const { runCommand } = require('./utils.js')

const getChainIdRPC = (rpcUrl) => {
  return runCommand(`cast chain-id --rpc-url ${rpcUrl}`)
}

const getAccountBalanceRPC = (rpcUrl, address) => {
  return runCommand(`cast balance --ether --rpc-url ${rpcUrl} ${address}`)
}

const getAccountNonceRPC = (rpcUrl, address) => {
  return runCommand(`cast nonce --rpc-url ${rpcUrl} ${address}`)
}

module.exports = {
  getChainIdRPC,
  getAccountBalanceRPC,
  getAccountNonceRPC,
}
