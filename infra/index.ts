import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as awsx from "@pulumi/awsx";
import * as random from "@pulumi/random";

const PRODUCTION = "prd";

const config = new pulumi.Config();
const stack = pulumi.getStack();

// VPC
const vpc = new awsx.ec2.Vpc(`terminal_web-${stack}`, {});

// Security groups
const appSG = new aws.ec2.SecurityGroup("app", {
  vpcId: vpc.id,
  ingress: [
    {
      protocol: "tcp",
      fromPort: 80,
      toPort: 80,
      cidrBlocks: ["0.0.0.0/0"],
    },
  ],
  egress: [
    {
      protocol: "-1",
      fromPort: 0,
      toPort: 0,
      cidrBlocks: ["0.0.0.0/0"],
    },
  ],
});
const dbSG = new aws.ec2.SecurityGroup("db", {
  vpcId: vpc.id,
  ingress: [
    {
      protocol: "tcp",
      fromPort: 5432,
      toPort: 5432,
      securityGroups: [appSG.id],
    },
  ],
});

// RDS
const dbSubnetGroup = new aws.rds.SubnetGroup(`terminal_web-${stack}`, {
  subnetIds: vpc.privateSubnetIds,
});
const db_pw = new random.RandomPassword(`terminal_web_db_password-${stack}`, {
  length: 16,
  special: false,
});
const dbName = "terminal_web_production";
let dbArgs = {
    
  engine: "postgres",
  engineVersion: "14.2",
  instanceClass: "db.t4g.micro",
  allocatedStorage: 10,
  dbName: dbName,
  username: "terminal_web",
  password: db_pw.result,
  backupWindow: "16:00-18:00", // 0:00-2:00 SGT
  vpcSecurityGroupIds: [dbSG.id],
  dbSubnetGroupName: dbSubnetGroup.name,
  skipFinalSnapshot: true,
};
if (stack === PRODUCTION) {
  Object.assign(dbArgs, { backupRetentionPeriod: 7 });
}
const db = new aws.rds.Instance(`terminal-web-${stack}`, dbArgs);
const dbUrl = pulumi.interpolate`postgresql://${db.username}:${db.password}@${db.endpoint}/${db.dbName}`;

// Fargate
const repo = new awsx.ecr.Repository(`terminal_web-${stack}`);
export const img = repo.buildAndPushImage("../group");
const cluster = new awsx.ecs.Cluster(`terminal_web-${stack}`, { vpc: vpc });
const lb = new awsx.lb.ApplicationListener(`lb`, {
  vpc,
  external: true,
  port: 80,
  targetGroup: {
    healthCheck: {
      path: "/health",
      matcher: "200,204",
    },
    port: 80,
  },
});
const taskRole = awsx.ecs.FargateTaskDefinition.createTaskRole(`terminal_web_app-${stack}-task`)
const sshPolicy = new aws.iam.Policy(`terminal_web-${stack}`, {
  description: "SSH into container policy",
  policy: `{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": [
          "ssmmessages:CreateControlChannel",
          "ssmmessages:CreateDataChannel",
          "ssmmessages:OpenControlChannel",
          "ssmmessages:OpenDataChannel"
        ],
        "Effect": "Allow",
        "Resource": "*"
      }
    ]
  }`,
});
const taskRoleSSHPolicyAttachment = new aws.iam.RolePolicyAttachment(`terminal_web_app_ssh-${stack}`, {
  role: taskRole.name,
  policyArn: sshPolicy.arn,
});
const app = new awsx.ecs.FargateService(`terminal_web_app-${stack}`, {
  cluster,
  securityGroups: [appSG.id],
  taskDefinitionArgs: {
    containers: {
      web: {
        image: img,
        cpu: 256,
        portMappings: [lb],
      },
    },
    taskRole
  },
  desiredCount: 1,
  enableExecuteCommand: true,
  waitForSteadyState: false,
});

export const APP_URL = lb.endpoint.hostname;
