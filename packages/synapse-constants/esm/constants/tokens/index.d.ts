import { Token } from '../types/index';
export * from './bridgeable';
interface TokensByChain {
    [cID: string]: Token[];
}
interface TokenByKey {
    [cID: string]: Token;
}
interface TokenMap {
    [chainId: string]: {
        [address: string]: Token;
    };
}
export declare const sortTokens: (tokens: Token[]) => Token[];
export declare const PAUSED_TOKENS_BY_CHAIN: {
    [key: string]: any;
};
export declare const findChainIdsWithPausedToken: (routeSymbol: string) => never[];
export declare const getBridgeableTokens: () => TokensByChain;
export declare const TOKENS_SORTED_BY_SWAPABLETYPE: (string | undefined)[];
export declare const TOKENS_SORTED_BY_SYMBOL: (string | undefined)[];
export declare const BRIDGABLE_TOKENS: TokensByChain;
export declare const tokenSymbolToToken: (chainId: number, symbol: string) => Token | undefined;
export declare const tokenAddressToToken: (chainId: number, tokenAddress: string) => Token | undefined;
export declare const TOKEN_HASH_MAP: TokenMap;
export declare const POOL_PRIORITY_RANKING: {
    [key: string]: any;
};
export declare const POOL_CHAINS_BY_NAME: {
    [key: string]: any;
};
export declare const POOL_BY_ROUTER_INDEX: TokenByKey;
export declare const POOLS_BY_CHAIN: TokensByChain;
export declare const DISPLAY_POOLS_BY_CHAIN: TokensByChain;
export declare const USD_POOLS_BY_CHAIN: {
    [key: string]: any;
};
export declare const ETH_POOLS_BY_CHAIN: {
    [key: string]: any;
};
export declare const LEGACY_POOLS_BY_CHAIN: TokensByChain;
interface StakableTokens {
    [chainId: string]: Token[];
}
export declare const STAKABLE_TOKENS: StakableTokens;
export declare const STAKING_MAP_TOKENS: {
    [key: string]: any;
};
