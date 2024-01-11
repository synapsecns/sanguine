import { getModuleSet } from './bridge'
import { Query, applySlippage } from '../module'
import { SynapseSDK } from '../sdk'

export function applyBridgeSlippage(
  this: SynapseSDK,
  bridgeModuleName: string,
  originQueryPrecise: Query,
  destQueryPrecise: Query,
  slipNumerator: number,
  slipDenominator: number = 10000
): { originQuery: Query; destQuery: Query } {
  const moduleSet = getModuleSet.call(this, bridgeModuleName)
  return moduleSet.applySlippage(
    originQueryPrecise,
    destQueryPrecise,
    slipNumerator,
    slipDenominator
  )
}

export const applySwapSlippage = (
  query: Query,
  slipNumerator: number,
  slipDenominator: number = 10000
): Query => {
  return applySlippage(query, slipNumerator, slipDenominator)
}
