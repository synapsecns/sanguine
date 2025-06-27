const fs = require('fs')
const path = require('path')

async function extractChainInfo() {
  const parentDir = path.join(__dirname, '..')
  const deploymentsDir = path.join(parentDir, 'deployments')
  const outputPath = path.join(parentDir, 'configs', 'global', 'chains.json')

  // Folder name mappings
  const nameMap = { kaia: 'klaytn', cronos: 'cronosevm', bnb: 'bsc' }

  // Fetch metadata
  const metadata = await fetch(
    'https://metadata.layerzero-api.com/v1/metadata/deployments'
  ).then((r) => r.json())

  // Process chains
  const allChains = fs
    .readdirSync(deploymentsDir)
    .filter((f) => fs.statSync(path.join(deploymentsDir, f)).isDirectory())

  const chains = {}
  const failed = []

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
