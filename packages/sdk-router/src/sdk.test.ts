import { Provider } from '@ethersproject/abstract-provider'
import { providers as etherProvider } from 'ethers'
import { BigNumber } from '@ethersproject/bignumber'

import { SynapseSDK } from './sdk'
import { CCTP_ROUTER_ADDRESS, SupportedChainId } from './constants'

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

    it('succeeds with SynapseCCTPRouter instantiation for keys in CCTP_ROUTER_ADDRESS', () => {
      const chainIds = [
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        SupportedChainId.AVALANCHE,
      ]
      const providers = [ethProvider, arbitrumProvider, avalancheProvider]

      const sdk = new SynapseSDK(chainIds, providers)

      for (const chainId of chainIds) {
        if (CCTP_ROUTER_ADDRESS.hasOwnProperty(chainId)) {
          expect(sdk.synapseCCTPRouters[chainId]).toBeDefined()
        } else {
          expect(sdk.synapseCCTPRouters[chainId]).toBeUndefined()
        }
      }
    })
  })

  describe('getBridgeTokens', () => {
    const destChainId = SupportedChainId.ETH
    const tokenOut = '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48'
    const sdk = new SynapseSDK([destChainId], [ethProvider])

    it('fetches bridge tokens for Synapse router', async () => {
      // Assuming you have a way to mock the bridge tokens returned by the router
      const mockedRouterBridgeTokens = [
        { symbol: 'nUSD', token: '0x1B84765dE8B7566e4cEAF4D0fD3c5aF52D3DdE4F' },
        { symbol: 'USDC', token: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48' },
      ]

      const routerBridgeTokens = await sdk.getBridgeTokens(
        destChainId,
        tokenOut,
        sdk.synapseRouters[destChainId]
      )

      // Assert that the function returned the correct bridge tokens for the router
      expect(routerBridgeTokens).toEqual(mockedRouterBridgeTokens)
    })

    it('fetches bridge tokens for CCTP router', async () => {
      // Assuming you have a way to mock the bridge tokens returned by the CCTP router
      const mockedCCTPBridgeTokens = [
        {
          symbol: 'CCTP.USDC',
          token: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
        },
      ]

      const routerCCTPTokens = await sdk.getBridgeTokens(
        destChainId,
        tokenOut,
        sdk.synapseCCTPRouters[destChainId]
      )

      // Assert that the function returned the correct bridge tokens for the CCTP router
      expect(routerCCTPTokens).toEqual(mockedCCTPBridgeTokens)
    })
  })

  describe('getOriginQueries', () => {
    it('fetches origin queries from both SynapseRouter and SynapseCCTPRouter', async () => {
      const originChainId = SupportedChainId.ETH
      const tokenIn = '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48'
      const routerTokenSymbols = ['USDC']
      const cctpTokenSymbols = ['CCTP.USDC']
      const amountIn = BigNumber.from('100000000')
      const sdk = new SynapseSDK([originChainId], [ethProvider])

      const synapseRouterOriginQuery = await sdk.getOriginQueries(
        sdk.synapseRouters[originChainId],
        tokenIn,
        routerTokenSymbols,
        amountIn
      )
      const synapseCCTPRouterOriginQuery = await sdk.getOriginQueries(
        sdk.synapseCCTPRouters[originChainId],
        tokenIn,
        cctpTokenSymbols,
        amountIn
      )

      expect(synapseRouterOriginQuery).toBeTruthy()
      expect(synapseCCTPRouterOriginQuery).toBeTruthy()

      expect(synapseRouterOriginQuery[0].minAmountOut.gt(0)).toBeTruthy()
      expect(synapseCCTPRouterOriginQuery[0].minAmountOut.gt(0)).toBeTruthy()
    })
  })

  describe('getDestinationQueries', () => {
    it('fetches destination queries from both SynapseRouter and SynapseCCTPRouter', async () => {
      const destChainId = SupportedChainId.ETH
      const tokenOut = '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48'
      const routerRequests = [
        { symbol: 'nUSD', amountIn: BigNumber.from('100000000000000000000') },
      ]
      const cctpRequests = [
        { symbol: 'CCTP.USDC', amountIn: BigNumber.from('100000000') },
      ]
      const sdk = new SynapseSDK([destChainId], [ethProvider])

      const synapseRouterDestinationQuery = await sdk.getDestinationQueries(
        sdk.synapseRouters[destChainId],
        routerRequests,
        tokenOut
      )
      const synapseCCTPRouterDestinationQuery = await sdk.getDestinationQueries(
        sdk.synapseCCTPRouters[destChainId],
        cctpRequests,
        tokenOut
      )

      // Ensure the function returned queries
      expect(synapseRouterDestinationQuery).toBeTruthy()
      expect(synapseCCTPRouterDestinationQuery).toBeTruthy()

      // Ensure minAmountOut is greater than 0
      expect(synapseRouterDestinationQuery[0].minAmountOut.gt(0)).toBeTruthy()
      expect(
        synapseCCTPRouterDestinationQuery[0].minAmountOut.gt(0)
      ).toBeTruthy()
    })
  })

  describe('bridgeQuote', () => {
    it('CCTP: ETH > Arbitrum', async () => {
      const chainIds = [1, 42161]
      const providers = [ethProvider, arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const result = await Synapse.bridgeQuote(
        1,
        42161,
        '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
        '0xaf88d065e77c8cc2239327c5edb3a432268e5831',
        BigNumber.from('100000000')
      )
      if (!result) {
        // console.log(result)
        throw Error
      }

      const { feeConfig, originQuery, destQuery, routerAddress } = result

      expect(feeConfig?.bridgeFee).toBeGreaterThan(0)
      expect(originQuery).not.toBeNull()
      expect(destQuery).not.toBeNull()

      expect(routerAddress?.length).toBeGreaterThan(0)

      const { data, to } = await Synapse.bridge(
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        '0xD359bc471554504f683fbd4f6e36848612349DDF',
        1,
        42161,
        '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
        BigNumber.from('100000000'),
        originQuery!,
        destQuery!
      )
      console.log(data)
      console.log(to)
      expect(data?.length).toBeGreaterThan(0)
      expect(to?.length).toBeGreaterThan(0)
    })
  })

  // test arb usdc > op usdc
  describe('bridgeQuote', () => {
    it('test', async () => {
      const chainIds = [42161, 10]
      const providers = [arbitrumProvider, optimisimProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const result = await Synapse.bridgeQuote(
        42161,
        10,
        '0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8',
        '0x7F5c764cBc14f9669B88837ca1490cCa17c31607',
        BigNumber.from('100000000')
      )

      if (!result) {
        // console.log(result)
        throw Error
      }

      const { feeConfig, originQuery, destQuery, routerAddress } = result

      expect(feeConfig?.bridgeFee).toBeGreaterThan(0)
      expect(originQuery).not.toBeNull()
      expect(destQuery).not.toBeNull()
      checkQueryFields(originQuery)
      checkQueryFields(destQuery)

      expect(routerAddress?.length).toBeGreaterThan(0)

      const { data, to } = await Synapse.bridge(
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        '0x7e7a0e201fd38d3adaa9523da6c109a07118c96a',
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

  // // test avax usdc.e > bsc usdc
  describe('bridgeQuote', () => {
    it('test', async () => {
      const chainIds = [43114, 56]
      const providers = [avalancheProvider, bscProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const result = await Synapse.bridgeQuote(
        43114,
        56,
        '0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664',
        '0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d',
        BigNumber.from('100000000')
      )

      if (!result) {
        // console.log(result)
        throw Error
      }

      const { feeConfig, originQuery, destQuery } = result

      expect(feeConfig?.bridgeFee).toBeGreaterThan(0)
      expect(originQuery).not.toBeNull()
      expect(destQuery).not.toBeNull()
      checkQueryFields(originQuery)
      checkQueryFields(destQuery)
      const { data, to } = await Synapse.bridge(
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        '0x7e7a0e201fd38d3adaa9523da6c109a07118c96a',
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
      const result = await Synapse.bridgeQuote(
        43114,
        56,
        '0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664',
        '0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d',
        BigNumber.from('100000000'),
        BigNumber.from('100000000')
      )

      if (!result) {
        // console.log(result)
        throw Error
      }

      const { originQuery, destQuery } = result

      expect(originQuery?.deadline).toStrictEqual(BigNumber.from('100000000'))
      checkQueryFields(originQuery)
      checkQueryFields(destQuery)
      const { data, to } = await Synapse.bridge(
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        '0x7e7a0e201fd38d3adaa9523da6c109a07118c96a',
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
      const result = await Synapse.bridgeQuote(
        42161,
        43114,
        '0x8D9bA570D6cb60C7e3e0F31343Efe75AB8E65FB1',
        '0x321E7092a180BB43555132ec53AaA65a5bF84251',
        BigNumber.from('10000000000000000000')
      )
      if (!result) {
        // console.log(result)
        throw Error
      }

      const { feeConfig, originQuery, destQuery } = result

      expect(feeConfig?.bridgeFee).toBeGreaterThan(0)
      expect(originQuery).not.toBeNull()
      expect(destQuery).not.toBeNull()
      checkQueryFields(originQuery)
      checkQueryFields(destQuery)
      const { data, to } = await Synapse.bridge(
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        '0x7e7a0e201fd38d3adaa9523da6c109a07118c96a',
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
      const result = await Synapse.bridgeQuote(
        1,
        42161,
        '',
        '',
        BigNumber.from('1000000000000000000')
      )

      if (!result) {
        // console.log(result)
        throw Error
      }

      const { feeConfig, originQuery, destQuery } = result

      expect(feeConfig?.bridgeFee).toBeGreaterThan(0)
      expect(originQuery).not.toBeNull()
      expect(destQuery).not.toBeNull()
      const { data, to } = await Synapse.bridge(
        '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
        '0x7e7a0e201fd38d3adaa9523da6c109a07118c96a',
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
  // this test uses the swap eth wrapper pool address to test if calculate remove liquidity works on ETH Wrapper pools
  describe('calculate remove liquidity', () => {
    it('test', async () => {
      const chainIds = [42161]
      const providers = [arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const { amounts, routerAddress } = await Synapse.calculateRemoveLiquidity(
        42161,
        '0x1c3fe783a7c06bfAbd124F2708F5Cc51fA42E102',
        BigNumber.from('1000000')
      )
      expect(amounts.length).toBeGreaterThan(0)
      expect(amounts[0].value.toNumber()).toBeGreaterThan(0)
      expect(routerAddress?.length).toBeGreaterThan(0)
    })
  })
  describe('calculate remove liquidity one', () => {
    it('test', async () => {
      const chainIds = [42161]
      const providers = [arbitrumProvider]
      const Synapse = new SynapseSDK(chainIds, providers)
      const { amount, routerAddress } =
        await Synapse.calculateRemoveLiquidityOne(
          42161,
          '0x1c3fe783a7c06bfAbd124F2708F5Cc51fA42E102',
          BigNumber.from('1000000'),
          1
        )
      expect(Object.keys(amount)?.length).toBeGreaterThan(0)
      expect(amount.value.toNumber()).toBeGreaterThan(0)
      expect(routerAddress.length).toBeGreaterThan(0)
    })
  })
})
