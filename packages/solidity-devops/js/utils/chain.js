const {
  getChainIdRPC,
  getChainGasPricingRPC,
  getAccountBalanceRPC,
  getAccountNonceRPC,
  hasCodeRPC,
} = require('./cast.js')
const { tryReadConfigValue } = require('./config.js')
const { readEnv } = require('./env.js')
const { logError, logInfo } = require('./logger.js')
const { readWalletAddress, readWalletType } = require('./wallet.js')

const OPTION_AUTO_FILL_GAS_PRICE_LEGACY = '--auto-gas-legacy'
const OPTION_AUTO_FILL_GAS_PRICE_1559 = '--auto-gas-1559'

const VERIFIER_ETHERSCAN = 'etherscan'
const VERIFIER_BLOCKSCOUT = 'blockscout'
const VERIFIER_SOURCIFY = 'sourcify'

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
  const options = tryReadConfigValue('chains', chainName) || ''
  return applyAutoFillGasPrice(chainName, options)
}

const readChainVerificationOptions = (chainName) => {
  const verifier = readEnv(chainName, 'VERIFIER')
  switch (verifier) {
    case VERIFIER_ETHERSCAN:
      return readEtherscanOptions(chainName)
    case VERIFIER_BLOCKSCOUT:
      return readBlockscoutOptions(chainName)
    case VERIFIER_SOURCIFY:
      return readSourcifyOptions(chainName)
    default:
      logError(`Unknown verifier: ${verifier}`)
      return null
  }
}

const readEtherscanOptions = (chainName) => {
  const url = readEnv(chainName, 'VERIFIER_URL')
  const key = readEnv(chainName, 'VERIFIER_KEY')
  return `--verifier etherscan --verifier-url ${url} --etherscan-api-key ${key}`
}

const readBlockscoutOptions = (chainName) => {
  const url = readEnv(chainName, 'VERIFIER_URL')
  return `--verifier blockscout --verifier-url ${url}`
}

const readSourcifyOptions = (chainName) => {
  return '--verifier sourcify'
}

/**
 * Fetches the gas price from the chain's RPC and updates the options, if the auto-fill gas price option is present.
 *
 * @param {string} chainName - The name of the chain
 * @param {string} options - The chain specific options
 * @returns {string} The chain specific options with the gas price filled in, if the auto-fill gas price option is present
 */
const applyAutoFillGasPrice = (chainName, options) => {
  if (options.includes(OPTION_AUTO_FILL_GAS_PRICE_LEGACY)) {
    const { gasPrice } = getChainGasPricing(chainName)
    return options.replace(
      OPTION_AUTO_FILL_GAS_PRICE_LEGACY,
      `--with-gas-price ${gasPrice}`
    )
  } else if (options.includes(OPTION_AUTO_FILL_GAS_PRICE_1559)) {
    const { baseFee, gasPrice } = getChainGasPricing(chainName)
    if (baseFee > gasPrice) {
      logError(
        `Base fee (${baseFee}) is greater than gas price (${gasPrice}), using default gas price instead`
      )
      return options.replace(OPTION_AUTO_FILL_GAS_PRICE_1559, '')
    }
    const priorityFee = gasPrice - baseFee
    // Use 2*base + priority as the max gas price
    const maxGasPrice = 2 * baseFee + priorityFee
    return options.replace(
      OPTION_AUTO_FILL_GAS_PRICE_1559,
      `--with-gas-price ${maxGasPrice} --priority-gas-price ${priorityFee}`
    )
  }
  return options
}

const getChainId = (chainName) => {
  return getChainIdRPC(readChainRPC(chainName))
}

const getChainGasPricing = (chainName) => {
  return getChainGasPricingRPC(readChainRPC(chainName))
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
  readChainVerificationOptions,
  getChainId,
  getChainGasPricing,
  getAccountBalance,
  getAccountNonce,
  hasCode,
  logWallet,
  VERIFIER_ETHERSCAN,
  VERIFIER_BLOCKSCOUT,
  VERIFIER_SOURCIFY,
}
