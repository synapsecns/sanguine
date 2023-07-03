import { JsonRpcProvider } from '@ethersproject/providers'
import { hexZeroPad } from '@ethersproject/bytes'
import { GETLOGS_SIZE } from '@constants/bridgeWatcher'
import { Contract } from '@ethersproject/contracts'
import { toHexStr } from '@utils/toHexStr'
import { BridgeWatcherTx } from '@types'
import { id } from '@ethersproject/hash'
import { TOKEN_HASH_MAP } from '@constants/tokens'
import { getAddress, isAddress } from '@ethersproject/address'
import * as CHAINS from '@constants/chains/master'
import { WETH } from '@constants/tokens/swapMaster'
import {
  SYN,
  NUSD,
  NETH,
  FRAX,
  SYN_FRAX,
  WBTC,
  KLAYTN_WETH,
  DOG,
  LINK,
  GOHM,
  HIGHSTREET,
  JUMP,
  NFD,
  NEWO,
  VSTA,
  GMX,
  SDT,
  UNIDX,
  USDB,
  SFI,
  H2O,
  L2DAO,
  PLS,
  AGEUR,
  NOTE,
  USDC
} from '@constants/tokens/master'
import { parse } from 'path'
export const getTransactionReceipt = async (
  txHash: string,
  provider: JsonRpcProvider
) => {
  const receipt = await provider.getTransactionReceipt(txHash)
  return receipt
}
export const getBlock = async (
  blockNumber: number,
  provider: JsonRpcProvider
) => {
  const block = await provider.getBlock(blockNumber)
  return block
}
export const getLogs = async (
  currentBlock: number,
  provider: JsonRpcProvider,
  contract: Contract,
  address: string
) => {
  const filter = {
    address: contract.address,
    topics: [null, hexZeroPad(address, 32)],
    fromBlock: toHexStr(currentBlock - GETLOGS_SIZE),
    toBlock: toHexStr(currentBlock),
  }
  try {
    const logs = await provider.send('eth_getLogs', [filter])
    return logs
  } catch (e) {
    console.log('getLogs error', e)
    return []
  }
}

export const checkTxIn = (tx) => {
  return tx?.chainId ? true : false
}

export const generateBridgeTx = (
  isFrom,
  address,
  chainId,
  parsedLog,
  timestampObj,
  txReceipt,
  destinationAddress
): BridgeWatcherTx => {
  const swapTokenAddr = getAddress(parsedLog.token)

  let tokenAddr
  if (isFrom) {
    if (txReceipt.logs[1].address === GMX.addresses[CHAINS.AVALANCHE.id]) {
      tokenAddr = GMX.addresses[CHAINS.AVALANCHE.id]
    } else {
      tokenAddr = txReceipt.logs[0].address
    }
  } else {
    if (
      [
        SYN,
        LINK,
        HIGHSTREET,
        DOG,
        JUMP,
        FRAX,
        NFD,
        GOHM,
        AGEUR,
        H2O,
        L2DAO,
        PLS,
        NEWO,
        VSTA,
        SFI,
        SDT,
        UNIDX,
        USDB,
        GMX,
        WBTC,
        KLAYTN_WETH,
        NOTE,
      ]
        .map((t) => t.addresses[chainId])
        .includes(swapTokenAddr)
    ) {
      tokenAddr = TOKEN_HASH_MAP[chainId][swapTokenAddr].addresses[chainId]
    } else if (swapTokenAddr === SYN_FRAX.addresses[chainId]) {
      tokenAddr = FRAX.addresses[chainId]
    } else if (swapTokenAddr === GMX.wrapperAddresses[chainId]) {
      tokenAddr = GMX.addresses[chainId]
    } else if (swapTokenAddr === NETH.addresses[chainId]) {
      tokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
    } else if (swapTokenAddr === WETH.addresses[chainId]) {
      if (chainId === CHAINS.ETH.id) {
        tokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
      } else {
        tokenAddr = txReceipt.logs[txReceipt.logs.length - 1].address
      }
    } else if (swapTokenAddr === NUSD.addresses[chainId]) {
      if (chainId === CHAINS.ETH.id) {
        if (parsedLog.event === 'TokenWithdraw') {
          tokenAddr = txReceipt.logs[txReceipt.logs.length - 1].address
        } else {
          tokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
        }
      } else if (chainId === CHAINS.POLYGON.id) {
        tokenAddr = txReceipt.logs[txReceipt.logs.length - 3].address
      } else {
        tokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
      }
    } else if(!isFrom && swapTokenAddr === USDC.addresses[chainId] && [CHAINS.ARBITRUM.id, CHAINS.ETH.id, CHAINS.AVALANCHE.id].includes(chainId)) {
      tokenAddr = txReceipt.logs[txReceipt.logs.length - 3].address
    } else {
      tokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
    }
  }
  const token = TOKEN_HASH_MAP[chainId][tokenAddr]

  let inputTokenAmount
  if (
    getAddress(txReceipt.logs[0]?.address) ===
      GMX.addresses[CHAINS.ARBITRUM.id] ||
    getAddress(txReceipt.logs[1]?.address) ===
      GMX.addresses[CHAINS.AVALANCHE.id]
  ) {
    inputTokenAmount = txReceipt.logs[1].data
  } else {
    inputTokenAmount = txReceipt.logs[0].data
  }

  return {
    isFrom,
    amount: isFrom ? inputTokenAmount : parsedLog.amount,
    timestamp: timestampObj.timestamp,
    blockNumber: parsedLog.blockNumber,
    chainId,
    address,
    txHash: txReceipt.transactionHash,
    txReceipt,
    token,
    kappa: parsedLog.requestID ? parsedLog.requestID : removePrefix(id(parsedLog.transactionHash)),
    toChainId: isFrom ? Number(parsedLog.chainId.toString()) : chainId,
    toAddress: isAddress(destinationAddress) ? destinationAddress : address,
    contractEmittedFrom: parsedLog.contractEmittedFrom,
  }
}

export const getHighestBlock = async (
  chainId: number,
  provider: JsonRpcProvider
) => {
  const highestBlock = await provider.getBlockNumber()
  return highestBlock
}

const removePrefix = (str: string): string => {
  if (str.startsWith('0x')) {
    return str.substring(2)
  }
  return str
}
