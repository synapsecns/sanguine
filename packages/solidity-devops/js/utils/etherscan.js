const { VERIFIER_ETHERSCAN } = require('./chain.js')
const { readEnv } = require('./env.js')
const { getCommandOutput, runCommand, syncSleep } = require('./utils.js')

const MAX_ATTEMPTS = 10

const getResponseResult = (response) => {
  if (!response.result) {
    logError(
      `Result not found in verification response: ${JSON.stringify(
        response,
        null,
        2
      )}`
    )
    return null
  }
  return response.result
}

/**
 * Initiates the process of proxy verification on Etherscan and returns the GUID of the verification request.
 *
 * @param {string} chainName - The name of the chain to verify on
 * @param {string} address - The address of the proxy contract to verify
 * @returns {string} The GUID of the verification request
 */
const initiateVerifyProxy = (chainName, address) => {
  const verifier = readEnv(chainName, 'VERIFIER')
  if (verifier !== VERIFIER_ETHERSCAN) {
    logError(`Verifier ${verifier} is not supported by this command`)
    return null
  }
  const url = readEnv(chainName, 'VERIFIER_URL')
  const key = readEnv(chainName, 'VERIFIER_KEY')
  const params = `module=contract&action=verifyproxycontract&apikey=${key}&address=${address}`
  const response = JSON.parse(
    getCommandOutput(`curl -s -d "${params}" "${url}"`)
  )
  return getResponseResult(response)
}

const getRequestStatus = (chainName, actionName, guid) => {
  const verifier = readEnv(chainName, 'VERIFIER')
  if (verifier !== VERIFIER_ETHERSCAN) {
    logError(`Verifier ${verifier} is not supported by this command`)
    return null
  }
  const url = readEnv(chainName, 'VERIFIER_URL')
  const key = readEnv(chainName, 'VERIFIER_KEY')
  const params = `module=contract&action=${actionName}&apikey=${key}&guid=${guid}`
  for (let i = 0; i < MAX_ATTEMPTS; i++) {
    syncSleep(5, `before attempt ${i + 1} to ${actionName}`)
    const response = JSON.parse(
      getCommandOutput(`curl -s -d "${params}" "${url}"`)
    )
    const result = getResponseResult(response)
    if (!result) {
      return null
    }
    if (result === 'Pending in queue') {
      logInfo(`  ${result}...`)
      continue
    }
    return response
  }
  logError(
    `Request with GUID ${guid} not completed after ${MAX_ATTEMPTS} attempts. Consider checking the network status or increasing the retry interval.`
  )
  return null
}

module.exports = {
  initiateVerifyProxy,
  getRequestStatus,
}
