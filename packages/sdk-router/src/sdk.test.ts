import { Provider } from '@ethersproject/abstract-provider'
import { providers as etherProvider } from 'ethers'
import { BigNumber } from '@ethersproject/bignumber'

import { SynapseSDK } from './sdk'

describe('SynapseSDK', () => {
  const arbitrumProvider: Provider = new etherProvider.JsonRpcProvider(
    'https://arb1.arbitrum.io/rpc'
  )
  const avalancheProvider: Provider = new etherProvider.JsonRpcProvider(
    'https://api.avax.network/ext/bc/C/rpc'
  )

  describe('#constructor', () => {
    it('fails with unequal amount of chains to providers', () => {
      const chainIds = [42161, 43114]
      const providers = [arbitrumProvider]
      expect(() => new SynapseSDK(chainIds, providers)).toThrowError(
        'Amount of chains and providers does not equal'
      )
    })

  })

  describe('bridgeQuote', () => {
    it('test', async () => {
      const chainIds = [42161, 43114]
      const providers = [arbitrumProvider, avalancheProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const quotes = await Synapse.bridgeQuote(
        42161,
        43114,
        '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8',
        '0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664',
        BigNumber.from('10000000000000000000')
      )
      console.log(quotes)
      // await expect(bridgeTokens.length).toEqual(1)
    })
  })

  describe('bridge', () => {
    it('test', async () => {
      const chainIds = [42161, 43114]
      const providers = [arbitrumProvider, avalancheProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const quotes = await Synapse.bridgeQuote(
        42161,
        43114,
        '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8',
        '0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664',
        BigNumber.from('20000000')
      )
      console.log(quotes)

      console.log(
        await Synapse.bridge(
          '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
          42161,
          43114,
          '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8',
          BigNumber.from('20000000'),
          quotes.originQuery!,
          quotes.destQuery!
        )
      )
    })
  })
})
