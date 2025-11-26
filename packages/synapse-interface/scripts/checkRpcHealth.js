/**
 * RPC Health Check Script
 * Dynamically parses chain config files to extract RPC URLs
 * Tests all URLs in parallel with configurable concurrency
 *
 * Usage:
 *   node scripts/checkRpcHealth.js                         # default: synapse-interface
 *   node scripts/checkRpcHealth.js --package rest-api
 *   node scripts/checkRpcHealth.js --package widget
 *   node scripts/checkRpcHealth.js --package synapse-constants
 *   node scripts/checkRpcHealth.js --package all           # check all packages
 */

const fs = require('fs')
const path = require('path')

const MAX_RETRIES = 3
const TIMEOUT_MS = 10000
const CONCURRENCY = 10

// Chains where chain ID mismatch is expected (deliberate choice)
const IGNORE_CHAIN_ID_MISMATCH = [998] // Hyperliquid

const PACKAGE_PATHS = {
  'synapse-interface': '../constants/chains/master.tsx',
  'synapse-constants': '../../synapse-constants/src/constants/chains/master.ts',
  'rest-api': '../../rest-api/src/constants/chains.ts',
  widget: '../../widget/src/constants/chains.ts',
}

/**
 * Parse CLI arguments
 */
function parseArgs() {
  const args = process.argv.slice(2)
  let packageName = 'synapse-interface'

  for (let i = 0; i < args.length; i++) {
    if (args[i] === '--package' && args[i + 1]) {
      packageName = args[i + 1]
      i++
    }
  }

  const validOptions = [...Object.keys(PACKAGE_PATHS), 'all']
  if (!validOptions.includes(packageName)) {
    console.error(
      `Unknown package: ${packageName}. Valid options: ${validOptions.join(
        ', '
      )}`
    )
    process.exit(1)
  }

  return { packageName }
}

/**
 * Parse chain config file to extract chain names and RPC URLs
 * Skips dynamic URLs like getOmniRpcUrl()
 */
function parseChainRpcUrls(filePath) {
  const content = fs.readFileSync(filePath, 'utf-8')
  const chains = []

  // Match each chain export block
  const chainBlockRegex = /export const (\w+): Chain = \{([\s\S]*?)\n\}/g
  let match

  while ((match = chainBlockRegex.exec(content)) !== null) {
    const chainName = match[1]
    const blockContent = match[2]

    // Extract name field
    const nameMatch = blockContent.match(/name:\s*['"]([^'"]+)['"]/)
    const displayName = nameMatch ? nameMatch[1] : chainName

    // Extract rpcUrls block
    const rpcUrlsMatch = blockContent.match(/rpcUrls:\s*\{([\s\S]*?)\}/)
    if (!rpcUrlsMatch) continue

    const rpcBlock = rpcUrlsMatch[1]

    // Extract primary URL (only hardcoded strings, skip function calls like getOmniRpcUrl)
    const primaryMatch =
      rpcBlock.match(/primary:\s*['"]([^'"]+)['"]/) ||
      rpcBlock.match(/primary:\s*\n\s*['"]([^'"]+)['"]/)

    // Extract fallback URL (only hardcoded strings)
    const fallbackMatch =
      rpcBlock.match(/fallback:\s*['"]([^'"]+)['"]/) ||
      rpcBlock.match(/fallback:\s*\n\s*['"]([^'"]+)['"]/)

    // Extract networkUrl (only hardcoded strings, outside rpcUrls block)
    const networkUrlMatch = blockContent.match(/networkUrl:\s*['"]([^'"]+)['"]/)

    // Extract chain ID
    const idMatch = blockContent.match(/id:\s*(\d+)/)
    const chainId = idMatch ? parseInt(idMatch[1], 10) : null

    if (primaryMatch || fallbackMatch || networkUrlMatch) {
      chains.push({
        chain: chainName,
        name: displayName,
        id: chainId,
        primary: primaryMatch ? primaryMatch[1] : null,
        fallback: fallbackMatch ? fallbackMatch[1] : null,
        networkUrl: networkUrlMatch ? networkUrlMatch[1] : null,
      })
    }
  }

  return chains
}

/**
 * Run async tasks with limited concurrency
 */
async function runWithConcurrency(tasks, concurrency) {
  const results = []
  const executing = new Set()

  for (const task of tasks) {
    const promise = task().then((result) => {
      executing.delete(promise)
      return result
    })
    executing.add(promise)
    results.push(promise)

    if (executing.size >= concurrency) {
      await Promise.race(executing)
    }
  }

  return Promise.all(results)
}

/**
 * Test a single RPC endpoint with retries
 * Validates that the returned chain ID matches the expected value
 */
async function testRpc(url, expectedChainId, retries = MAX_RETRIES) {
  for (let attempt = 1; attempt <= retries; attempt++) {
    try {
      const controller = new AbortController()
      const timeoutId = setTimeout(() => controller.abort(), TIMEOUT_MS)

      const response = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          jsonrpc: '2.0',
          method: 'eth_chainId',
          params: [],
          id: 1,
        }),
        signal: controller.signal,
      })

      clearTimeout(timeoutId)

      const data = await response.json()
      if (data.result) {
        const returnedChainId = parseInt(data.result, 16)
        if (expectedChainId && returnedChainId !== expectedChainId) {
          // Check if this chain should ignore chain ID mismatch
          if (IGNORE_CHAIN_ID_MISMATCH.includes(expectedChainId)) {
            return {
              success: true,
              warning: `Chain ID mismatch (expected ${expectedChainId}, got ${returnedChainId})`,
              chainId: returnedChainId,
            }
          }
          return {
            success: false,
            error: `Chain ID mismatch: expected ${expectedChainId}, got ${returnedChainId}`,
          }
        }
        return { success: true, chainId: returnedChainId }
      }
      if (data.error) {
        throw new Error(data.error.message || 'RPC error')
      }
      throw new Error('No result in response')
    } catch (error) {
      if (attempt < retries) {
        await new Promise((r) => setTimeout(r, 1000))
      } else {
        return { success: false, error: error.message }
      }
    }
  }
  return { success: false, error: 'Max retries exceeded' }
}

/**
 * Check a single package and return results
 */
async function checkPackage(packageName) {
  const configPath = path.join(__dirname, PACKAGE_PATHS[packageName])

  // Check if file exists
  if (!fs.existsSync(configPath)) {
    console.error(`Config file not found: ${configPath}`)
    return { failed: [], duplicates: [], total: 0, success: 0 }
  }

  const chains = parseChainRpcUrls(configPath)

  // Build list of all URLs to test
  const urlsToTest = []
  for (const chain of chains) {
    if (chain.primary) {
      urlsToTest.push({
        chain: chain.chain,
        name: chain.name,
        chainId: chain.id,
        type: 'primary',
        url: chain.primary,
      })
    }
    if (chain.fallback) {
      urlsToTest.push({
        chain: chain.chain,
        name: chain.name,
        chainId: chain.id,
        type: 'fallback',
        url: chain.fallback,
      })
    }
    if (chain.networkUrl) {
      urlsToTest.push({
        chain: chain.chain,
        name: chain.name,
        chainId: chain.id,
        type: 'networkUrl',
        url: chain.networkUrl,
      })
    }
  }

  console.log(`Package: ${packageName}`)
  console.log(`Parsed ${chains.length} chains`)
  console.log(
    `Testing ${urlsToTest.length} URLs with ${CONCURRENCY} concurrent workers`
  )
  console.log('')

  const failed = []
  const warnings = []

  // Create tasks for parallel execution
  const tasks = urlsToTest.map((item) => async () => {
    const result = await testRpc(item.url, item.chainId)
    let status
    if (result.success && result.warning) {
      status = '\x1b[33m⚠\x1b[0m' // Yellow warning
    } else if (result.success) {
      status = '\x1b[32m✓\x1b[0m' // Green success
    } else {
      status = '\x1b[31m✗\x1b[0m' // Red failure
    }
    console.log(`${status} ${item.chain} (${item.type}): ${item.url}`)

    if (!result.success) {
      failed.push({ ...item, error: result.error, package: packageName })
    } else if (result.warning) {
      warnings.push({ ...item, warning: result.warning, package: packageName })
    }
    return { ...item, ...result }
  })

  // Run all tests in parallel with concurrency limit
  await runWithConcurrency(tasks, CONCURRENCY)

  // Check for duplicate primary/fallback URLs
  const duplicates = chains
    .filter(
      (chain) =>
        chain.primary && chain.fallback && chain.primary === chain.fallback
    )
    .map((d) => ({ ...d, package: packageName }))

  return {
    failed,
    warnings,
    duplicates,
    total: urlsToTest.length,
    success: urlsToTest.length - failed.length,
  }
}

async function main() {
  const { packageName } = parseArgs()

  console.log('RPC Health Check')
  console.log('================')
  console.log(`Max retries: ${MAX_RETRIES}, Timeout: ${TIMEOUT_MS}ms`)
  console.log('')

  const packages =
    packageName === 'all' ? Object.keys(PACKAGE_PATHS) : [packageName]

  let totalUrls = 0
  let totalSuccess = 0
  const allFailed = []
  const allWarnings = []
  const allDuplicates = []

  for (const pkg of packages) {
    if (packages.length > 1) {
      console.log('----------------')
    }
    const result = await checkPackage(pkg)
    totalUrls += result.total
    totalSuccess += result.success
    allFailed.push(...result.failed)
    allWarnings.push(...result.warnings)
    allDuplicates.push(...result.duplicates)
    console.log('')
  }

  // Print summary
  console.log('================')
  console.log('Summary')
  console.log('================')

  console.log(`Total: ${totalUrls} URLs`)
  console.log(`\x1b[32mWorking: ${totalSuccess}\x1b[0m`)
  console.log(`\x1b[33mWarnings: ${allWarnings.length}\x1b[0m`)
  console.log(`\x1b[31mFailed: ${allFailed.length}\x1b[0m`)

  if (allWarnings.length > 0) {
    console.log('')
    console.log('\x1b[33mChain ID warnings (expected):\x1b[0m')
    for (const w of allWarnings) {
      const pkgPrefix = packages.length > 1 ? `[${w.package}] ` : ''
      console.log(`  - ${pkgPrefix}${w.chain} (${w.type}): ${w.warning}`)
    }
  }

  if (allFailed.length > 0) {
    console.log('')
    console.log('Failed URLs:')
    for (const f of allFailed) {
      const pkgPrefix = packages.length > 1 ? `[${f.package}] ` : ''
      console.log(`  - ${pkgPrefix}${f.chain} (${f.type}): ${f.url}`)
      console.log(`    Error: ${f.error}`)
    }
  }

  if (allDuplicates.length > 0) {
    console.log('')
    console.log(
      `\x1b[33mDuplicate primary/fallback URLs (${allDuplicates.length}):\x1b[0m`
    )
    for (const d of allDuplicates) {
      const pkgPrefix = packages.length > 1 ? `[${d.package}] ` : ''
      console.log(`  - ${pkgPrefix}${d.chain} (${d.name}): ${d.primary}`)
    }
  }
}

main().catch(console.error)
