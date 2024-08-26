const fs = require('fs')
const path = require('path')

// Tokens to exclude from the check (tokens that are in bridgeable but not exported)
const excludedTokens = ['USDB', 'SPECTRAL']

const packageExportContent = fs.readFileSync(
  path.join(__dirname, '../../src', 'index.tsx'),
  'utf8'
)
const bridgeableFileContent = fs.readFileSync(
  path.join(__dirname, '../../src', 'constants', 'bridgeable.ts'),
  'utf8'
)

// Extract token names from bridgeable file
const bridgeableTokens = bridgeableFileContent
  .match(/export const (\w+):/g)
  .map((match) => match.split(' ')[2].replace(':', ''))

// Extract exported token names from main package export file
const exportedTokens = packageExportContent
  .match(/export const (\w+) = BRIDGEABLE\.\1/g)
  .map((match) => match.split(' ')[2])

// Find tokens that are in BRIDGEABLE but not exported in the main file, excluding specified tokens
const missingExports = bridgeableTokens.filter(
  (token) => !exportedTokens.includes(token) && !excludedTokens.includes(token)
)

if (missingExports.length > 0) {
  console.log(
    'The following tokens are in bridgeable.ts but not exported in the main file:'
  )
  missingExports.forEach((token) => console.log(token))
  process.exit(1)
} else {
  console.log(
    'All required tokens from bridgeable.ts are correctly exported in the main file.'
  )
  process.exit(0)
}
