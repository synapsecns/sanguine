import * as fs from 'fs'
import * as path from 'path'

// Type definitions for LayerZero metadata
interface EndpointV2 {
  address: string
}

interface Deployment {
  version: number
  eid?: string
  endpointV2?: EndpointV2
}

interface ChainMetadata {
  deployments?: Deployment[]
}

interface MetadataResponse {
  [key: string]: ChainMetadata
}

interface ChainInfo {
  eid: number
  endpointV2: string
}

interface ChainsConfig {
  [chain: string]: ChainInfo
}

// Type guard to validate metadata structure
function isMetadataResponse(data: unknown): data is MetadataResponse {
  return typeof data === 'object' && data !== null
}

async function extractChainInfo(): Promise<void> {
  const parentDir = path.join(__dirname, '../..')
  const deploymentsDir = path.join(parentDir, 'deployments')
  const outputPath = path.join(parentDir, 'configs', 'global', 'chains.json')

  // Folder name mappings
  const nameMap: Record<string, string> = {
    kaia: 'klaytn',
    cronos: 'cronosevm',
    bnb: 'bsc',
  }

  // Fetch metadata
  const response = await fetch(
    'https://metadata.layerzero-api.com/v1/metadata/deployments'
  )
  const metadata = await response.json()

  if (!isMetadataResponse(metadata)) {
    throw new Error('Invalid metadata response format')
  }

  // Process chains
  const allChains = fs
    .readdirSync(deploymentsDir)
    .filter((f) => fs.statSync(path.join(deploymentsDir, f)).isDirectory())

  const chains: ChainsConfig = {}
  const failed: string[] = []

  allChains.forEach((chain) => {
    const key = `${nameMap[chain] || chain}-mainnet`
    const v2 = metadata[key]?.deployments?.find((d) => d.version === 2)

    if (v2?.eid && v2?.endpointV2?.address) {
      chains[chain] = {
        eid: parseInt(v2.eid),
        endpointV2: v2.endpointV2.address,
      }
    } else {
      failed.push(chain)
    }
  })

  // Save result
  fs.mkdirSync(path.dirname(outputPath), { recursive: true })
  fs.writeFileSync(outputPath, JSON.stringify(chains, null, 2))

  console.log(`Saved ${Object.keys(chains).length} chains to ${outputPath}`)
  if (failed.length > 0) {
    console.log(`Failed to extract data for: ${failed.join(', ')}`)
  }
}

extractChainInfo().catch(console.error)
