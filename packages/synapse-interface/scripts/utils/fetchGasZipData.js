const GAS_ZIP_URL = 'https://backend.gas.zip/v2/chains'
// TODO: filter out chains without direct deposit support
// https://dev.gas.zip/gas/chain-support/deposit#direct-deposit
const fetchGasZipData = async () => {
  const response = await fetch(GAS_ZIP_URL)
  const data = await response.json()
  return data.chains.map((chain) => chain.chain)
}

module.exports = { fetchGasZipData }
