const { runCommand } = require('./utils.js')

const getChainId = (rpcUrl) => {
  return runCommand(`cast chain-id --rpc-url ${rpcUrl}`)
}

module.exports = { getChainId }
