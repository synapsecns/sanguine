#!/usr/bin/env node

/**
 * Copyright (c) 2021-present, Matti Bar-Zeev.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

const { fork } = require('child_process')
const path = require('path')
const COLLECT_FILES = 'collectFiles'

const execute = ({ command, commandArgs }) => {
  let scriptPath
  switch (command) {
    case COLLECT_FILES:
      scriptPath = path.resolve(__dirname, '../lib/coverage-aggregator.js')
      break
  }

  fork(scriptPath, commandArgs)
}

const args = process.argv.slice(2)
const command = args[0]
const commandArgs = args.slice(1)

execute({ command, commandArgs })

module.exports = {
  execute,
  COLLECT_FILES,
}
