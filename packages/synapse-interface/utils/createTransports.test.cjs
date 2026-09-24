const assert = require('node:assert/strict')
const fs = require('node:fs')
const path = require('node:path')
const test = require('node:test')
const vm = require('node:vm')
const ts = require('typescript')
const { extractRpcUrls } = require('@wagmi/core')
const { arbitrum, mainnet } = require('@wagmi/core/chains')
const { HttpConnection } = require('@walletconnect/jsonrpc-http-connection')

const chains = [arbitrum, mainnet]
const chainConfig = Object.fromEntries(
  chains.map((chain) => [
    chain.id,
    {
      rpcUrls: {
        primary: chain.rpcUrls.default.http[0],
        fallback: `https://fallback.example/${chain.id}`,
      },
    },
  ])
)
const source = fs.readFileSync(
  path.join(__dirname, 'createTransports.ts'),
  'utf8'
)
const compiled = ts.transpileModule(source, {
  compilerOptions: { module: ts.ModuleKind.CommonJS },
}).outputText

function loadTransports(origin) {
  const context = {
    exports: {},
    URL,
    require: (id) =>
      id === '@/constants/chains'
        ? { CHAINS_BY_ID: chainConfig }
        : require(id),
  }
  if (origin) context.window = { location: { origin } }
  vm.runInNewContext(compiled, context)
  return context.exports.createTransports(chains)
}

for (const origin of [
  'https://bridge.synapseprotocol.com',
  'https://preview.example',
  'http://localhost:3000',
]) {
  test(`WalletConnect accepts extracted RPC URLs on ${origin}`, () => {
    const transports = loadTransports(origin)
    for (const chain of chains) {
      // Exercise the same extraction and URL validation as the connector.
      const urls = extractRpcUrls({ chain, transports })
      assert.deepEqual(Array.from(urls), [
        `${origin}/api/rpc/${chain.id}`,
        chainConfig[chain.id].rpcUrls.primary,
        chainConfig[chain.id].rpcUrls.fallback,
      ])
      assert.doesNotThrow(() => new HttpConnection(urls[0]))
    }
  })
}

test('transport configuration remains safe to load during SSR', () => {
  const transports = loadTransports()
  const urls = extractRpcUrls({ chain: arbitrum, transports })
  assert.equal(urls[0], '/api/rpc/42161')
})
