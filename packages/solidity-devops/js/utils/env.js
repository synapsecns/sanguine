const { assertCondition } = require('./utils.js')

// Read environment variables into process.env
const loadEnv = () => {
  require('dotenv').config()
}

/**
 * Merge the variable names into a single name, and convert it to uppercase.
 * E.g. mergeVarNames('foo', 'bar') will return 'FOO_BAR'.
 *
 * @param  {...string} varNames - The names of the environment variables to merge
 * @returns The merged variable name
 */
const mergeVarNames = (...varNames) => {
  return varNames.map((name) => name.toUpperCase()).join('_')
}

/**
 * Tries to read the value of an environment variable. If the variable is not set, returns undefined.
 * The variable names are case-insensitive and are concatenated with underscores.
 * E.g. mergeAndTryReadEnv('foo', 'bar') will try to read the value of process.env.FOO_BAR.
 *
 * @param  {...string} varNames - The names of the environment variables to read
 * @returns The merged variable name and value of the environment variable
 */
const mergeAndTryReadEnv = (...varNames) => {
  const varName = mergeVarNames(...varNames)
  const value = process.env[varName]
  return { varName, value }
}

/**
 * Tries to read the value of an environment variable. If the variable is not set, returns undefined.
 * The variable names are case-insensitive and are concatenated with underscores.
 * E.g. tryReadEnv('foo', 'bar') will try to read the value of process.env.FOO_BAR.
 *
 * @param  {...string} varNames - The names of the environment variables to read
 * @returns The value of the environment variable, or undefined if it is not set
 */
const tryReadEnv = (...varNames) => {
  const { value } = mergeAndTryReadEnv(...varNames)
  return value
}

/**
 * Reads the value of an environment variable. If the variable is not set, logs an error message and exits the process.
 * The variable names are case-insensitive and are concatenated with underscores.
 * E.g. readEnv('foo', 'bar') will try to read the value of process.env.FOO_BAR.
 *
 * @param  {...string} varNames - The names of the environment variables to read
 * @returns {string} The value of the environment variable
 * @throws If the environment variable is not set
 */
const readEnv = (...varNames) => {
  const { varName, value } = mergeAndTryReadEnv(...varNames)
  assertCondition(
    value !== undefined,
    `Environment variable ${varName} is not set`
  )
  return value
}

module.exports = { loadEnv, mergeVarNames, tryReadEnv, readEnv }
