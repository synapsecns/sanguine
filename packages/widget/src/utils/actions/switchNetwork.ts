import { toHexStr } from '@/utils/toHexStr'
import { CHAINS_BY_ID } from '@/constants/chains'

export const switchNetwork = async (chainId: number, provider: any) => {
  try {
    if (typeof chainId !== 'number') {
      throw new Error('Require ChainId')
    }
    if (!provider) {
      throw new Error('Require Provider')
    }

    const chain = CHAINS_BY_ID[chainId]
    const hexChainId: string = toHexStr(chainId)

    await provider.send('wallet_addEthereumChain', [
      {
        chainId: hexChainId,
        chainName: chain.networkName,
        nativeCurrency: chain.nativeCurrency,
        rpcUrls: [chain.networkUrl],
        blockExplorerUrls: [chain.explorerUrl],
      },
    ])
    await provider.send('wallet_switchEthereumChain', [{ chainId: hexChainId }])
  } catch (error) {
    console.error('[Synapse Widget] Error switching networks: ', error)
  }
}
