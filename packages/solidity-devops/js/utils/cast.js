const { getCommandOutput } = require('./utils.js')

const getChainIdRPC = (rpcUrl) => {
  return getCommandOutput(`cast chain-id --rpc-url ${rpcUrl}`)
}

const getAccountBalanceRPC = (rpcUrl, address) => {
  return getCommandOutput(`cast balance --ether --rpc-url ${rpcUrl} ${address}`)
}

const getAccountNonceRPC = (rpcUrl, address) => {
  return getCommandOutput(`cast nonce --rpc-url ${rpcUrl} ${address}`)
}

module.exports = {
  getChainIdRPC,
  getAccountBalanceRPC,
  getAccountNonceRPC,
}
