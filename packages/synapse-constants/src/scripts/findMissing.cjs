//WIP

const fs = require('fs')
const path = require('path')

const ts = require('typescript')

// Read and parse bridgeMap.ts
const bridgeMapPath = path.join(__dirname, '../constants/bridgeMap.ts')
const bridgeMapContent = fs.readFileSync(bridgeMapPath, 'utf8')
const bridgeMapAst = ts.createSourceFile(
  'bridgeMap.ts',
  bridgeMapContent,
  ts.ScriptTarget.Latest,
  true
)
// Extract addresses from bridgeMap.ts
const addresses = []
ts.forEachChild(bridgeMapAst, function visit(node) {
  if (ts.isObjectLiteralExpression(node)) {
    node.properties.forEach((prop) => {
      if (ts.isPropertyAssignment(prop) && ts.isObjectLiteralExpression(prop.initializer)) {
        prop.initializer.properties.forEach((innerProp) => {
          if (ts.isPropertyAssignment(innerProp)) {
            const address = innerProp.name.getText().replace(/['"`]/g, '')
            addresses.push(address)
          }
        });
      }
    });
  }
  ts.forEachChild(node, visit)
})

// Create a copy of addresses to track missing ones
const missingAddresses = [...addresses]

// Read all files in /tokens
const tokensDir = path.join(__dirname, '../constants/tokens')
const tokenFiles = fs.readdirSync(tokensDir)

// Check each file for addresses
tokenFiles.forEach((file) => {
  const filePath = path.join(tokensDir, file)
  const fileContent = fs.readFileSync(filePath, 'utf8')

  addresses.forEach((address, index) => {
    const regex = new RegExp(address, 'i'); // Create a case-insensitive regex

    if (regex.test(fileContent)) {
      // Remove the address from missingAddresses if it's found
      const missingIndex = missingAddresses.indexOf(address)
      if (missingIndex > -1) {
        missingAddresses.splice(missingIndex, 1)
      }
    }
  })
})

// Convert the array to a string, with each address on a new line
const missingAddressesStr = missingAddresses.join('\n')

// Write the string to a file
fs.writeFileSync('unsupportedTokens.txt', missingAddressesStr)
console.log('Missing Token Addresses can be found in unsupportedTokens.txt')
