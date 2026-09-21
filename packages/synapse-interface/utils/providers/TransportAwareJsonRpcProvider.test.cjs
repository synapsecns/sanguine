const assert = require('node:assert/strict')
const fs = require('node:fs')
const Module = require('node:module')
const path = require('node:path')
const test = require('node:test')
const ts = require('typescript')
const { FallbackProvider } = require('@ethersproject/providers')

const sourcePath = path.join(__dirname, 'TransportAwareJsonRpcProvider.ts')
const source = fs.readFileSync(sourcePath, 'utf8')
const compiled = ts.transpileModule(source, {
  compilerOptions: { module: ts.ModuleKind.CommonJS },
}).outputText
const compiledModule = new Module(sourcePath, module)
compiledModule.filename = sourcePath
compiledModule.paths = Module._nodeModulePaths(__dirname)
compiledModule._compile(compiled, sourcePath)
const { TransportAwareJsonRpcProvider } = compiledModule.exports

const transaction = {
  to: '0x0000000000000000000000000000000000000001',
  data: '0x1234',
}

function createFallback(primary, secondary) {
  return new FallbackProvider(
    [
      { provider: primary, priority: 1, stallTimeout: 50 },
      { provider: secondary, priority: 2, stallTimeout: 50 },
    ],
    1
  )
}

for (const status of [500, 429]) {
  test(`an HTTP ${status} eth_call uses a healthy fallback provider`, async () => {
    const primary = new TransportAwareJsonRpcProvider('http://primary.test', 1)
    const secondary = new TransportAwareJsonRpcProvider(
      'http://secondary.test',
      1
    )
    let fallbackCalls = 0
    primary.send = async () => {
      const error = new Error(`HTTP ${status}`)
      error.code = 'SERVER_ERROR'
      error.status = status
      error.body = 'Request failed'
      throw error
    }
    secondary.send = async (method) => {
      if (method === 'eth_call') fallbackCalls += 1
      return '0x1234'
    }

    const provider = createFallback(primary, secondary)
    assert.equal(
      await provider.perform('call', { transaction, blockTag: 'latest' }),
      '0x1234'
    )
    assert.equal(fallbackCalls, 1)
  })
}

for (const status of [undefined, 500]) {
  test(`a JSON-RPC revert remains a call exception (HTTP ${status ??
    200})`, async () => {
    const primary = new TransportAwareJsonRpcProvider('http://primary.test', 1)
    const secondary = new TransportAwareJsonRpcProvider(
      'http://secondary.test',
      1
    )
    let fallbackCalls = 0
    primary.send = async () => {
      const error = new Error('execution reverted')
      error.code = 'SERVER_ERROR'
      error.status = status
      error.body = JSON.stringify({
        jsonrpc: '2.0',
        error: { code: 3, message: 'execution reverted' },
        id: 1,
      })
      throw error
    }
    secondary.send = async (method) => {
      if (method === 'eth_call') fallbackCalls += 1
      return '0x1234'
    }

    const provider = createFallback(primary, secondary)
    await assert.rejects(
      provider.perform('call', { transaction, blockTag: 'latest' }),
      { code: 'CALL_EXCEPTION' }
    )
    assert.equal(fallbackCalls, 0)
  })
}
