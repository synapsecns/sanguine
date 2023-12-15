import * as all from './master';
export * from './master';
export const sortChains = (chains) => Object.values(chains).sort((a, b) => { var _a, _b; return ((_a = b.priorityRank) !== null && _a !== void 0 ? _a : 0) - ((_b = a.priorityRank) !== null && _b !== void 0 ? _b : 0); });
export const CHAINS_ARR = Object.values(all)
    .filter((item) => typeof item !== 'number')
    .sort((a, b) => { var _a, _b; return ((_a = b.priorityRank) !== null && _a !== void 0 ? _a : 0) - ((_b = a.priorityRank) !== null && _b !== void 0 ? _b : 0); });
const getChainEnumById = () => {
    const outObj = {};
    CHAINS_ARR.map((chain) => {
        outObj[chain.id] = chain.codeName;
    });
    return outObj;
};
const getids = () => {
    const outObj = {};
    CHAINS_ARR.map((chain) => {
        outObj[chain.chainSymbol] = chain.id;
    });
    return outObj;
};
const getChainsByID = () => {
    const outObj = {};
    CHAINS_ARR.map((chain) => {
        outObj[chain.id] = chain;
    });
    return outObj;
};
export const CHAIN_ENUM_BY_ID = getChainEnumById();
export const CHAIN_IDS = getids(); // used to be ids
export const CHAINS_BY_ID = getChainsByID();
export const ORDERED_CHAINS_BY_ID = CHAINS_ARR.map((chain) => String(chain.id));
export const PAUSED_FROM_CHAIN_IDS = [];
export const PAUSED_TO_CHAIN_IDS = [all.DOGE.id];
export const ChainId = {
    ETH: 1,
    ROPSTEN: 3,
    RINKEBY: 4,
    GÖRLI: 5,
    OPTIMISM: 10,
    CRONOS: 25,
    KOVAN: 42,
    BSC: 56,
    POLYGON: 137,
    FANTOM: 250,
    BOBA: 288,
    METIS: 1088,
    MOONBEAM: 1284,
    MOONRIVER: 1285,
    DOGECHAIN: 2000,
    CANTO: 7700,
    KLAYTN: 8217,
    HARDHAT: 31337,
    ARBITRUM: 42161,
    BASE: 8453,
    AVALANCHE: 43114,
    DFK: 53935,
    AURORA: 1313161554,
    HARMONY: 1666600000,
    TERRA: 121014925, //"columbus-5", the day columbus reportedly landed in america followed by 5
};
export const AcceptedChainId = Object.fromEntries(Object.entries(ChainId).map(([key, value]) => [value, key]));
