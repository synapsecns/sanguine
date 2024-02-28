const fs = require('fs')

const { assertCondition } = require('./utils.js')

const DEVOPS_CONFIG_PATH = './devops.json'

let cachedConfig = null

const readConfig = () => {
  if (!cachedConfig) {
    const configContent = fs.readFileSync(DEVOPS_CONFIG_PATH)
    cachedConfig = JSON.parse(configContent)
  }
  return cachedConfig
}

const readValueOrMissingKey = (...keys) => {
  let config = readConfig()
  let accKey = ''
  for (const key of keys) {
    accKey += `.${key}`
    if (key in config) {
      config = config[key]
    } else {
      return {
        missingKey: accKey,
        value: undefined,
      }
    }
  }
  return {
    missingKey: undefined,
    value: config,
  }
}

/**
 * Try to read the deep value of the configuration object.
 * For example, tryReadConfigValue('chains', 'mainnet') will return the value of `config.chains.mainnet`.
 * If the key is not found, returns undefined.
 *
 * @param  {...string} keys - The keys to read
 * @returns {string|undefined} The value of the configuration object, or undefined if the key is not found
 */
const tryReadConfigValue = (...keys) => {
  const { value } = readValueOrMissingKey(...keys)
  return value
}

/**
 * Reads the deep value of the configuration object.
 * For example, readConfigValue('chains', 'mainnet') will return the value of `config.chains.mainnet`.
 *
 * @param  {...string} keys - The keys to read
 * @returns {string} The value of the configuration object
 */
const readConfigValue = (...keys) => {
  const { missingKey, value } = readValueOrMissingKey(...keys)
  assertCondition(
    missingKey === undefined,
    `Key ${missingKey} is not found in the configuration`
  )
  return value
}

module.exports = { tryReadConfigValue, readConfigValue }
