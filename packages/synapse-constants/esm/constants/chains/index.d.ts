import { Chain } from '../types';
export * from './master';
export type ChainsByChainID = {
    [cID: number]: Chain;
};
export declare const sortChains: (chains: Chain[]) => Chain[];
export declare const CHAINS_ARR: Chain[];
export declare const CHAIN_ENUM_BY_ID: Record<number, string>;
export declare const CHAIN_IDS: {
    [key: string]: any;
};
export declare const CHAINS_BY_ID: ChainsByChainID;
export declare const ORDERED_CHAINS_BY_ID: string[];
export declare const PAUSED_FROM_CHAIN_IDS: never[];
export declare const PAUSED_TO_CHAIN_IDS: number[];
export declare const ChainId: {
    ETH: number;
    ROPSTEN: number;
    RINKEBY: number;
    GÖRLI: number;
    OPTIMISM: number;
    CRONOS: number;
    KOVAN: number;
    BSC: number;
    POLYGON: number;
    FANTOM: number;
    BOBA: number;
    METIS: number;
    MOONBEAM: number;
    MOONRIVER: number;
    DOGECHAIN: number;
    CANTO: number;
    KLAYTN: number;
    HARDHAT: number;
    ARBITRUM: number;
    BASE: number;
    AVALANCHE: number;
    DFK: number;
    AURORA: number;
    HARMONY: number;
    TERRA: number;
};
export declare const AcceptedChainId: {
    [k: string]: string;
};
