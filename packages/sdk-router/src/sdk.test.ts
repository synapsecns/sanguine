import { SynapseSDK } from "./sdk";
import { Provider } from '@ethersproject/abstract-provider'
import { providers } from 'ethers'
import {BigNumber} from '@ethersproject/bignumber'

describe('SynapseSDK', () => {
    const arbitrumProvider: Provider = new providers.JsonRpcProvider("https://arb1.arbitrum.io/rpc")
    const avalancheProvider: Provider = new providers.JsonRpcProvider("https://api.avax.network/ext/bc/C/rpc")

    describe('#constructor', () => {
        it('fails with unequal amount of chains to providers', () => {
            let chainIds = [42161, 43114];
            let providers = [arbitrumProvider]
            expect(() => new SynapseSDK(chainIds, providers)).toThrowError("Amount of chains and providers does not equal")
        })

        it('succeeds with equal amount of chains to providers', async () => {
            let chainIds = [42161, 43114];
            let providers = [arbitrumProvider, avalancheProvider]
            let Synapse = new SynapseSDK(chainIds, providers);
            expect(() => Synapse.synapseRouters.length.toEqual(2))
        })
    })

    describe('bridgeQuote', function() {
        it('test', async() => {
            let chainIds = [42161, 43114];
            let providers = [arbitrumProvider, avalancheProvider]
            let Synapse = new SynapseSDK(chainIds, providers);
            let quotes = await Synapse.bridgeQuote(42161, 43114, "0xff970a61a04b1ca14834a43f5de4533ebddb5cc8", "0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664", BigNumber.from("10000000000000000000"))
            console.log(quotes)
            // await expect(bridgeTokens.length).toEqual(1)
        })
    })

    describe('bridge', function() {
        it('test', async() => {
            let chainIds = [42161, 43114];
            let providers = [arbitrumProvider, avalancheProvider]
            let Synapse = new SynapseSDK(chainIds, providers);
            let quotes = await Synapse.bridgeQuote(42161, 43114, "0xff970a61a04b1ca14834a43f5de4533ebddb5cc8", "0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664", BigNumber.from("20000000"))
            console.log(quotes)

            console.log(await Synapse.bridge(
                "0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9",
                42161,
                43114,
                "0xff970a61a04b1ca14834a43f5de4533ebddb5cc8",
                BigNumber.from("20000000"),
                quotes.originQuery,
                quotes.destQuery,
            ))
        })
    })
})
