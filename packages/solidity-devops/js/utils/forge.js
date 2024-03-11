const { runCommand } = require('./utils.js')

const forgeScript = (scriptFN, options) => {
  return runCommand(`forge script ${scriptFN} ${options}`)
}

const forgeVerify = (options) => {
  return runCommand(`forge verify-contract ${options}`)
}

module.exports = {
  forgeScript,
  forgeVerify,
}
