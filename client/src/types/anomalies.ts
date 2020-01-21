import { Entity } from 'parsing/common';
import { AlertAction } from 'types/common';


export enum NetworkNodeType {
  SingleEntity = 'SINGLE_ENTITY',
  EntityList = 'ENTITY_LIST',
}

export interface AnomalyAlertItem {
  anomalyType: string,
  anomalyDescription: string,
  sender: AnomalyAlertNetworkNode
  receiver: AnomalyAlertNetworkNode
  baseline: number
  baselineHigh: number|null
  baselineLow: number|null
  actual: number
  actualHigh: number|null
  actualLow: number|null
  count: number
  firstSeen: string,
  lastSeen: string,
  actionTaken: AlertAction,
  subItems: AnomalyAlertItem[]|null,
}

export interface AnomalyAlertNetworkNode {
  nodeType: NetworkNodeType,
  entity: Entity|null,
  entities: Entity[]|null,
}
