import { Provider } from '@ethersproject/abstract-provider'
import { AddressZero } from '@ethersproject/constants'
import { SynapseRouter } from './synapseRouter'
import invariant from 'tiny-invariant';
import { BigintIsh } from './constants'
import {BigNumber} from '@ethersproject/bignumber'
import { BytesLike } from '@ethersproject/bytes'

export class SynapseSDK { 
    public synapseRouters: any

    constructor(chainIds: number[], providers: Provider[]) {
        invariant(chainIds.length == providers.length, `Amount of chains and providers does not equal`);
        this.synapseRouters = {}
        for (let i = 0; i < chainIds.length; i++) {
            this.synapseRouters[chainIds[i]] = new SynapseRouter(chainIds[i] as any, providers[i])
        }
    }

    public async bridgeQuote(
        originChainId: number,
        destChainId: number,
        tokenIn: string,
        tokenOut: string,
        amountIn: BigintIsh
    ): Promise<any> {
        let originQuery
        let destQuery
        let originRouter: SynapseRouter = this.synapseRouters[originChainId];
        let destRouter: SynapseRouter = this.synapseRouters[destChainId];
        // Step 0: find connected bridge tokens on destination
        let bridgeTokens = await destRouter.routerContract.getConnectedBridgeTokens(tokenOut);

        if (bridgeTokens.length == 0) throw Error('No bridge tokens found for this route')

        let symbols: string[] = [];
        for (let token of bridgeTokens) {
            if (token.symbol.length == 0) continue
            if (token.token == AddressZero) continue
            symbols.push(token.symbol)
        }

        // Step 1: perform a call to origin SynapseRouter

        let originQueries = await originRouter.routerContract.getOriginAmountOut(tokenIn, symbols, amountIn);

        // Step 2: form a list of Destination Requests
        // In practice, there is no need to pass the requests with amountIn = 0, but we will do it for code simplicity

        let requests: { symbol: string; amountIn: BigintIsh; }[] = [];

        for (let i = 0; i < bridgeTokens.length; i++) {
            requests.push({
                symbol: symbols[i],
                amountIn: originQueries[i].minAmountOut
            })
        }

        // Step 3: perform a call to destination SynapseRouter
        let destQueries = await destRouter.routerContract.getDestinationAmountOut(requests, tokenOut);
        // Step 4: find the best query (in practice, we could return them all)

        let maxAmountOut: BigNumber = BigNumber.from(0)
        for (let i = 0; i < destQueries.length; i++) {
            if (destQueries[i].minAmountOut.gt(maxAmountOut)) {
                maxAmountOut = destQueries[i].minAmountOut
                originQuery = originQueries[i]
                destQuery = destQueries[i]
            }
        };

        // // Throw error if origin quote is zero
        // if (originQuery.minAmountOut == 0) throw Error("No path found on origin")

        return {originQuery, destQuery}

    }

    public async bridge(
        to: string,
        originChainId: number,
        destChainId: number,
        token: string,
        amount: BigintIsh,
        originQuery: { swapAdapter: string; tokenOut: string; minAmountOut: BigintIsh; deadline: BigintIsh; rawParams: BytesLike; },
        destQuery: { swapAdapter: string; tokenOut: string; minAmountOut: BigintIsh; deadline: BigintIsh; rawParams: BytesLike; }
    ): Promise<any> {
        let originRouter: SynapseRouter = this.synapseRouters[originChainId];
        console.log(originQuery)
        console.log(destQuery)
        return await originRouter.routerContract.populateTransaction.bridge(to, destChainId, token, amount, originQuery, destQuery)
    }
}