#!/usr/bin/env node

const path = require('path')
const fs = require('fs')

const yargs = require('yargs/yargs')
const glob = require('glob')

const GREEN = '\x1b[32m%s\x1b[0m'

const collectFiles = ({ pattern, target }) => {
  if (!pattern || !target) {
    throw new Error('Missing either pattern or target params')
  }

  console.log(GREEN, `Collecting files... into ${target}`)

  glob(pattern, {}, (err, files) => {
    if (err) {
      throw err
    }
    files.forEach((file, index) => {
      if (!file.includes('node_modules')) {
        fs.copyFileSync(
          file,
          path.resolve(target, `${index}-${path.basename(file)}`)
        )
      }
    })
  })

  console.log(GREEN, `Done.`)
}

const args = yargs(process.argv.slice(2)).argv

collectFiles(args)
