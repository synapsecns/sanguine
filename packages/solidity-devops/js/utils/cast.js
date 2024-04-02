const { getCommandOutput } = require('./utils.js')

const getChainIdRPC = (rpcUrl) => {
  return getCommandOutput(`cast chain-id --rpc-url ${rpcUrl}`)
}

const getChainGasPricingRPC = (rpcUrl) => {
  const baseFee = getCommandOutput(`cast base-fee --rpc-url ${rpcUrl}`)
  const gasPrice = getCommandOutput(`cast gas-price --rpc-url ${rpcUrl}`)
  return {
    baseFee,
    gasPrice,
  }
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
  getChainGasPricingRPC,
  getAccountBalanceRPC,
  getAccountNonceRPC,
  hasCodeRPC,
}
