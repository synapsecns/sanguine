const fs = require('fs')
const { execSync } = require('child_process')

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
 * Creates a directory recursively if it doesn't exist.
 * No-op if the directory already exists.
 *
 * @param {string[]} dirNames - The names of the directories to create
 * @returns {string} The path of the directory (whether it was created or not)
 */
const createDir = (...dirNames) => {
  const dirPath = dirNames.join('/')
  if (!fs.existsSync(dirPath)) {
    fs.mkdirSync(dirPath, { recursive: true })
  }
  return dirPath
}

/**
 * Runs a command and returns its output.
 * If the command fails, exits the process.
 *
 * @param {string} command - The command to run
 * @returns {string} The output of the command
 */
const getCommandOutput = (command) => {
  try {
    const output = execSync(command)
    return output.toString().trim()
  } catch (error) {
    process.exit(1)
  }
}

/**
 * Runs a command (printing its output to the console), and exits the process if it fails.
 *
 * @param {string} command - The command to run
 */
const runCommand = (command) => {
  try {
    execSync(command, { stdio: 'inherit' })
  } catch (error) {
    process.exit(1)
  }
}

module.exports = {
  assertCondition,
  createDir,
  exitWithError,
  getCommandOutput,
  runCommand,
}
