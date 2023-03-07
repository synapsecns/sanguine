import BRIDGE_ZAP_ABI from '@abis/bridgeZap.json'
import L1_BRIDGE_ZAP_ABI from '@abis/l1bridgezap.json'
import META_BRIDGE_ZAP_ABI from '@abis/l2bridgezap.json'

import {
  BRIDGE_ZAP_ADDRESSES,
} from '@constants/bridge'

import { ChainId } from '@constants/networks'

import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useContract, useGenericContract } from '@hooks/contracts/useContract'




export function useBridgeZapContract() {
  const { chainId } = useActiveWeb3React()
  let abi
  if (chainId == ChainId.ETH) {
    abi = BRIDGE_ZAP_ABI
  } else if (chainId == ChainId.DFK) {
    abi = L1_BRIDGE_ZAP_ABI
  } else {
    abi = META_BRIDGE_ZAP_ABI
  }
  const address = BRIDGE_ZAP_ADDRESSES[chainId]
  const bridgeZapContract = useContract(address, abi)

  return bridgeZapContract
}


export function useGenericBridgeZapContract(chainId) {
  let abi
  if (chainId == ChainId.ETH) {
    abi = BRIDGE_ZAP_ABI
  } else {
    abi = META_BRIDGE_ZAP_ABI
  }
  const address = BRIDGE_ZAP_ADDRESSES[chainId]
  const bridgeZapContract = useGenericContract(chainId, address, abi)

  return bridgeZapContract
}