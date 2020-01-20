
export enum Protocol {
  Any = '*',
  Http = 'http',
  Tcp = 'tcp',
  Kafka = 'kafka',
}


export interface SocketInfo {
  port: number | null,
  address: string | null,
}


export type EnumValueMap<E extends string, T> = {
  [key in E]: T;
};

export type EnumValuePartialMap<E extends string, T> = {
  [key in E]?: T;
};

export enum FetchStatus {
  Fetched = 'FETCHED',
  Fetching = 'FETCHING',
  Error = 'ERROR',
}

export enum AlertAction {
  // these names match the value from the API
  // don't change unless the API changes
  Allow = 'allow',
  Alert = 'alert',
  Block = 'block',
}

export enum AlertReason {
  // these names match the value from the API
  // don't change unless the API changes
  NONE = 'NONE',
  InboundTrafficDefault = 'InboundTrafficDefault',
  PrivateEgressTrafficDefault = 'PrivateEgressTrafficDefault',
  PublicEgressTrafficDefault = 'PublicEgressTrafficDefault',
  MaliciousEgressTrafficDefault = 'MaliciousEgressTrafficDefault',
  OutboundTrafficDefault = 'OutboundTrafficDefault',
}

export enum AlertType {
  K8s = 'K8S',
  Anomaly = 'ANOMALY',
  Threat = 'THREAT',
  Network = 'NETWORK',
}
