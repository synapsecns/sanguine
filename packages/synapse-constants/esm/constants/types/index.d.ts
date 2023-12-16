import { BigNumber } from '@ethersproject/bignumber';
export type Chain = {
    id: number;
    chainSymbol: string;
    name: string;
    altName?: string;
    codeName: string;
    chainImg: any;
    layer: number;
    rpcUrls: {
        primary: string;
        fallback: string;
    };
    explorerUrl: string;
    explorerName: string;
    explorerImg: any;
    blockTime: number;
    nativeCurrency: {
        name: string;
        symbol: string;
        decimals: number;
    };
    priorityRank?: number;
    color?: string;
};
export type PoolToken = {
    symbol: string;
    percent: string;
    balance: string;
    balanceStr: string;
    token: Token;
    isLP: boolean;
    rawBalance: bigint;
};
export type Query = [string, string, BigNumber, BigNumber, string] & {
    swapAdapter: string;
    tokenOut: string;
    minAmountOut: BigNumber;
    deadline: BigNumber;
    rawParams: string;
};
export type PoolUserData = {
    name: string;
    tokens: PoolToken[];
    lpTokenBalance: bigint;
    nativeTokens?: any;
};
export type PoolData = {
    name: string;
    tokens: PoolToken[];
    totalLocked: number;
    totalLockedUSD: number;
    virtualPrice?: bigint;
    nativeTokens?: any;
    swapFee?: bigint;
};
interface TokensByChain {
    [cID: string]: Token[];
}
export type PoolCardInfo = {
    index: number;
    label: string;
    poolsByChain: TokensByChain;
};
export declare enum WalletId {
    MetaMask = "metaMask",
    WalletConnect = "walletConnect",
    CoinbaseWallet = "coinbaseWallet"
}
export interface IconProps {
    walletId?: string;
    className?: string;
}
export type PoolTokenObject = {
    token: Token;
    balance: string;
    rawBalance: bigint;
    isLP: boolean;
};
/**
 * Represents an ERC20-like token with a unique address, chainId, and some metadata.
 */
export declare class Token {
    addresses: {
        [x: number]: string;
    };
    wrapperAddresses?: Record<number, string>;
    decimals: number | Record<number, number>;
    symbol?: string;
    name?: string;
    logo?: any;
    icon?: any;
    poolName?: string;
    swapAddresses?: Record<number, string>;
    swapWrapperAddresses?: Record<number, string>;
    swapDepositAddresses?: Record<number, string>;
    swapEthAddresses?: Record<number, string>;
    routerIndex?: string;
    poolId: number | Record<number, number>;
    poolType?: string;
    poolTokens?: Token[];
    depositTokens?: Token[];
    nativeTokens?: Token[];
    description?: string;
    docUrl: string;
    forceMeta?: boolean;
    swapableType?: string;
    isNative: boolean;
    swapExceptions: number | Record<number, number[]>;
    visibilityRank: number;
    isMeta: boolean;
    isEthSwap: boolean;
    category: {
        bridge: boolean;
        swap: boolean;
        pool: boolean;
    };
    swapableOn: number[];
    display: boolean;
    legacy: boolean;
    priorityRank: number;
    chainId?: number;
    incentivized?: boolean;
    customRewardToken?: string;
    miniChefAddress: string;
    priorityPool?: boolean;
    color?: 'gray' | 'yellow' | 'green' | 'lime' | 'sky' | 'blue' | 'orange' | 'purple' | 'indigo' | 'cyan' | 'red';
    priceUnits?: string;
    notStake?: boolean;
    routeSymbol?: string;
    constructor({ addresses, wrapperAddresses, decimals, symbol, name, logo, poolName, swapAddresses, swapWrapperAddresses, swapDepositAddresses, swapEthAddresses, routerIndex, poolId, poolType, poolTokens, depositTokens, nativeTokens, description, docUrl, forceMeta, swapableType, isNative, swapExceptions, visibilityRank, isMeta, isEthSwap, category, swapableOn, display, legacy, priorityRank, chainId, incentivized, customRewardToken, miniChefAddress, priorityPool, color, priceUnits, notStake, routeSymbol, }: {
        addresses: {
            [x: number]: string;
        };
        wrapperAddresses?: Record<number, string>;
        decimals?: number | Record<number, number>;
        symbol?: string;
        name?: string;
        logo?: any;
        poolName?: string;
        swapAddresses?: Record<number, string>;
        swapWrapperAddresses?: Record<number, string>;
        swapDepositAddresses?: Record<number, string>;
        swapEthAddresses?: Record<number, string>;
        routerIndex?: string;
        poolId?: number | Record<number, number>;
        poolType?: string;
        poolTokens?: Token[];
        depositTokens?: Token[];
        nativeTokens?: Token[];
        description?: string;
        docUrl?: string;
        forceMeta?: boolean;
        swapableType?: string;
        isNative?: boolean;
        swapExceptions?: number | Record<number, number[]>;
        visibilityRank?: number;
        isMeta?: boolean;
        isEthSwap?: boolean;
        category?: {
            bridge: boolean;
            swap: boolean;
            pool: boolean;
        };
        swapableOn?: number[];
        display?: boolean;
        legacy?: boolean;
        priorityRank: number;
        chainId?: number;
        incentivized?: boolean;
        customRewardToken?: string;
        miniChefAddress?: string;
        priorityPool?: boolean;
        color?: 'gray' | 'yellow' | 'green' | 'lime' | 'sky' | 'blue' | 'orange' | 'purple' | 'indigo' | 'cyan' | 'red';
        priceUnits?: string;
        notStake?: boolean;
        routeSymbol?: string;
    });
}
export {};
