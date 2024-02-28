const chalk = require('chalk')

const { getAccountBalance, getAccountNonce } = require('./chain.js')
const { readWalletAddress, readWalletType } = require('./wallet.js')

const logWallet = (chainName, walletName) => {
  const walletAddr = readWalletAddress(walletName)
  const walletType = readWalletType(walletName)
  logInfo(`Wallet: ${walletAddr} [${walletName}, ${walletType}]`)
  const balance = getAccountBalance(chainName, walletAddr)
  const nonce = getAccountNonce(chainName, walletAddr)
  logInfo(`  Balance: ${balance}`)
  logInfo(`  Nonce: ${nonce}`)
}

const logSuccess = (msg) => {
  console.log(chalk.green.bold(msg))
}

const logWarning = (msg) => {
  console.log(chalk.yellow.bold(msg))
}

const logError = (msg) => {
  console.log(chalk.red.bold(msg))
}

const logInfo = (msg) => {
  console.log(chalk.cyanBright(msg))
}

module.exports = { logWallet, logSuccess, logWarning, logError, logInfo }
