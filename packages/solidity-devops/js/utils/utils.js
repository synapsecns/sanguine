/**
 * Asserts that a condition is true. If not, logs an error message and exits the process.
 *
 * @param {bool} condition - The condition to assert
 * @param {string} message - The error message to log if the condition is false
 */
const assertCondition = (condition, message) => {
  if (!condition) {
    console.error(message)
    process.exit(1)
  }
}

/**
 * Logs an error message and exits the process.
 *
 * @param {string} message - The error message to log
 */
const exitWithError = (message) => {
  assertCondition(false, message)
}

/**
 * Parses the command line arguments and returns an object with the positional arguments and options.
 * Positional arguments are the arguments before the '--' separator, and options are the arguments after it.
 *
 * @returns The positional arguments and options
 */
const parseCommandLineArgs = () => {
  // Remove the first two elements (node and script path)
  const args = process.argv.slice(2)
  // Find the index of '--' which separates positional arguments from options
  const dashIndex = args.indexOf('--')
  let positionalArgs
  let options

  if (dashIndex !== -1) {
    // If '--' is present, separate the arguments before and after it
    positionalArgs = args.slice(0, dashIndex)
    options = args.slice(dashIndex + 1).join(' ') // Join the options back into a single string
  } else {
    // If '--' is not present, all arguments are considered positional
    positionalArgs = args
    options = '' // No options provided
  }

  return { positionalArgs, options }
}

module.exports = { assertCondition, exitWithError, parseCommandLineArgs }
