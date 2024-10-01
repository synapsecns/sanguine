const fs = require('fs')
const { execSync } = require('child_process')

// Writes map export to a TypeScript file, then runs prettier on the file
const prettyPrintTS = (map, mapName, fn) => {
  console.log(`Writing ${mapName} to ${fn}`)
  const json = JSON.stringify(map)
  fs.writeFileSync(fn, `export const ${mapName} = ${json}`)
  // Run prettier on the file using terminal command:
  execSync(`npx prettier --write ${fn}`)
}

module.exports = { prettyPrintTS }
