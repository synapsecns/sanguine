import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, providers } from 'ethers'
import { AddressZero } from '@ethersproject/constants'
import { describe, expect, it } from 'vitest'

import {
  ETH_NUSD,
  ETH_USDC,
  PUBLIC_PROVIDER_URLS,
  ROUTER_ADDRESS_MAP,
  SupportedChainId,
} from '../constants'
import { SynapseRouter } from './synapseRouter'
import { RouterQuery } from './query'
import { BridgeToken, DestRequest } from './types'

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

    describe('bridge', () => {
      it('Bridge with nUSD', async () => {
        const { data } = await synapseRouter.bridge(
          recipient,
          SupportedChainId.ARBITRUM,
          ethNusd,
          BigNumber.from(10),
          emptyQuery,
          emptyQuery
        )
        // TODO: check data correctness
        expect(data?.length).toBeGreaterThan(0)
      })
    })

    describe('getBridgeTokens', () => {
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

    describe('getOriginQueries', () => {
      it('Fetches origin queries for USDC', async () => {
        const originQueries = await synapseRouter.getOriginQueries(
          ETH_USDC,
          ['USDC', 'nUSD'],
          BigNumber.from(10).pow(9)
        )
        expect(originQueries.length).toBe(2)
        expect(originQueries[0].tokenOut).toBe(ETH_USDC)
        expect(originQueries[0].minAmountOut.gt(0)).toBe(true)
        expect(originQueries[1].tokenOut).toBe(ETH_NUSD)
        expect(originQueries[1].minAmountOut.gt(0)).toBe(true)
      })

      it('Does not filter zero amount queries', async () => {
        const originQueries = await synapseRouter.getOriginQueries(
          ETH_USDC,
          ['USDC', 'fakeSymbol'],
          BigNumber.from(10).pow(9)
        )
        expect(originQueries.length).toBe(2)
        expect(originQueries[0].tokenOut).toBe(ETH_USDC)
        expect(originQueries[0].minAmountOut.gt(0)).toBe(true)
        expect(originQueries[1].tokenOut).toBe(AddressZero)
        expect(originQueries[1].minAmountOut).toEqual(BigNumber.from(0))
      })
    })

    describe('getDestinationQueries', () => {
      it('Fetches destination queries for USDC', async () => {
        // $1000 requests: should be above minimum bridge fee
        const destRequests: DestRequest[] = [
          {
            symbol: 'USDC',
            amountIn: BigNumber.from(10).pow(9),
          },
          {
            symbol: 'nUSD',
            amountIn: BigNumber.from(10).pow(21),
          },
        ]
        const destQueries = await synapseRouter.getDestinationQueries(
          destRequests,
          ETH_USDC
        )
        expect(destQueries.length).toBe(2)
        expect(destQueries[0].tokenOut).toBe(ETH_USDC)
        expect(destQueries[0].minAmountOut.gt(0)).toBe(true)
        expect(destQueries[1].tokenOut).toBe(ETH_USDC)
        expect(destQueries[1].minAmountOut.gt(0)).toBe(true)
      })

      it('Does not filter zero amount queries', async () => {
        const destRequests: DestRequest[] = [
          {
            symbol: 'USDC',
            amountIn: BigNumber.from(10).pow(9),
          },
          {
            symbol: 'fakeSymbol',
            amountIn: BigNumber.from(10).pow(9),
          },
        ]
        const destQueries = await synapseRouter.getDestinationQueries(
          destRequests,
          ETH_USDC
        )
        expect(destQueries.length).toBe(2)
        expect(destQueries[0].tokenOut).toBe(ETH_USDC)
        expect(destQueries[0].minAmountOut.gt(0)).toBe(true)
        expect(destQueries[1].minAmountOut).toEqual(BigNumber.from(0))
      })
    })
  })

  // TODO (Chi): figure out mocking, add more tests
})
