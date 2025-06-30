import * as fs from 'fs'
import * as path from 'path'
import { getAddress, type Address } from 'viem'

// Type definitions for DVN metadata
interface DVNInfo {
  version: number
  canonicalName: string
  id: string
  deprecated?: boolean
}

interface ChainMetadata {
  created: string
  updated: string
  tableName: string
  environment: string
  chainName: string
  dvns?: {
    [dvnAddress: string]: DVNInfo
  }
}

interface DVNMetadataResponse {
  [chainName: string]: ChainMetadata
}

interface DVNConfig {
  [chain: string]: {
    [dvnName: string]: Address
  }
}

// Type guard to validate DVN metadata structure
function isDVNMetadataResponse(data: unknown): data is DVNMetadataResponse {
  return typeof data === 'object' && data !== null
}

async function extractDVNs(): Promise<void> {
  const parentDir = path.join(__dirname, '../..')
  const deploymentsDir = path.join(parentDir, 'deployments')
  const outputPath = path.join(parentDir, 'configs', 'global', 'dvns.json')

  // Folder name mappings for LayerZero (same as in extractChainInfo.ts)
  const lzNameMap: Record<string, string> = {
    kaia: 'klaytn',
    cronos: 'cronosevm',
    bnb: 'bsc',
  }

  // Fetch DVN metadata
  console.log('Fetching DVN metadata from LayerZero...')
  const response = await fetch(
    'https://metadata.layerzero-api.com/v1/metadata/dvns'
  )
  const metadata = await response.json()

  if (!isDVNMetadataResponse(metadata)) {
    throw new Error('Invalid DVN metadata response format')
  }

  // Get all chains from deployments directory
  const allChains = fs
    .readdirSync(deploymentsDir)
    .filter((f) => fs.statSync(path.join(deploymentsDir, f)).isDirectory())

  console.log(`Found ${allChains.length} chains in deployments directory`)

  // Create mapping of LayerZero chain names to our chain names
  const lzToOurChainMap: Record<string, string> = {}
  for (const chain of allChains) {
    const lzChainName = lzNameMap[chain] || chain
    lzToOurChainMap[lzChainName] = chain
  }

  // Process DVN data - collect DVN coverage across chains
  const dvnCoverage: Map<string, Map<string, string>> = new Map()

  for (const [lzChainName, chainData] of Object.entries(metadata)) {
    const ourChainName = lzToOurChainMap[lzChainName]

    if (!ourChainName || !chainData.dvns) continue

    for (const [dvnAddress, dvnInfo] of Object.entries(chainData.dvns)) {
      // Skip deprecated DVNs
      if (dvnInfo.deprecated) continue

      // Use canonicalName as the DVN identifier
      // Imagine putting spaces in JSON keys
      const dvnName = dvnInfo.canonicalName.replace(/ /g, '_')

      if (!dvnCoverage.has(dvnName)) {
        dvnCoverage.set(dvnName, new Map())
      }

      dvnCoverage.get(dvnName)!.set(ourChainName, dvnAddress)
    }
  }

  // Filter DVNs that are present on all required chains and build output
  const dvnsByChain: DVNConfig = {}
  const universalDVNs: string[] = []

  // Initialize chain objects
  for (const chain of allChains) {
    dvnsByChain[chain] = {}
  }

  // Add DVNs that exist on all chains
  for (const [dvnName, chainMap] of dvnCoverage) {
    if (allChains.every((chain) => chainMap.has(chain))) {
      universalDVNs.push(dvnName)

      for (const chain of allChains) {
        dvnsByChain[chain][dvnName] = getAddress(chainMap.get(chain)!)
      }
    }
  }

  console.log(`Found ${universalDVNs.length} DVNs with full chain coverage`)

  // Save result
  fs.mkdirSync(path.dirname(outputPath), { recursive: true })
  fs.writeFileSync(outputPath, JSON.stringify(dvnsByChain, null, 2))

  console.log(
    `Saved DVN data for ${
      Object.keys(dvnsByChain).length
    } chains to ${outputPath}`
  )

  // Log summary
  console.log(`\nUniversal DVNs included: ${universalDVNs.sort().join(', ')}`)
}

extractDVNs().catch(console.error)
