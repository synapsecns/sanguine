import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, providers } from 'ethers'

import {
  PUBLIC_PROVIDER_URLS,
  CCTP_ROUTER_ADDRESS_MAP,
  SupportedChainId,
} from '../constants'
import { SynapseCCTPRouter } from './synapseCCTPRouter'
import { CCTPRouterQuery } from './query'

describe('SynapseCCTPRouter', () => {
  const ethAddress = CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
  const ethProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.ETH]
  )

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

  describe('Bridge', () => {
    const cctpRouter = new SynapseCCTPRouter(
      SupportedChainId.ETH,
      ethProvider,
      ethAddress
    )

    it('Bridge with USDC', async () => {
      const { data } = await cctpRouter.bridge(
        recipient,
        SupportedChainId.ARBITRUM,
        ethUSDC,
        BigNumber.from(10),
        emptyQuery,
        emptyQuery
      )
      expect(data?.length).toBeGreaterThan(0)
    })
  })

  // TODO (Chi): figure out mocking, add more tests
})
