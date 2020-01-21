import { Protocol } from 'types/common';

export enum Operation {
  Any = '*',
  Read = 'read',
  Write = 'write',
  Produce = 'produce',
  Consume = 'consume',
}

export enum Decision {
  Allow = 'allow',
  Alert = 'alert',
  Block = 'block',
}

export enum MaplNodeType {
  Workload = 'workload',
  Subnet = 'subnet',
  Hostname = 'hostname',
  Any = '*',
}

export type Receiver = {
  receiverName: string,
  receiverType: MaplNodeType,
}

export type Sender = {
  senderName: string,
  senderType: MaplNodeType,
}

export enum ResourceType {
  All = '*',
  Path = 'path',
  Port = 'port',
  KafkaTopic = 'kafkatopic',
  KafkaGroup = 'kafkagroup',
}

export type Resource = {
  resourceNames: string[],
  resourceType: ResourceType,
}

export interface AndCondition {
  attribute: string,
  method: string,
  value: string,
}

export type Condition = {
  andConditions: AndCondition[],
}

export type MaplRuleType = {
  sender: Sender | null;
  receiver: Receiver | null;
  protocol: Protocol | null;
  operation: Operation | null;
  decision: Decision | null;
  resource: Resource | null;
  conditions: Condition[] | null;
}

export enum TimeFrame {
  Day = 'DAY',
  Week = 'WEEK',
}

