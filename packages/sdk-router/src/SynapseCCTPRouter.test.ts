import { Provider } from '@ethersproject/abstract-provider'
import { providers } from 'ethers'
import JSBI from 'jsbi'

import { CCTP_ROUTER_ADDRESS } from './constants'
import { SynapseCCTPRouter } from './SynapseCCTPRouter'

describe('SynapseCCTPRouter', () => {
  const arbitrumProvider: Provider = new providers.JsonRpcProvider(
    'https://arb1.arbitrum.io/rpc'
  )

  const provider: Provider = new providers.AlchemyProvider('mainnet', 'demo')

  describe('#constructor', () => {
    it('fails with undefined chain id', () => {
      expect(() => new SynapseCCTPRouter(undefined as any, provider)).toThrow(
        'CHAIN_ID_UNDEFINED'
      )
    })

    it('fails with undefined provider', () => {
      expect(() => new SynapseCCTPRouter(1, undefined as any)).toThrow(
        'PROVIDER_UNDEFINED'
      )
    })

    it('succeeds with correct contract address', () => {
      for (const chainId of Object.keys(CCTP_ROUTER_ADDRESS)) {
        expect(
          new SynapseCCTPRouter(Number(chainId), provider).routerContract
            .address
        ).toEqual(
          CCTP_ROUTER_ADDRESS[
            // eslint-disable-next-line
            parseInt(chainId, 10) as keyof typeof CCTP_ROUTER_ADDRESS
          ]
        )
      }
    })
  })

  describe('Bridge', () => {
    it('Bridge with nUSD', async () => {
      const synapseCCTPRouter = new SynapseCCTPRouter(42161, arbitrumProvider)
      const { data } =
        await synapseCCTPRouter.routerContract.populateTransaction.bridge(
          '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
          43114,
          '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
          JSBI.BigInt(10),
          {
            routerAdapter: '0x0000000000000000000000000000000000000000',
            tokenOut: '0x0000000000000000000000000000000000000000',
            minAmountOut: 0,
            deadline: 0,
            rawParams: '0x0000000000000000000000000000000000000000',
          },
          {
            routerAdapter: '0x0000000000000000000000000000000000000000',
            tokenOut: '0x0000000000000000000000000000000000000000',
            minAmountOut: 0,
            deadline: 0,
            rawParams: '0x0000000000000000000000000000000000000000',
          }
        )
      expect(data?.length).toBeGreaterThan(0)
    })
  })
})
