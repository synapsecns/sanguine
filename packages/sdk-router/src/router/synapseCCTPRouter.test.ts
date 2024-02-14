import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from 'ethers'
import { AddressZero } from '@ethersproject/constants'

import { ETH_USDC } from '../constants/testValues'
import { getTestProvider } from '../constants/testProviders'
import { CCTP_ROUTER_ADDRESS_MAP, SupportedChainId } from '../constants'
import { SynapseCCTPRouter } from './synapseCCTPRouter'
import { BridgeToken, CCTPRouterQuery } from '../module'
import { DestRequest } from './types'

describe('SynapseCCTPRouter', () => {
  const ethAddress = CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
  const ethProvider: Provider = getTestProvider(SupportedChainId.ETH)

  const recipient = '0x0000000000000000000000000000000000001337'
  const ethUSDC = '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48'
  const emptyQuery: CCTPRouterQuery = {
    routerAdapter: '0x0000000000000000000000000000000000000000',
    tokenOut: '0x0000000000000000000000000000000000000000',
    minAmountOut: BigNumber.from(0),
    deadline: BigNumber.from(0),
    rawParams: '0x',
  }

  describe('#constructor', () => {
    it('fails with undefined chain id', () => {
      expect(
        () => new SynapseCCTPRouter(undefined as any, ethProvider, ethAddress)
      ).toThrow('CHAIN_ID_UNDEFINED')
    })

    it('fails with zero chain id', () => {
      expect(() => new SynapseCCTPRouter(0, ethProvider, ethAddress)).toThrow(
        'CHAIN_ID_UNDEFINED'
      )
    })

    it('fails with undefined provider', () => {
      expect(
        () =>
          new SynapseCCTPRouter(
            SupportedChainId.ETH,
            undefined as any,
            ethAddress
          )
      ).toThrow('PROVIDER_UNDEFINED')
    })

    it('fails with undefined address', () => {
      expect(
        () =>
          new SynapseCCTPRouter(
            SupportedChainId.ETH,
            ethProvider,
            undefined as any
          )
      ).toThrow('ADDRESS_UNDEFINED')
    })

    it('fails with empty address', () => {
      expect(
        () => new SynapseCCTPRouter(SupportedChainId.ETH, ethProvider, '')
      ).toThrow('ADDRESS_UNDEFINED')
    })
  })

  describe('ETH SynapseCCTPRouter', () => {
    const cctpRouter = new SynapseCCTPRouter(
      SupportedChainId.ETH,
      ethProvider,
      ethAddress
    )

    describe('bridge', () => {
      it('Bridge with USDC', async () => {
        const { data } = await cctpRouter.bridge(
          recipient,
          SupportedChainId.ARBITRUM,
          ethUSDC,
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
            symbol: 'CCTP.USDC',
            token: ETH_USDC,
          },
        ]
        const tokens = await cctpRouter.getBridgeTokens(ETH_USDC)
        expect(tokens).toEqual(expectedTokens)
      })
    })

    describe('getOriginQueries', () => {
      it('Fetches origin queries for USDC', async () => {
        const originQueries = await cctpRouter.getOriginQueries(
          ETH_USDC,
          ['CCTP.USDC'],
          BigNumber.from(10).pow(9)
        )
        expect(originQueries.length).toBe(1)
        expect(originQueries[0].tokenOut).toBe(ETH_USDC)
        expect(originQueries[0].minAmountOut.gt(0)).toBe(true)
      })

      it('Does not filter zero amount queries', async () => {
        const originQueries = await cctpRouter.getOriginQueries(
          ETH_USDC,
          ['CCTP.USDC', 'fakeSymbol'],
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
            symbol: 'CCTP.USDC',
            amountIn: BigNumber.from(10).pow(9),
          },
        ]
        const destQueries = await cctpRouter.getDestinationQueries(
          destRequests,
          ETH_USDC
        )
        expect(destQueries.length).toBe(1)
        expect(destQueries[0].tokenOut).toBe(ETH_USDC)
        expect(destQueries[0].minAmountOut.gt(0)).toBe(true)
      })

      it('Does not filter zero amount queries', async () => {
        const destRequests: DestRequest[] = [
          {
            symbol: 'CCTP.USDC',
            amountIn: BigNumber.from(10).pow(9),
          },
          {
            symbol: 'fakeSymbol',
            amountIn: BigNumber.from(10).pow(9),
          },
        ]
        const destQueries = await cctpRouter.getDestinationQueries(
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
