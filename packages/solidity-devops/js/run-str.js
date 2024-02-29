#!/usr/bin/env node
const { parseCommandLineArgs, addOptions } = require('./utils/options.js')
const { runCommand } = require('./utils/utils.js')

const { positionalArgs, options } = parseCommandLineArgs({
  requiredArgsCount: 4,
  usage:
    'Usage: "yarn sol-run-str <path-to-script> <chain-name> <wallet-name> <string-arg> [<options>]"',
})

const [scriptFN, chainName, walletName, stringArg] = positionalArgs

const newOptions = addOptions(options, `--sig "run(string)" "${stringArg}"`)
// launch run.js with the new options
runCommand(`yarn sol-run ${scriptFN} ${chainName} ${walletName} ${newOptions}`)
