import { execSync } from 'child_process'
import fs from 'fs'
import path from 'path'

import { BRIDGE_MAP } from '../constants/bridgeMap'

// Writes map export to a TypeScript file, then runs prettier on the file
const prettyPrintTS = (map: any, mapName: string, fn: string) => {
  console.log(`Writing ${mapName} to ${fn}`)
  const json = JSON.stringify(map)
  fs.writeFileSync(fn, `export const ${mapName} = ${json}`)
  // Run prettier on the file using terminal command:
  execSync(`npx prettier --write ${fn}`)
}

// Define types based on the structure of BRIDGE_MAP
type TokenInfo = {
  decimals: number
  symbol: string
  origin: string[]
  destination: string[]
  swappable: string[]
}

type TokensByAddress = {
  [tokenAddress: string]: TokenInfo
}

type ChainTokenMap = {
  [chainId: string]: TokensByAddress
}

const bridgeMapObject = BRIDGE_MAP as ChainTokenMap

// Adjust the bridgeMap by adding RFQ symbols to origin arrays
Object.entries(bridgeMapObject).forEach(([_, tokens]) => {
  // Find all RFQ.* symbols in tokens' origin arrays
  const rfqSymbols = new Set<string>()

  Object.values(tokens).forEach((token: TokenInfo) => {
    token.origin.forEach((symbol: string) => {
      if (symbol.startsWith('RFQ.')) {
        rfqSymbols.add(symbol)
      }
    })
  })

  // Add these symbols to the origin arrays of all tokens
  Object.values(tokens).forEach((token: TokenInfo) => {
    const newOrigin = new Set([...token.origin])
    rfqSymbols.forEach((symbol) => newOrigin.add(symbol))
    token.origin = Array.from(newOrigin).sort()
  })
})

const outputPath = path.resolve(__dirname, '../constants/bridgeMap.ts')
prettyPrintTS(bridgeMapObject, 'BRIDGE_MAP', outputPath)

console.log('BridgeMap adjustment completed successfully!')
