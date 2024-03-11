const { assertCondition } = require('./utils')

/**
 * Parses the command line arguments and returns an object with the positional arguments and options.
 * Positional arguments are the first `requiredArgsCount` arguments, and options are the rest.
 * If there are no options, the options string is empty.
 *
 * @param {Object} options - The options object.
 * @param {number} options.requiredArgsCount - The number of required arguments.
 * @param {string} options.usage - The usage information string.
 * @returns {{ positionalArgs: string[], options: string }} The positional arguments and options.
 */
const parseCommandLineArgs = ({ requiredArgsCount, usage }) => {
  // Remove the first two elements (node and script path)
  const args = process.argv.slice(2)
  assertCondition(args.length >= requiredArgsCount, usage)
  const positionalArgs = args.slice(0, requiredArgsCount)
  // Join the rest of the args into the options string, wrapping them in quotes
  // Do not wrap the options that start with '--'
  const options = args
    .slice(requiredArgsCount)
    .map((arg) => (arg.startsWith('--') ? arg : `"${arg}"`))
    .join(' ')
  return { positionalArgs, options }
}

const isBroadcasted = (options) => {
  return options.includes('--broadcast')
}

const addVerifyOptions = (options) => {
  return options.includes('--verify') ? options : `${options} --verify`
}

const addOptions = (options, newOptions) => {
  return newOptions ? `${options} ${newOptions}` : options
}

module.exports = {
  parseCommandLineArgs,
  isBroadcasted,
  addVerifyOptions,
  addOptions,
}
