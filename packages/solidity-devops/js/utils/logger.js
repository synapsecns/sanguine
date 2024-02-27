const { getAccountBalance, getAccountNonce } = require('./chain.js')
const { readWalletAddress, readWalletType } = require('./wallet.js')

const logWallet = (chainName, walletName) => {
  const walletAddr = readWalletAddress(walletName)
  const walletType = readWalletType(walletName)
  console.log(`Wallet ${walletAddr} [${walletName}, ${walletType}]`)
  const balance = getAccountBalance(chainName, walletAddr)
  const nonce = getAccountNonce(chainName, walletAddr)
  console.log(`  Balance: ${balance}`)
  console.log(`  Nonce: ${nonce}`)
}

module.exports = { logWallet }
