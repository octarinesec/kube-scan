import {AlertAction, AlertReason, Protocol} from 'types/common';
import { Entity } from 'parsing/common';

export { Protocol } from 'types/common';

export enum AlertCode {
  Violation = 'violation',
  Unknown = 'unknown',
}

export enum AlertSeverity {
  Violation = 'violation',
  Unknown = 'unknown',
}

export enum AlertMethod {
  Read = 'read',
  Write = 'write',
  Get = 'get',
  Post = 'post',
  Put = 'put',
  Options = 'options',
  Delete = 'delete',
  Produce = 'produce',
  Consume = 'consume',
}

interface RequestInfo {
  protocol: Protocol,
  endpoint: string,
  method: AlertMethod | null,
  topic: string | null,
  group: string | null,
  senderPort: number | null,
  receiverPort: number | null,
}

export interface AlertItem {
  id: string,
  sender: Entity,
  receiver: Entity,
  code: AlertCode,
  severity: AlertSeverity,
  firstSeen: string,
  lastSeen: string,
  actionTaken: AlertAction | null,
  actionReason: AlertReason | null,
  count: number | null,
  requestInfo: RequestInfo,
}

export interface AlertGroup {
  sender: Entity,
  receiver: Entity,
  code: AlertCode,
  severity: AlertSeverity,
  firstSeen: string,
  lastSeen: string,
  actionTaken: AlertAction | null,
  actionReason: AlertReason | null,
  count: number | null,
  items: AlertItem[],
  methods: AlertMethod[],
  endpoints: string[],
  groupKey: string,
  protocol: Protocol,
  topics: string[],
  groups: string[],
}
