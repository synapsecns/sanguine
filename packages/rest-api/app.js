"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.server = void 0;
var providers_1 = require("@ethersproject/providers");
var sdk_router_1 = require("@synapsecns/sdk-router");
var bignumber_1 = require("@ethersproject/bignumber");
var units_1 = require("@ethersproject/units");
var express = require("express");
// import express from 'express'
var chainsData = require("./config/chains.json");
var tokensData = require("./config/tokens.json");
var tokens = tokensData;
var chains = chainsData;
// Constants
var TEN = bignumber_1.BigNumber.from(10);
var tokenHtml = Object.keys(tokens)
    .map(function (symbol) {
    return '<b>' +
        String(symbol) +
        '</b>: <br/>' +
        String(Object.keys(tokens[symbol].addresses)
            .map(function (addrChainId) {
            return '<li>' +
                String(addrChainId) +
                ': ' +
                tokens[symbol].addresses[addrChainId] +
                '</li>';
        })
            .join('')) +
        '</br>';
})
    .join('');
// Set up Synapse SDK
var providers = [];
var chainIds = [];
for (var _i = 0, chains_1 = chains; _i < chains_1.length; _i++) {
    var chain = chains_1[_i];
    providers.push(new providers_1.JsonRpcProvider(chain.rpc));
    chainIds.push(chain.id);
}
// Define the sdk
var Synapse = new sdk_router_1.SynapseSDK(chainIds, providers);
// Set up express server
var app = express();
var port = process.env.PORT || 3000;
//Intro Message for UI
app.get('/', function (req, res) {
    res.send("\n    <h1>Welcome to the Synapse Rest API for swap and bridge quotes</h1>\n    <hr/>\n    <h2>Available Chains</h2>\n    <ul>\n     ".concat(chains
        .map(function (chain) {
        return '<li>' + String(chain.name) + ' (' + String(chain.id) + ')' + '</li>';
    })
        .join(''), "\n    </ul>\n    <h2>Available Tokens (symbols to use)</h2>\n    ").concat(tokenHtml));
});
//Swap Quote get request
app.get('/swap', function (req, res) { return __awaiter(void 0, void 0, void 0, function () {
    var query, chainId, fromTokenSymbol, toTokenSymbol, fromTokenAddress, toTokenAddress, fromTokenDecimals, toTokenDecimals, amount;
    var _a, _b, _c, _d, _e, _f, _g, _h;
    return __generator(this, function (_j) {
        query = req.query;
        chainId = query.chain;
        fromTokenSymbol = String(query.fromToken);
        toTokenSymbol = String(query.toToken);
        fromTokenAddress = (_b = (_a = tokens[fromTokenSymbol]) === null || _a === void 0 ? void 0 : _a.addresses) === null || _b === void 0 ? void 0 : _b[chainId];
        toTokenAddress = (_d = (_c = tokens[toTokenSymbol]) === null || _c === void 0 ? void 0 : _c.addresses) === null || _d === void 0 ? void 0 : _d[chainId];
        fromTokenDecimals = (_f = (_e = tokens[fromTokenSymbol]) === null || _e === void 0 ? void 0 : _e.decimals) === null || _f === void 0 ? void 0 : _f[chainId];
        toTokenDecimals = (_h = (_g = tokens[toTokenSymbol]) === null || _g === void 0 ? void 0 : _g.decimals) === null || _h === void 0 ? void 0 : _h[chainId];
        // Handle invalid params (either token symbols or chainIDs)
        // TODO: add error handling for missing params
        if (!fromTokenAddress ||
            !toTokenAddress ||
            !fromTokenDecimals ||
            !toTokenDecimals) {
            res.send("\n      <h1>Invalid Params</h1>\n      <hr/>\n      <b>Ensure that your request matches the following format: /swap?chain=1&fromToken=USDC&toToken=DAI&amount=100</b>\n      <h2>Available Tokens (symbols to use)</h2>\n      ".concat(tokenHtml));
            return [2 /*return*/];
        }
        amount = bignumber_1.BigNumber.from(query.amount).mul(TEN.pow(fromTokenDecimals));
        // Send request w/Synapse SDK
        Synapse.swapQuote(Number(chainId), fromTokenAddress, toTokenAddress, bignumber_1.BigNumber.from(amount))
            .then(function (resp) {
            // Check for stable swap (going in its 6 decimals but coming out its 18 decimals so we need to adjust)
            // Using arbitrary 6 decimals as a threshold for now
            // TODO: Router contract v2 should return the amount out with decimals for the out-out token not the out-in token (eg.nusd).
            // Add response field with adjusted maxAmountOutStr (to account for decimals)
            var payload = resp;
            payload.maxAmountOutStr = formatBNToString(resp.maxAmountOut, toTokenDecimals);
            res.json(payload);
        })
            .catch(function (err) {
            // TODO: do a better return here
            res.send("\n      <h1>Invalid Request</h1>\n      <code>".concat(err, "</code>\n      <hr/>\n      <b>Ensure that your request matches the following format: /swap?chain=1&fromToken=USDC&toToken=DAI&amount=100</b>\n      <h2>Available Tokens (symbols to use)</h2>\n      ").concat(tokenHtml));
        });
        return [2 /*return*/];
    });
}); });
//BridgeQuote get request
app.get('/bridge', function (req, res) { return __awaiter(void 0, void 0, void 0, function () {
    var query, fromChain, toChain, fromTokenSymbol, toTokenSymbol, fromTokenAddress, toTokenAddress, fromTokenDecimals, toTokenDecimals, amount;
    var _a, _b, _c, _d, _e, _f, _g, _h;
    return __generator(this, function (_j) {
        query = req.query;
        fromChain = query.fromChain;
        toChain = query.toChain;
        fromTokenSymbol = String(query.fromToken);
        toTokenSymbol = String(query.toToken);
        fromTokenAddress = (_b = (_a = tokens[fromTokenSymbol]) === null || _a === void 0 ? void 0 : _a.addresses) === null || _b === void 0 ? void 0 : _b[fromChain];
        toTokenAddress = (_d = (_c = tokens[toTokenSymbol]) === null || _c === void 0 ? void 0 : _c.addresses) === null || _d === void 0 ? void 0 : _d[toChain];
        fromTokenDecimals = (_f = (_e = tokens[fromTokenSymbol]) === null || _e === void 0 ? void 0 : _e.decimals) === null || _f === void 0 ? void 0 : _f[fromChain];
        toTokenDecimals = (_h = (_g = tokens[toTokenSymbol]) === null || _g === void 0 ? void 0 : _g.decimals) === null || _h === void 0 ? void 0 : _h[toChain];
        // Handle invalid params (either token symbols or chainIDs)
        // TODO: add error handling for missing params
        if (!fromTokenAddress ||
            !toTokenAddress ||
            !fromTokenDecimals ||
            !toTokenDecimals) {
            res.send("\n        <h1>Invalid Request</h1>\n        <hr/>\n        <b>Ensure that your request matches the following format: /bridge?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000</b>\n        <h2>Available Tokens (symbols to use)</h2>\n        ".concat(tokenHtml));
            return [2 /*return*/];
        }
        amount = bignumber_1.BigNumber.from(query.amount).mul(TEN.pow(fromTokenDecimals));
        // Send request w/Synapse SDK
        Synapse.bridgeQuote(Number(fromChain), Number(toChain), fromTokenAddress, toTokenAddress, bignumber_1.BigNumber.from(amount))
            .then(function (resp) {
            // Need to add some sort of execute function here
            // TODO: Router contract v2 should return the amount out with decimals for the out-out token not the out-in token (eg.nusd).
            // Add response field with adjusted maxAmountOutStr (to account for decimals)
            var payload = resp;
            payload.maxAmountOutStr = formatBNToString(resp.maxAmountOut, toTokenDecimals);
            res.json(payload);
        })
            .catch(function (err) {
            // TODO: do a better return here
            res.send("\n        <h1>Invalid Request</h1>\n        <code>".concat(err, "</code>\n        <hr/>\n        <b>Ensure that your request matches the following format: /bridge?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000</b>\n        <h2>Available Tokens (symbols to use)</h2>\n        ").concat(tokenHtml));
        });
        return [2 /*return*/];
    });
}); });
// Beginning of txInfo functions --> These return the txInfo to actually bridge
app.get('/swapTxInfo', function (req, res) { return __awaiter(void 0, void 0, void 0, function () {
    var query, chainId, fromTokenSymbol, toTokenSymbol, fromTokenAddress, toTokenAddress, fromTokenDecimals, toTokenDecimals, amount;
    var _a, _b, _c, _d, _e, _f, _g, _h;
    return __generator(this, function (_j) {
        query = req.query;
        chainId = query.chain;
        fromTokenSymbol = String(query.fromToken);
        toTokenSymbol = String(query.toToken);
        fromTokenAddress = (_b = (_a = tokens[fromTokenSymbol]) === null || _a === void 0 ? void 0 : _a.addresses) === null || _b === void 0 ? void 0 : _b[chainId];
        toTokenAddress = (_d = (_c = tokens[toTokenSymbol]) === null || _c === void 0 ? void 0 : _c.addresses) === null || _d === void 0 ? void 0 : _d[chainId];
        fromTokenDecimals = (_f = (_e = tokens[fromTokenSymbol]) === null || _e === void 0 ? void 0 : _e.decimals) === null || _f === void 0 ? void 0 : _f[chainId];
        toTokenDecimals = (_h = (_g = tokens[toTokenSymbol]) === null || _g === void 0 ? void 0 : _g.decimals) === null || _h === void 0 ? void 0 : _h[chainId];
        // Handle invalid params (either token symbols or chainIDs)
        // TODO: add error handling for missing params
        if (!fromTokenAddress ||
            !toTokenAddress ||
            !fromTokenDecimals ||
            !toTokenDecimals) {
            res.send("\n      <h1>Invalid Params</h1>\n      <hr/>\n      <b>Ensure that your request matches the following format: /swap?chain=1&fromToken=USDC&toToken=DAI&amount=100</b>\n      <h2>Available Tokens (symbols to use)</h2>\n      ".concat(tokenHtml));
            return [2 /*return*/];
        }
        amount = bignumber_1.BigNumber.from(query.amount).mul(TEN.pow(fromTokenDecimals));
        // Send request w/Synapse SDK
        Synapse.swapQuote(Number(chainId), fromTokenAddress, toTokenAddress, bignumber_1.BigNumber.from(amount))
            .then(function (resp) {
            Synapse.swap(Number(chainId), fromTokenAddress, toTokenAddress, bignumber_1.BigNumber.from(amount), resp.query).then(function (txInfo) {
                res.json(txInfo);
            });
        })
            .catch(function (err) {
            // TODO: do a better return here
            res.send("\n      <h1>Invalid Request</h1>\n      <code>".concat(err, "</code>\n      <hr/>\n      <b>Ensure that your request matches the following format: /swapTxInfo?chain=1&fromToken=USDC&toToken=DAI&amount=100</b>\n      <h2>Available Tokens (symbols to use)</h2>\n      ").concat(tokenHtml));
        });
        return [2 /*return*/];
    });
}); });
//BridgeTxInfo
app.get('/bridgeTxInfo', function (req, res) { return __awaiter(void 0, void 0, void 0, function () {
    var query, fromChain, toChain, fromTokenSymbol, toTokenSymbol, fromTokenAddress, toTokenAddress, fromTokenDecimals, toTokenDecimals, destAddress, routerAddress, amount;
    var _a, _b, _c, _d, _e, _f, _g, _h;
    return __generator(this, function (_j) {
        query = req.query;
        fromChain = query.fromChain;
        toChain = query.toChain;
        fromTokenSymbol = String(query.fromToken);
        toTokenSymbol = String(query.toToken);
        fromTokenAddress = (_b = (_a = tokens[fromTokenSymbol]) === null || _a === void 0 ? void 0 : _a.addresses) === null || _b === void 0 ? void 0 : _b[fromChain];
        toTokenAddress = (_d = (_c = tokens[toTokenSymbol]) === null || _c === void 0 ? void 0 : _c.addresses) === null || _d === void 0 ? void 0 : _d[toChain];
        fromTokenDecimals = (_f = (_e = tokens[fromTokenSymbol]) === null || _e === void 0 ? void 0 : _e.decimals) === null || _f === void 0 ? void 0 : _f[fromChain];
        toTokenDecimals = (_h = (_g = tokens[toTokenSymbol]) === null || _g === void 0 ? void 0 : _g.decimals) === null || _h === void 0 ? void 0 : _h[toChain];
        destAddress = String(query.destAddress);
        routerAddress = '0x7e7a0e201fd38d3adaa9523da6c109a07118c96a';
        // Handle invalid params (either token symbols or chainIDs)
        // TODO: add error handling for missing params
        if (!fromTokenAddress ||
            !toTokenAddress ||
            !fromTokenDecimals ||
            !toTokenDecimals) {
            res.send("\n        <h1>Invalid Request</h1>\n        <hr/>\n        <b>Ensure that your request matches the following format: /bridgeTxInfo?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000&destAddress=0xcc78d2f004c9de9694ff6a9bbdee4793d30f3842</b>\n        <h2>Available Tokens (symbols to use)</h2>\n        ".concat(tokenHtml));
            return [2 /*return*/];
        }
        amount = bignumber_1.BigNumber.from(query.amount).mul(TEN.pow(fromTokenDecimals));
        // Send request w/Synapse SDK
        Synapse.bridgeQuote(Number(fromChain), Number(toChain), fromTokenAddress, toTokenAddress, bignumber_1.BigNumber.from(amount))
            .then(function (resp) {
            Synapse.bridge(destAddress, routerAddress, Number(fromChain), Number(toChain), fromTokenAddress, bignumber_1.BigNumber.from(amount), resp.originQuery, resp.destQuery).then(function (txInfo) {
                res.json(txInfo);
            });
        })
            .catch(function (err) {
            // TODO: do a better return here
            res.send("\n        <h1>Invalid Request</h1>\n        <code>".concat(err, "</code>\n        <hr/>\n        <b>Ensure that your request matches the following format: /bridgeTxInfo?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000&destAddress=0xcc78d2f004c9de9694ff6a9bbdee4793d30f3842</b>\n        <h2>Available Tokens (symbols to use)</h2>\n        ").concat(tokenHtml));
        });
        return [2 /*return*/];
    });
}); });
exports.server = app.listen(port, function () {
    console.log("Server listening at ".concat(port));
});
var formatBNToString = function (bn, nativePrecison, decimalPlaces) {
    if (decimalPlaces === void 0) { decimalPlaces = 18; }
    var fullPrecision = (0, units_1.formatUnits)(bn, nativePrecison);
    var decimalIdx = fullPrecision.indexOf('.');
    if (decimalPlaces === undefined || decimalIdx === -1) {
        return fullPrecision;
    }
    else {
        var rawNumber = Number(fullPrecision);
        if (rawNumber === 0) {
            return rawNumber.toFixed(1);
        }
        return rawNumber.toString();
    }
};
