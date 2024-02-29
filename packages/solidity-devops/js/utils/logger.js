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
  console.log(chalk.blue(command))
}

module.exports = { logSuccess, logWarning, logError, logInfo, logCommand }
