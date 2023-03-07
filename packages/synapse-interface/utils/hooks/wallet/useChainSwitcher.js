import _ from 'lodash'

import { ChainId, CHAIN_PARAMS } from '@constants/networks'

import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useTerraWallet } from '@hooks/terra/useTerraWallet'

const ETH_KEYS_TO_REMOVE = [
  'chainName',
  'nativeCurrency',
  'rpcUrls',
  'blockExplorerUrls'
]


export function useChainSwitcher() {
  const { account, library } = useActiveWeb3React()
  const { connectTerraStation, terraAddress, disconnect } = useTerraWallet()

  return async function triggerChainSwitch(itemChainId) {
    const params = CHAIN_PARAMS[itemChainId]

    if (itemChainId == ChainId.TERRA) {
      return connectTerraStation()
    } else if (itemChainId == ChainId.ETH) {
      const ethParams = _.pickBy(params, (value, key) => ETH_KEYS_TO_REMOVE.indexOf(key) === -1)
      return library?.send('wallet_switchEthereumChain', [ethParams, account])
    } else {
      return library?.send('wallet_addEthereumChain', [params, account])
    }
  }
}


