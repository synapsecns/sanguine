const chalk = require('chalk')

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
  console.log(chalk.cyanBright.bold(msg))
}

const logCommand = (command) => {
  console.log(chalk.blue(redactKeys(command)))
}

const redactKeys = (command) => {
  // Find all options that end with -key and redact the following value
  const keyRegex = /(--\S+-key) (\S+)/g
  const redactedCommand = command.replace(keyRegex, '$1 <REDACTED-VALUE>')
  return redactedCommand
}

module.exports = { logSuccess, logWarning, logError, logInfo, logCommand }
