import { getWalletClient, getPublicClient } from '@wagmi/core'
import { providers } from 'ethers'
import { type HttpTransport, type WalletClient, type PublicClient } from 'viem'

import { wagmiConfig } from './wagmiConfig'

export const publicClientToProvider = (publicClient: PublicClient) => {
  const { chain, transport } = publicClient
  const network = {
    chainId: chain.id,
    name: chain.name,
    ensAddress: chain.contracts?.ensRegistry?.address,
  }
  if (transport.type === 'fallback') {
    return new providers.FallbackProvider(
      (transport.transports as ReturnType<HttpTransport>[]).map(
        ({ value }) => new providers.JsonRpcProvider(value?.url, network)
      )
    )
  }
  return new providers.JsonRpcProvider(transport.url, network)
}

/** Action to convert a viem Public Client to an ethers.js Provider. */
export const getEthersProvider = ({ chainId }: { chainId?: number } = {}) => {
  const publicClient = getPublicClient(wagmiConfig, { chainId })
  return publicClientToProvider(publicClient as any)
}

export const walletClientToSigner = (walletClient: WalletClient) => {
  console.log(walletClient)
  const { account, chain, transport } = walletClient
  const network = {
    chainId: chain.id,
    name: chain.name,
    ensAddress: chain.contracts?.ensRegistry?.address,
  }
  // @ts-ignore
  const provider = new providers.Web3Provider(transport, network)
  const signer = provider.getSigner(account.address)
  return signer
}

/** Action to convert a viem Wallet Client to an ethers.js Signer. */
export const getEthersSigner = async ({
  chainId,
}: { chainId?: number } = {}) => {
  const walletClient = await getWalletClient(wagmiConfig, {
    chainId,
  })
  if (!walletClient) {
    return undefined
  }
  return walletClientToSigner(walletClient)
}
