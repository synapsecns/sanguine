import BRIDGE_CONFIG_ABI from '@abis/bridgeConfig.json'
import { ChainId } from '@constants/networks'
import { BRIDGE_CONFIG_ADDRESSES } from '@constants/bridge'

import { useGenericContract } from '@hooks/contracts/useContract'



export function useBridgeConfigContract() {

  // return useGenericContract(ChainId.POLYGON, BRIDGE_CONFIG_ADDRESSES[ChainId.POLYGON], BRIDGE_CONFIG_ABI)
  return useGenericContract(ChainId.ETH, BRIDGE_CONFIG_ADDRESSES[ChainId.ETH], BRIDGE_CONFIG_ABI)
}