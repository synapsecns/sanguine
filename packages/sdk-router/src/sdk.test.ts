import { Provider } from '@ethersproject/abstract-provider'
import { providers as etherProvider } from 'ethers'
import { BigNumber } from '@ethersproject/bignumber'

import { SynapseSDK } from './sdk'

const checkQueryFields = (query: any) => {
  expect(query.swapAdapter).not.toBeNull()
  expect(query.tokenOut).not.toBeNull()
  expect(query.minAmountOut).not.toBeNull()
  expect(query.deadline).not.toBeNull()
  expect(query.rawParams).not.toBeNull()
}
jest.setTimeout(30000)
// TODO add more tests checking parity of to/from values
// as well as more token/chain combinations
describe('SynapseSDK', () => {
  const ethProvider: Provider = new etherProvider.JsonRpcProvider(
    'https://rpc.builder0x69.io	'
  )
  const arbitrumProvider: Provider = new etherProvider.JsonRpcProvider(
    'https://arb1.arbitrum.io/rpc'
  )
  const avalancheProvider: Provider = new etherProvider.JsonRpcProvider(
    'https://api.avax.network/ext/bc/C/rpc'
  )
  const optimisimProvider: Provider = new etherProvider.JsonRpcProvider(
    'https://1rpc.io/op'
  )
  const bscProvider: Provider = new etherProvider.JsonRpcProvider(
    'https://endpoints.omniatech.io/v1/bsc/mainnet/public'
  )
  // test constructor
  describe('Test Constructor', () => {
    it('fails with unequal amount of chains to providers', () => {
      const chainIds = [42161, 43114, 10]
      const providers = [arbitrumProvider]
      expect(() => new SynapseSDK(chainIds, providers)).toThrowError(
        'Amount of chains and providers does not equal'
      )
    })
  })

  // test arb usdc > op usdc
  describe('bridgeQuote', () => {
    it('test', async () => {
      const chainIds = [42161, 10]
      const providers = [arbitrumProvider, optimisimProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const { feeConfig, originQuery, destQuery, routerAddress } =
        await Synapse.bridgeQuote(
          42161,
          10,
          '0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8',
          '0x7F5c764cBc14f9669B88837ca1490cCa17c31607',
          BigNumber.from('100000000')
        )
      expect(feeConfig?.bridgeFee).toBeGreaterThan(0)
      expect(originQuery).not.toBeNull()
      expect(destQuery).not.toBeNull()
      checkQueryFields(originQuery)
      checkQueryFields(destQuery)

      expect(routerAddress?.length).toBeGreaterThan(0)

      const { data, to } = await Synapse.bridge(
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        42161,
        10,
        '0x7F5c764cBc14f9669B88837ca1490cCa17c31607',
        BigNumber.from('100000000'),
        originQuery!,
        destQuery!
      )
      expect(data?.length).toBeGreaterThan(0)
      expect(to?.length).toBeGreaterThan(0)
    })
  })

  // test avax usdc.e > bsc usdc
  describe('bridgeQuote', () => {
    it('test', async () => {
      const chainIds = [43114, 56]
      const providers = [avalancheProvider, bscProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const { feeConfig, destQuery, originQuery } = await Synapse.bridgeQuote(
        43114,
        56,
        '0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664',
        '0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d',
        BigNumber.from('100000000')
      )
      expect(feeConfig?.bridgeFee).toBeGreaterThan(0)
      expect(originQuery).not.toBeNull()
      expect(destQuery).not.toBeNull()
      checkQueryFields(originQuery)
      checkQueryFields(destQuery)
      const { data, to } = await Synapse.bridge(
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        43114,
        56,
        '0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d',
        BigNumber.from('100000000'),
        originQuery!,
        destQuery!
      )
      expect(data?.length).toBeGreaterThan(0)
      expect(to?.length).toBeGreaterThan(0)
    })
  })

  // test avax usdc.e > bsc usdc
  describe('test custom deadline', () => {
    it('test', async () => {
      const chainIds = [43114, 56]
      const providers = [avalancheProvider, bscProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const { destQuery, originQuery } = await Synapse.bridgeQuote(
        43114,
        56,
        '0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664',
        '0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d',
        BigNumber.from('100000000'),
        BigNumber.from('100000000')
      )

      expect(originQuery?.deadline).toStrictEqual(BigNumber.from('100000000'))
      checkQueryFields(originQuery)
      checkQueryFields(destQuery)
      const { data, to } = await Synapse.bridge(
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        43114,
        56,
        '0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d',
        BigNumber.from('100000000'),
        originQuery!,
        destQuery!
      )
      expect(data?.length).toBeGreaterThan(0)
      expect(to?.length).toBeGreaterThan(0)
    })
  })
  // test gohn arb > gohn avax
  describe('bridge', () => {
    it('test', async () => {
      const chainIds = [42161, 43114]
      const providers = [arbitrumProvider, avalancheProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const { feeConfig, originQuery, destQuery } = await Synapse.bridgeQuote(
        42161,
        43114,
        '0x8D9bA570D6cb60C7e3e0F31343Efe75AB8E65FB1',
        '0x321E7092a180BB43555132ec53AaA65a5bF84251',
        BigNumber.from('10000000000000000000')
      )
      expect(feeConfig?.bridgeFee).toBeGreaterThan(0)
      expect(originQuery).not.toBeNull()
      expect(destQuery).not.toBeNull()
      checkQueryFields(originQuery)
      checkQueryFields(destQuery)
      const { data, to } = await Synapse.bridge(
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        42161,
        43114,
        '0x321E7092a180BB43555132ec53AaA65a5bF84251',
        BigNumber.from('10000000000000000000'),
        originQuery!,
        destQuery!
      )
      expect(data?.length).toBeGreaterThan(0)
      expect(to?.length).toBeGreaterThan(0)
    })
  })
  describe('bridge', () => {
    it('test', async () => {
      const chainIds = [1, 42161]
      const providers = [ethProvider, arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const { feeConfig, originQuery, destQuery } = await Synapse.bridgeQuote(
        1,
        42161,
        '',
        '',
        BigNumber.from('1000000000000000000')
      )

      expect(feeConfig?.bridgeFee).toBeGreaterThan(0)
      expect(originQuery).not.toBeNull()
      expect(destQuery).not.toBeNull()
      const { data, to } = await Synapse.bridge(
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        42161,
        43114,
        '0x321E7092a180BB43555132ec53AaA65a5bF84251',
        BigNumber.from('10000000000000000000'),
        originQuery!,
        destQuery!
      )
      expect(data?.length).toBeGreaterThan(0)
      expect(to?.length).toBeGreaterThan(0)
    })
  })
  describe('swap quote', () => {
    it('test', async () => {
      const chainIds = [42161]
      const providers = [arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const { query } = await Synapse.swapQuote(
        42161,
        '0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8',
        '0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9',
        BigNumber.from('1000000')
      )
      expect(query).not.toBeNull()
      checkQueryFields(query)
      const { data, to } = await Synapse.swap(
        42161,
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        '0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8',
        BigNumber.from('1000000'),
        query!
      )
      expect(data?.length).toBeGreaterThan(0)
      expect(to?.length).toBeGreaterThan(0)
    })
  })
  describe('bridge gas', () => {
    it('test', async () => {
      const chainIds = [42161]
      const providers = [arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const gas = await Synapse.getBridgeGas(42161)
      expect(Number(gas.toString())).toBeGreaterThan(-1)
    })
  })
  describe('get all pools', () => {
    it('test', async () => {
      const chainIds = [42161]
      const providers = [arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const pools = await Synapse.getAllPools(42161)
      expect(pools?.length).toBeGreaterThan(0)
      expect(pools?.[0]?.tokens?.[0]?.token?.length).toBeGreaterThan(0)
    })
  })
  describe('get pool info', () => {
    it('test', async () => {
      const chainIds = [42161]
      const providers = [arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const poolInfo = await Synapse.getPoolInfo(
        42161,
        '0x9Dd329F5411466d9e0C488fF72519CA9fEf0cb40'
      )
      expect(poolInfo?.tokens?.toString()?.length).toBeGreaterThan(0)
      expect(poolInfo?.lpToken?.length).toBeGreaterThan(0)
    })
  })
  describe('get pool tokens', () => {
    it('test', async () => {
      const chainIds = [42161]
      const providers = [arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const poolTokens = await Synapse.getPoolTokens(
        42161,
        '0x9Dd329F5411466d9e0C488fF72519CA9fEf0cb40'
      )
      expect(poolTokens?.length).toBeGreaterThan(0)
      expect(poolTokens?.[0]?.token?.length).toBeGreaterThan(0)
    })
  })

  describe('calculate add liquidity', () => {
    it('test', async () => {
      const chainIds = [42161]
      const providers = [arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const tokenAmount = BigNumber.from('1000000')
      const tokenAmount2 = BigNumber.from('2000000')
      const amount = await Synapse.calculateAddLiquidity(
        42161,
        '0xa067668661C84476aFcDc6fA5D758C4c01C34352',
        {
          '0x3ea9B0ab55F34Fb188824Ee288CeaEfC63cf908e': tokenAmount,
          '0x82aF49447D8a07e3bd95BD0d56f35241523fBab1': tokenAmount2,
        }
      )
      expect(amount?.amount.toString()?.length).toBeGreaterThan(0)
      expect(amount?.routerAddress.length).toBeGreaterThan(0)
    })
  })
  describe('calculate add liquidity 2', () => {
    it('test', async () => {
      const chainIds = [1]
      const providers = [ethProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const tokenAmount = BigNumber.from('1000000')
      const tokenAmount2 = BigNumber.from('0')
      const tokenAmount3 = BigNumber.from('0')
      const amount = await Synapse.calculateAddLiquidity(
        1,
        '0x1116898DdA4015eD8dDefb84b6e8Bc24528Af2d8',
        {
          '0x6b175474e89094c44da98b954eedeac495271d0f': tokenAmount,
          '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48': tokenAmount2,
          '0xdac17f958d2ee523a2206206994597c13d831ec7': tokenAmount3,
        }
      )
      expect(amount?.toString()?.length).toBeGreaterThan(0)
    })
  })
  describe('calculate remove liquidity', () => {
    it('test', async () => {
      const chainIds = [42161]
      const providers = [arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const amounts = await Synapse.calculateRemoveLiquidity(
        42161,
        '0xa067668661C84476aFcDc6fA5D758C4c01C34352',
        BigNumber.from('1000000')
      )
      expect(Object.keys(amounts.amounts)?.length).toBeGreaterThan(0)
      expect(
        amounts.amounts[Object.keys(amounts.amounts)[0]].value.toNumber()
      ).toBeGreaterThan(0)
      expect(amounts?.routerAddress.length).toBeGreaterThan(0)
    })
  })
  describe('calculate remove liquidity one', () => {
    it('test', async () => {
      const chainIds = [42161]
      const providers = [arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const amounts = await Synapse.calculateRemoveLiquidityOne(
        42161,
        '0xa067668661C84476aFcDc6fA5D758C4c01C34352',
        BigNumber.from('1000000'),
        '0x6b175474e89094c44da98b954eedeac495271d0f'
      )
      expect(Object.keys(amounts.amount)?.length).toBeGreaterThan(0)
      expect(amounts.amount.value.toNumber()).toBeGreaterThan(0)
      expect(amounts?.routerAddress.length).toBeGreaterThan(0)
    })
  })
})
