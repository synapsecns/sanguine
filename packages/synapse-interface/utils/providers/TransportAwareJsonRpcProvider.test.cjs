const assert = require('node:assert/strict')
const fs = require('node:fs')
const http = require('node:http')
const Module = require('node:module')
const path = require('node:path')
const test = require('node:test')

const ts = require('typescript')
const {
  FallbackProvider,
  StaticJsonRpcProvider,
} = require('@ethersproject/providers')

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
const callParams = { transaction, blockTag: 'latest' }
const testOptions = { timeout: 4000 }

const createFallback = (primary, secondary) => {
  return new FallbackProvider(
    [
      { provider: primary, priority: 1, stallTimeout: 250 },
      { provider: secondary, priority: 2, stallTimeout: 250 },
    ],
    1
  )
}

// Exercise ethers' real HTTP parsing and error wrapping. Every fixture binds an
// ephemeral loopback port and tears down even sockets that never send a reply.
const endpoint = async (t, handler) => {
  const requests = []
  const sockets = new Set()
  const server = http.createServer((req, res) => {
    let body = ''
    req.on('data', (chunk) => (body += chunk))
    req.on('end', () => {
      const request = JSON.parse(body)
      requests.push(request)
      if (request.method === 'eth_blockNumber') {
        res.end(
          JSON.stringify({ jsonrpc: '2.0', id: request.id, result: '0x1' })
        )
      } else {
        handler(request, res)
      }
    })
  })
  server.on('connection', (socket) => {
    sockets.add(socket)
    socket.on('close', () => sockets.delete(socket))
  })
  t.after(async () => {
    for (const socket of sockets) socket.destroy()
    await new Promise((resolve) => server.close(resolve))
  })
  await new Promise((resolve, reject) => {
    server.once('error', reject)
    server.listen(0, '127.0.0.1', resolve)
  })
  return {
    url: `http://127.0.0.1:${server.address().port}`,
    requests,
    get calls() {
      return requests.filter((request) => request.method === 'eth_call')
    },
  }
}

const reply = (status, body) => {
  return (_request, response) => {
    response.writeHead(status, { 'Content-Type': 'application/json' })
    response.end(typeof body === 'string' ? body : JSON.stringify(body))
  }
}

const rpcError = (code, message, extra = {}, envelope = true) => {
  return {
    ...(envelope ? { jsonrpc: '2.0' } : {}),
    id: 1,
    error: { code, message, ...extra },
  }
}

const healthyEndpoint = async (t) => {
  return endpoint(t, (request, response) => {
    response.writeHead(200, { 'Content-Type': 'application/json' })
    response.end(
      JSON.stringify({ jsonrpc: '2.0', id: request.id, result: '0x1234' })
    )
  })
}

const outcome = async (action) => {
  try {
    return { result: await action() }
  } catch (error) {
    return { error }
  }
}

// Request IDs and stack traces differ across calls, but the RPC error, revert
// data, HTTP response, and nested transport cause must keep ethers' semantics.
const diagnostics = (error) => {
  if (!error) return undefined
  return Object.fromEntries(
    [
      'code',
      'reason',
      'status',
      'body',
      'data',
      'transaction',
      'error',
      'serverError',
    ]
      .filter((key) => error[key] !== undefined)
      .map((key) => [
        key,
        key === 'error' || key === 'serverError'
          ? diagnostics(error[key])
          : error[key],
      ])
  )
}

const endpointFailures = [
  {
    name: 'HTTP 500 with a generic JSON error',
    handler: reply(500, { error: { code: 500 } }),
    reason: 'bad response',
    status: 500,
  },
  {
    name: 'HTTP 403 with an inactive-team JSON-RPC error',
    handler: reply(403, rpcError(-32600, 'This team has been deactivated')),
    reason: 'bad response',
    status: 403,
  },
  {
    name: 'HTTP 500 with a plain-text body',
    handler: reply(500, 'upstream unavailable'),
    reason: 'bad response',
    status: 500,
  },
  ...[
    [-32603, 'Internal error'],
    [-32601, 'Method not supported'],
    [-32002, 'Resource unavailable'],
    [-32005, 'Limit exceeded'],
  ].map(([code, message]) => ({
    name: `HTTP 200 JSON-RPC ${code} (${message})`,
    handler: reply(200, rpcError(code, message)),
    reason: 'processing response error',
    rpcCode: code,
  })),
  {
    name: 'a connection that closes without a response',
    handler: (_request, response) => response.socket.destroy(),
    reason: 'missing response',
    serverCode: 'ECONNRESET',
  },
  {
    name: 'a connection that never responds before its timeout',
    handler: () => {},
    connection: { timeout: 100 },
    reason: 'timeout',
    code: 'TIMEOUT',
  },
  {
    name: 'HTTP 429 after exhausting the retry limit',
    handler: reply(429, 'Too many requests'),
    // ethers leaves its timeout timer alive on exhausted retries. Bound that
    // timer too, without waiting through its production retry schedule.
    connection: { throttleLimit: 1, throttleSlotInterval: 1, timeout: 200 },
    reason: 'failed response',
  },
  {
    name: 'malformed JSON in an HTTP 200 response',
    handler: reply(200, '{not-json'),
    reason: 'processing response error',
    nestedReason: 'invalid JSON',
  },
]

for (const failure of endpointFailures) {
  test(
    `${failure.name} reaches the healthy fallback`,
    testOptions,
    async (t) => {
      const failed = await endpoint(t, failure.handler)
      const healthy = await healthyEndpoint(t)
      const primary = new TransportAwareJsonRpcProvider(
        { url: failed.url, ...failure.connection },
        1
      )
      const secondary = new TransportAwareJsonRpcProvider(healthy.url, 1)

      // Check the surfaced cause directly as well: stall-timeout failover alone
      // must not make a still-wrapped CALL_EXCEPTION look like a passing test.
      const direct = await outcome(() => primary.perform('call', callParams))
      assert.ok(direct.error)
      assert.equal(direct.error.code, failure.code || 'SERVER_ERROR')
      assert.equal(direct.error.reason, failure.reason)
      if (failure.status) assert.equal(direct.error.status, failure.status)
      if (failure.rpcCode)
        assert.equal(direct.error.error.code, failure.rpcCode)
      if (failure.serverCode) {
        assert.equal(direct.error.serverError.code, failure.serverCode)
      }
      if (failure.nestedReason) {
        assert.equal(direct.error.error.reason, failure.nestedReason)
      }
      assert.equal(direct.error.url, failed.url)
      assert.equal(JSON.parse(direct.error.requestBody).method, 'eth_call')

      assert.equal(
        await createFallback(primary, secondary).perform('call', callParams),
        '0x1234'
      )
      assert.equal(healthy.calls.length, 1)
      assert.equal(healthy.calls[0].method, 'eth_call')
      assert.equal(failed.calls.length, 2)
    }
  )
}

const preservedErrors = [
  {
    statuses: [200],
    name: 'an invalid request',
    body: rpcError(-32600, 'Invalid request'),
  },
  {
    name: 'a nested data-less execution revert',
    body: rpcError(-32603, 'Internal error', {
      data: { originalError: { code: 3, message: 'execution reverted' } },
    }),
  },
  {
    name: 'a code 3 revert without data',
    body: rpcError(3, 'execution reverted'),
  },
  {
    name: 'a revert without a jsonrpc marker',
    body: rpcError(3, 'execution reverted', {}, false),
  },
  {
    name: 'a generic -32000 execution revert',
    body: rpcError(-32000, 'execution reverted'),
  },
  ...[
    ['custom error', '0xdeadbeef' + '0'.repeat(64)],
    ['panic', '0x4e487b71' + '0'.repeat(62) + '11'],
    ['empty revert', '0x'],
  ].map(([name, data]) => ({
    name: `${name} data`,
    body: rpcError(3, 'execution reverted', { data }),
    result: data,
  })),
  {
    name: 'invalid params',
    body: rpcError(-32602, 'Invalid params'),
  },
  {
    statuses: [200, 400],
    name: 'an ambiguous -32000 server error',
    body: rpcError(-32000, 'Something went wrong'),
  },
  {
    statuses: [200, 400],
    name: 'an unknown JSON-RPC error code',
    body: rpcError(-32099, 'Unknown request error'),
  },
]

for (const preserved of preservedErrors) {
  for (const status of preserved.statuses || [200, 500]) {
    test(
      `HTTP ${status} ${preserved.name} preserves ethers behavior`,
      testOptions,
      async (t) => {
        const failed = await endpoint(t, reply(status, preserved.body))
        const healthy = await healthyEndpoint(t)
        const baseline = new StaticJsonRpcProvider(failed.url, 1)
        const primary = new TransportAwareJsonRpcProvider(failed.url, 1)
        const secondary = new TransportAwareJsonRpcProvider(healthy.url, 1)
        const expected = await outcome(() =>
          baseline.perform('call', callParams)
        )
        const direct = await outcome(() => primary.perform('call', callParams))
        const actual = await outcome(() =>
          createFallback(primary, secondary).perform('call', callParams)
        )

        if (preserved.result !== undefined) {
          assert.equal(expected.result, preserved.result)
          assert.deepEqual(direct, expected)
          assert.deepEqual(actual, expected)
        } else {
          assert.equal(expected.error?.code, 'CALL_EXCEPTION')
          assert.equal(actual.error?.code, 'CALL_EXCEPTION')
          assert.deepEqual(
            diagnostics(direct.error),
            diagnostics(expected.error)
          )
          assert.equal(actual.error.reason, expected.error.reason)
        }
        assert.equal(healthy.calls.length, 0)
        assert.equal(failed.calls.length, 3)
      }
    )
  }
}

test(
  'HTTP 400 invalid params remains a call exception',
  testOptions,
  async (t) => {
    const failed = await endpoint(
      t,
      reply(400, rpcError(-32602, 'Invalid params'))
    )
    const healthy = await healthyEndpoint(t)
    const primary = new TransportAwareJsonRpcProvider(failed.url, 1)
    const secondary = new TransportAwareJsonRpcProvider(healthy.url, 1)
    await assert.rejects(
      createFallback(primary, secondary).perform('call', callParams),
      { code: 'CALL_EXCEPTION' }
    )
    assert.equal(healthy.calls.length, 0)
  }
)

for (const method of ['getBalance', 'estimateGas', 'sendTransaction']) {
  test(
    `${method} retains normal ethers error behavior`,
    testOptions,
    async (t) => {
      const failed = await endpoint(t, reply(500, 'upstream unavailable'))
      const baseline = new StaticJsonRpcProvider(failed.url, 1)
      const primary = new TransportAwareJsonRpcProvider(failed.url, 1)
      const params = {
        signedTransaction: '0x1234',
        address: transaction.to,
        transaction,
        blockTag: 'latest',
      }
      const expected = await outcome(() => baseline.perform(method, params))
      const actual = await outcome(() => primary.perform(method, params))
      assert.ok(expected.error)
      assert.deepEqual(diagnostics(actual.error), diagnostics(expected.error))
    }
  )
}

test(
  'all endpoints failing retains nested HTTP and RPC diagnostics',
  testOptions,
  async (t) => {
    const firstBody = { error: 'upstream unavailable' }
    const secondBody = rpcError(-32603, 'Internal error')
    const first = await endpoint(t, reply(500, firstBody))
    const second = await endpoint(t, reply(200, secondBody))
    const provider = createFallback(
      new TransportAwareJsonRpcProvider(first.url, 1),
      new TransportAwareJsonRpcProvider(second.url, 1)
    )
    const { error } = await outcome(() => provider.perform('call', callParams))
    assert.equal(error?.code, 'SERVER_ERROR')
    assert.equal(error.reason, 'failed to meet quorum')
    assert.equal(error.results.length, 2)
    const failures = error.results.map((result) => result.error)
    const httpFailure = failures.find((failure) => failure.url === first.url)
    const rpcFailure = failures.find((failure) => failure.url === second.url)
    assert.equal(httpFailure.code, 'SERVER_ERROR')
    assert.equal(httpFailure.status, 500)
    assert.deepEqual(JSON.parse(httpFailure.body), firstBody)
    assert.equal(rpcFailure.code, 'SERVER_ERROR')
    assert.equal(rpcFailure.error.code, -32603)
    assert.equal(rpcFailure.error.message, 'Internal error')
    assert.deepEqual(JSON.parse(rpcFailure.body), secondBody)
    assert.equal(first.calls.length, 1)
    assert.equal(second.calls.length, 1)
  }
)

test(
  'a healthy primary answers without a fallback call',
  testOptions,
  async (t) => {
    const first = await healthyEndpoint(t)
    const second = await healthyEndpoint(t)
    const provider = createFallback(
      new TransportAwareJsonRpcProvider(first.url, 1),
      new TransportAwareJsonRpcProvider(second.url, 1)
    )
    assert.equal(await provider.perform('call', callParams), '0x1234')
    assert.equal(first.calls.length, 1)
    assert.equal(second.calls.length, 0)
  }
)
