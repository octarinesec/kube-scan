export enum EntityType {
  Any = '00_ENTITY_TYPE_ANY',
  Workload = '01_ENTITY_TYPE_WORKLOAD',
  WorkloadReplica = '02_ENTITY_TYPE_WORKLOAD_REPLICA',
  Ip = '03_ENTITY_TYPE_IP',
  IpGroup = '05_ENTITY_TYPE_IP_GROUP',
  Subnet = '05_ENTITY_TYPE_SUBNET',
  Hostname = '06_ENTITY_TYPE_HOSTNAME',
  Labels = '07_ENTITY_TYPE_LABELS',
  External = '08_ENTITY_TYPE_EXTERNAL',
}


export enum ActivityDirection {
  Ingress = '01_INGRESS',
  Egress = '02_EGRESS',
  Internal = '03_Internal',
}

export enum ActivityHistoryStatus {
  Loading = 'Loading',
  Loaded = 'Loaded',
}

export type ActivityHistory = {
  status: ActivityHistoryStatus;
  counts: ActivityHistoryCountItem[],
  minWithCount: number,
  maxCount: number,
}

export type ActivityHistorySegment = {
  counts: ActivityHistoryCountItem[],
  totalItemCount: number,
}

export type ActivityHistoryCountItem = {
  index?: number,
  timeStr: string,
  count: number,
}

export type SenderReceiverPair = {
  senderID: string,
  receiverID: string,
}

export type ActivityHistoryMap = {[propName: string]: ActivityHistory};

