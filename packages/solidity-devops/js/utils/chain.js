const {
  getChainIdRPC,
  getAccountBalanceRPC,
  getAccountNonceRPC,
  hasCodeRPC,
} = require('./cast.js')
const { tryReadConfigValue } = require('./config.js')
const { readEnv } = require('./env.js')
const { logInfo } = require('./logger.js')
const { readWalletAddress, readWalletType } = require('./wallet.js')

/**
 * Reads the URL of the chain's RPC from the environment variables.
 *
 * @param {string} chainName - The name of the chain
 * @returns {string} The URL of the chain's RPC
 */
const readChainRPC = (chainName) => {
  return readEnv(chainName, 'RPC')
}

/**
 * Reads chain specific options from the devops configuration.
 * If no options are found, returns an empty string.
 *
 * @param {string} chainName - The name of the chain
 * @returns {string} The chain specific options
 */
const readChainSpecificOptions = (chainName) => {
  const options = tryReadConfigValue('chains', chainName)
  return options || ''
}

const getChainId = (chainName) => {
  return getChainIdRPC(readChainRPC(chainName))
}

const getAccountBalance = (chainName, address) => {
  return getAccountBalanceRPC(readChainRPC(chainName), address)
}

const getAccountNonce = (chainName, address) => {
  return getAccountNonceRPC(readChainRPC(chainName), address)
}

const hasCode = (chainName, address) => {
  return hasCodeRPC(readChainRPC(chainName), address)
}

const logWallet = (chainName, walletName) => {
  const walletAddr = readWalletAddress(walletName)
  const walletType = readWalletType(walletName)
  logInfo(`Wallet: ${walletAddr} [${walletName}, ${walletType}]`)
  const balance = getAccountBalance(chainName, walletAddr)
  const nonce = getAccountNonce(chainName, walletAddr)
  logInfo(`  Balance: ${balance}`)
  logInfo(`  Nonce: ${nonce}`)
}

module.exports = {
  readChainRPC,
  readChainSpecificOptions,
  getChainId,
  getAccountBalance,
  getAccountNonce,
  hasCode,
  logWallet,
}
