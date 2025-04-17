import { ethers } from 'ethers'

import fastBridgeAbi from '../constants/abis/fastBridge.json'
import fastBridgeRouterAbi from '../constants/abis/fastBridgeRouter.json'
import { FAST_BRIDGE_ROUTER_ADDRESS_MAP } from '../constants'
import { CHAINS_BY_ID } from '../constants/chains'

const providerCache = new Map()

export const getBridgeStatus = async (
  originChainId: string | number,
  kappa: string
) => {
  try {
    const chainInfo = CHAINS_BY_ID[originChainId]
    const rpcUrl = chainInfo.rpcUrls.primary || chainInfo.rpcUrls.fallback

    let provider = providerCache.get(rpcUrl)

    if (!provider) {
      provider = new ethers.providers.JsonRpcProvider(rpcUrl)
      providerCache.set(rpcUrl, provider)
    }

    const routerAddress = FAST_BRIDGE_ROUTER_ADDRESS_MAP[Number(originChainId)]

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
  } catch (error) {
    return null
  }
}
