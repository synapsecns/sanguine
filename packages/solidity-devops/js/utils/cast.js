const { logError } = require('./logger.js')
const { getCommandOutput } = require('./utils.js')

const getChainIdRPC = (rpcUrl) => {
  return getCommandOutput(`cast chain-id --rpc-url ${rpcUrl}`)
}

const getChainGasPriceRPC = (rpcUrl) => {
  const output = getCommandOutput(
    `cast gas-price --rpc-url ${rpcUrl}`,
    (exitOnError = false)
  )
  if (!output) {
    logError('  Failed to get gas price from the chain')
    process.exit(1)
  }
  // Output is returned without quotes
  return Number(output)
}

const getChainMaxPriorityFeeRPC = (rpcUrl) => {
  const output = getCommandOutput(
    `cast rpc --rpc-url ${rpcUrl} eth_maxPriorityFeePerGas`,
    (exitOnError = false)
  )
  if (!output) {
    logError('  EIP-1559 is not supported on this chain')
    process.exit(1)
  }
  // Remove quotes and convert from hex into decimal
  return Number(output.replace(/"/g, ''))
}

const getAccountBalanceRPC = (rpcUrl, address) => {
  return getCommandOutput(`cast balance --ether --rpc-url ${rpcUrl} ${address}`)
}

const getAccountNonceRPC = (rpcUrl, address) => {
  return getCommandOutput(`cast nonce --rpc-url ${rpcUrl} ${address}`)
}

const hasCodeRPC = (rpcUrl, address) => {
  const code = getCommandOutput(`cast code --rpc-url ${rpcUrl} ${address}`)
  // 0x is returned for an address without code
  return code.length > 2
}

module.exports = {
  getChainIdRPC,
  getChainGasPriceRPC,
  getChainMaxPriorityFeeRPC,
  getAccountBalanceRPC,
  getAccountNonceRPC,
  hasCodeRPC,
}
