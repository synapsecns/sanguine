import { ethers } from 'ethers'

import fastBridgeAbi from '../constants/abis/fastBridge.json'
import fastBridgeRouterAbi from '../constants/abis/fastBridgeRouter.json'
import { FAST_BRIDGE_ROUTER_ADDRESS_MAP } from '../constants'
import { CHAINS_BY_ID } from '../constants/chains'

export const getBridgeStatus = async (
  originChainId: string | number,
  kappa: string
) => {
  const chainInfo = CHAINS_BY_ID[originChainId]
  const rpcUrl = chainInfo.rpcUrls.primary || chainInfo.rpcUrls.fallback
  const provider = new ethers.providers.JsonRpcProvider(rpcUrl)

  const routerAddress = FAST_BRIDGE_ROUTER_ADDRESS_MAP[originChainId]

  const fastBridgeRouterContract = new ethers.Contract(
    routerAddress,
    fastBridgeRouterAbi,
    provider
  )

  const fastBridgeAddress = await fastBridgeRouterContract.fastBridge()

  const fastBridgeContract = new ethers.Contract(
    fastBridgeAddress,
    fastBridgeAbi,
    provider
  )

  const status = await fastBridgeContract.bridgeStatuses(
    kappa.startsWith('0x') ? kappa : `0x${kappa}`
  )

  return status
}
