//WIP

const fs = require('fs')
const path = require('path')

const ts = require('typescript')

// Read and parse bridgeMap.ts
const bridgeMapPath = path.join(__dirname, 'constants/tokens/bridgeMap.ts')
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
  if (ts.isPropertyAssignment(node) && node.name.getText() === 'addresses') {
    const initializer = node.initializer;
    if (ts.isObjectLiteralExpression(initializer)) {
      initializer.properties.forEach((prop) => {
        if (ts.isPropertyAssignment(prop)) {
          addresses.push(prop.initializer.getText().replace(/['"`]/g, ''));
        }
      });
    }
  }
  ts.forEachChild(node, visit);
})

// Read all files in /tokens
const tokensDir = path.join(__dirname, 'constants/tokens')
const tokenFiles = fs.readdirSync(tokensDir)

// Check each file for addresses
tokenFiles.forEach((file) => {
  const filePath = path.join(tokensDir, file)
  const fileContent = fs.readFileSync(filePath, 'utf8')

  addresses.forEach((address) => {
    if (fileContent.includes(address)) {
      console.log(`Address ${address} found in file ${file}`)
    }
  })
})
