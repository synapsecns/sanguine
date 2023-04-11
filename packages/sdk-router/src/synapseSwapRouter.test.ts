import { Provider } from '@ethersproject/abstract-provider'
import { providers } from 'ethers'
import { BigNumber } from '@ethersproject/bignumber'

import { SynapseSwapQuoter, LimitedTokenStructType } from './synapseSwapQuoter'

describe('SynapseSwapQuoter', () => {
  const arbitrumProvider: Provider = new providers.JsonRpcProvider(
    'https://arb1.arbitrum.io/rpc'
  )

  const provider: Provider = new providers.AlchemyProvider('mainnet', 'demo')

  describe('#constructor', () => {
    it('fails with undefined chain id', () => {
      expect(
        () => new SynapseSwapQuoter(undefined as any, provider, '')
      ).toThrow('CHAIN_ID_UNDEFINED')
    })

    it('fails with undefined provider', () => {
      expect(() => new SynapseSwapQuoter(1, undefined as any, '')).toThrow(
        'PROVIDER_UNDEFINED'
      )
    })
    it('fails with undefined address', () => {
      expect(
        () => new SynapseSwapQuoter(1, provider, undefined as any)
      ).toThrow('SWAP_QUOTER_ADDRESS_UNDEFINED')
    })
  })

  describe('get amount out', () => {
    it('get amount', async () => {
      const synapseSwapQuoter = new SynapseSwapQuoter(
        42161,
        arbitrumProvider,
        '0x78a83c17600add7447dbd6b8ab26330481075295'
      )
      console.log('hi', synapseSwapQuoter.swapQuoterContract.getAmountOut)

      const tokenIn: LimitedTokenStructType = {
        actionMask: 1,
        token: '0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8',
      }

      console.log(
        await synapseSwapQuoter.swapQuoterContract.getAmountOut(
          tokenIn,
          '0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1',
          BigNumber.from(10000)
        )
      )
    })
  })
})
