import * as CHAINS from '../chains/master';
import { getAddress } from '@ethersproject/address';
export var WalletId;
(function (WalletId) {
    WalletId["MetaMask"] = "metaMask";
    WalletId["WalletConnect"] = "walletConnect";
    WalletId["CoinbaseWallet"] = "coinbaseWallet";
})(WalletId || (WalletId = {}));
/**
 * Represents an ERC20-like token with a unique address, chainId, and some metadata.
 */
export class Token {
    constructor({ addresses, wrapperAddresses, decimals, symbol, name, logo, poolName, swapAddresses, swapWrapperAddresses, swapDepositAddresses, swapEthAddresses, routerIndex, poolId, poolType, poolTokens, depositTokens, nativeTokens, description, docUrl = '', forceMeta, swapableType, isNative = false, swapExceptions, visibilityRank, isMeta, isEthSwap, category, swapableOn, display, legacy, priorityRank, chainId, incentivized, customRewardToken, miniChefAddress, priorityPool, color, priceUnits, notStake, routeSymbol, }) {
        this.decimals = {}; // list of decimals on each chain
        this.poolId = {}; // list of pool ids on each chain
        this.docUrl = ''; // token doc url
        this.isNative = false; // is native
        this.swapExceptions = {}; // for specifying tokens where limited dest chains are available.
        this.visibilityRank = 0; // rank in which token is displayed, least visible is 0, there is no max
        this.isMeta = false; // is meta
        this.isEthSwap = false; // is eth swap
        this.category = {
            bridge: true,
            swap: true,
            pool: true,
        }; // list of categories on each chain
        this.swapableOn = []; // list of chains where token is swapable
        this.display = true; // display token
        this.legacy = false; // legacy token
        this.priorityPool = false; // priority pool
        const isMetaVar = Boolean(swapDepositAddresses || forceMeta);
        this.addresses = validateAddresses(addresses);
        this.wrapperAddresses = wrapperAddresses;
        // this.decimals             = decimals
        this.decimals = makeMultiChainObj(decimals);
        this.symbol = symbol;
        this.name = name;
        this.icon = logo;
        this.poolName = poolName;
        this.swapAddresses = swapAddresses;
        this.swapWrapperAddresses = swapWrapperAddresses;
        this.swapDepositAddresses = swapDepositAddresses;
        this.swapEthAddresses = swapEthAddresses;
        this.routerIndex = routerIndex;
        this.poolTokens = poolTokens;
        this.nativeTokens = nativeTokens !== null && nativeTokens !== void 0 ? nativeTokens : poolTokens;
        this.depositTokens = depositTokens !== null && depositTokens !== void 0 ? depositTokens : this.nativeTokens;
        this.description = description;
        this.docUrl = docUrl !== null && docUrl !== void 0 ? docUrl : '';
        this.poolId = makeMultiChainObj(poolId);
        this.poolType = poolType;
        this.visibilityRank = visibilityRank !== null && visibilityRank !== void 0 ? visibilityRank : 0;
        this.isMeta = isMeta !== null && isMeta !== void 0 ? isMeta : false;
        this.isEthSwap = swapEthAddresses ? true : false;
        this.isNative = isNative !== null && isNative !== void 0 ? isNative : false;
        this.swapableType = swapableType;
        this.swapExceptions = swapExceptions !== null && swapExceptions !== void 0 ? swapExceptions : [];
        this.category = category !== null && category !== void 0 ? category : { bridge: true, swap: true, pool: true };
        this.swapableOn = swapableOn !== null && swapableOn !== void 0 ? swapableOn : [];
        this.display = display !== null && display !== void 0 ? display : true;
        this.legacy = legacy !== null && legacy !== void 0 ? legacy : false;
        this.priorityRank = priorityRank;
        this.chainId = chainId;
        this.incentivized = incentivized;
        this.customRewardToken = customRewardToken;
        this.miniChefAddress = miniChefAddress !== null && miniChefAddress !== void 0 ? miniChefAddress : '';
        this.priorityPool = priorityPool !== null && priorityPool !== void 0 ? priorityPool : false;
        this.color = color !== null && color !== void 0 ? color : 'gray';
        this.priceUnits = priceUnits !== null && priceUnits !== void 0 ? priceUnits : 'USD';
        this.notStake = notStake !== null && notStake !== void 0 ? notStake : false;
        this.routeSymbol = routeSymbol;
    }
}
const makeMultiChainObj = (valOrObj) => {
    if (typeof valOrObj === 'object') {
        return valOrObj;
    }
    else {
        const obj = {};
        for (const chain of Object.values(CHAINS)) {
            obj[chain.id] = valOrObj;
        }
        return obj;
    }
};
const validateAddresses = (addresses) => {
    const reformatted = {};
    for (const chainId in addresses) {
        reformatted[chainId] = addresses[chainId]
            ? getAddress(addresses[chainId])
            : '';
    }
    return reformatted;
};
