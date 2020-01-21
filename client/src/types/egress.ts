import { MaplRule } from 'parsing/rules/MaplRule';
import { AlertAction, EnumValueMap } from 'types/common';


export type EgressRule = {
  id: string,
  created: string,
  enabled: boolean,
  maplRule: MaplRule,
}


export enum EgressTrafficType {
  Private = 'PRIVATE',
  Public = 'PUBLIC',
  Malicious = 'MALICIOUS',
}


export type EgressActionMapping = EnumValueMap<EgressTrafficType, AlertAction>;


