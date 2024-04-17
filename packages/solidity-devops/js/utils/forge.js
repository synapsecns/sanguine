const { runCommand } = require('./utils.js')

const forgeScript = (scriptFN, options, exitOnError = false) => {
  return runCommand(`forge script ${scriptFN} ${options}`, exitOnError)
}

const forgeVerify = (options, exitOnError = false) => {
  return runCommand(`forge verify-contract ${options}`, exitOnError)
}

module.exports = {
  forgeScript,
  forgeVerify,
}
