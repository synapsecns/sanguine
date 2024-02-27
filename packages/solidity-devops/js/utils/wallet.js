const { readEnv } = require('./env.js')

const TYPE_KEYSTORE = 'keystore'
const TYPE_LEDGER = 'ledger'
const TYPE_PRIVATE_KEY = 'pk'
const TYPE_TREZOR = 'trezor'

/**
 * Read the wallet configuration and return the command line options for the wallet.
 *
 * @param {string} walletName - The name of the wallet
 * @returns The command line options for the wallet, could be used with `forge` commands
 */
const readWalletOptions = (walletName) => {
  const walletAddr = readWalletAddress(walletName)
  const walletOptions = readWallet(walletName)
  return `${walletOptions} --sender ${walletAddr}`
}

/**
 * Read the wallet address from the environment variables.
 *
 * @param {string} walletName - The name of the wallet
 * @returns The address of the wallet
 */
const readWalletAddress = (walletName) => {
  return readEnv(walletName, 'ADDR')
}

/**
 * Read the wallet type from the environment variables.
 *
 * @param {string} walletName - The name of the wallet
 * @returns The type of the wallet
 */
const readWalletType = (walletName) => {
  return readEnv(walletName, 'TYPE').toLowerCase()
}

const readWallet = (walletName) => {
  const walletType = readWalletType(walletName)
  switch (walletType) {
    case TYPE_KEYSTORE:
      return readKeystoreOptions(walletName)
    case TYPE_LEDGER:
      return readLedgerOptions(walletName)
    case TYPE_PRIVATE_KEY:
      return readPrivateKeyOptions(walletName)
    case TYPE_TREZOR:
      return readTrezorOptions(walletName)
    default:
      return readInteractivePromptOptions(walletName)
  }
}

const readKeystoreOptions = (walletName) => {
  const walletJson = readEnv(walletName, 'JSON')
  return `--keystore ${walletJson}`
}

const readLedgerOptions = (walletName) => {
  return '--ledger'
}

const readPrivateKeyOptions = (walletName) => {
  const walletPK = readEnv(walletName, 'PK')
  return `--private-key ${walletPK}`
}

const readTrezorOptions = (walletName) => {
  return '--trezor'
}

const readInteractivePromptOptions = (walletName) => {
  // Use interactive prompt for private key as the last resort
  return '-i 1'
}

module.exports = { readWalletOptions, readWalletAddress, readWalletType }
