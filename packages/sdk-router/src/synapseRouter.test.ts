import { Provider } from '@ethersproject/abstract-provider'
import { providers } from 'ethers'
import JSBI from 'jsbi'

import { SupportedChainId, ROUTER_ADDRESS } from './constants'
import { SynapseRouter } from './synapseRouter'

describe('SynapseRouter', () => {
  const arbitrumProvider: Provider = new providers.JsonRpcProvider(
    'https://arb1.arbitrum.io/rpc'
  )

  const provider: Provider = new providers.AlchemyProvider('mainnet', 'demo')

  describe('#constructor', () => {
    it('fails with undefined chain id', () => {
      expect(() => new SynapseRouter(undefined as any, provider)).toThrow(
        'CHAIN_ID_UNDEFINED'
      )
    })

    it('fails with undefined provider', () => {
      expect(() => new SynapseRouter(1, undefined as any)).toThrow(
        'PROVIDER_UNDEFINED'
      )
    })

    it('succeeds with correct contract address', () => {
      for (const chainId of Object.keys(SupportedChainId).filter(
        (x) => parseInt(x, 10) > 0
      )) {
        expect(
          new SynapseRouter(chainId as any, provider).routerContract.address
        ).toEqual(ROUTER_ADDRESS[chainId as keyof object])
      }
    })
  })

  describe('Bridge', () => {
    it('Bridge with nUSD', async () => {
      const synapseRouter = new SynapseRouter(42161, arbitrumProvider)
      const { data } =
        await synapseRouter.routerContract.populateTransaction.bridge(
          '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
          43114,
          '0x2913E812Cf0dcCA30FB28E6Cac3d2DCFF4497688',
          JSBI.BigInt(10),
          {
            swapAdapter: '0x0000000000000000000000000000000000000000',
            tokenOut: '0x0000000000000000000000000000000000000000',
            minAmountOut: 0,
            deadline: 0,
            rawParams: '0x0000000000000000000000000000000000000000',
          },
          {
            swapAdapter: '0x0000000000000000000000000000000000000000',
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
