import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, providers } from 'ethers'

import {
  ETH_NUSD,
  ETH_USDC,
  PUBLIC_PROVIDER_URLS,
  ROUTER_ADDRESS_MAP,
  SupportedChainId,
} from '../constants'
import { SynapseRouter } from './synapseRouter'
import { RouterQuery } from './query'
import { BridgeToken } from './types'

describe('SynapseRouter', () => {
  const ethAddress = ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
  const ethProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.ETH]
  )

  const recipient = '0x0000000000000000000000000000000000001337'
  const ethNusd = '0x1B84765dE8B7566e4cEAF4D0fD3c5aF52D3DdE4F'
  const emptyQuery: RouterQuery = {
    swapAdapter: '0x0000000000000000000000000000000000000000',
    tokenOut: '0x0000000000000000000000000000000000000000',
    minAmountOut: BigNumber.from(0),
    deadline: BigNumber.from(0),
    rawParams: '0x',
  }

  describe('#constructor', () => {
    it('fails with undefined chain id', () => {
      expect(
        () => new SynapseRouter(undefined as any, ethProvider, ethAddress)
      ).toThrow('CHAIN_ID_UNDEFINED')
    })

    it('fails with zero chain id', () => {
      expect(() => new SynapseRouter(0, ethProvider, ethAddress)).toThrow(
        'CHAIN_ID_UNDEFINED'
      )
    })

    it('fails with undefined provider', () => {
      expect(
        () =>
          new SynapseRouter(SupportedChainId.ETH, undefined as any, ethAddress)
      ).toThrow('PROVIDER_UNDEFINED')
    })

    it('fails with undefined address', () => {
      expect(
        () =>
          new SynapseRouter(SupportedChainId.ETH, ethProvider, undefined as any)
      ).toThrow('ADDRESS_UNDEFINED')
    })

    it('fails with empty address', () => {
      expect(
        () => new SynapseRouter(SupportedChainId.ETH, ethProvider, '')
      ).toThrow('ADDRESS_UNDEFINED')
    })
  })

  describe('ETH SynapseRouter', () => {
    const synapseRouter = new SynapseRouter(
      SupportedChainId.ETH,
      ethProvider,
      ethAddress
    )

    it('Bridge with nUSD', async () => {
      const { data } = await synapseRouter.bridge(
        recipient,
        SupportedChainId.ARBITRUM,
        ethNusd,
        BigNumber.from(10),
        emptyQuery,
        emptyQuery
      )
      expect(data?.length).toBeGreaterThan(0)
    })

    it('Fetches bridge tokens for USDC', async () => {
      const expectedTokens: BridgeToken[] = [
        {
          symbol: 'nUSD',
          token: ETH_NUSD,
        },
        {
          symbol: 'USDC',
          token: ETH_USDC,
        },
      ]
      const bridgeTokens = await synapseRouter.getBridgeTokens(ETH_USDC)
      expect(bridgeTokens).toEqual(expectedTokens)
    })
  })

  // TODO (Chi): figure out mocking, add more tests
})
