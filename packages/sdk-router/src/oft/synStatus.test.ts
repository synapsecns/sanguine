import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
import { utils } from 'ethers'

import {
  HYPERCORE_CHAIN_ID,
  SYN_ADDRESS_MAP,
  SYN_COMPOSER_ADDRESS,
  SupportedChainId,
} from '../constants'
import { getSynBridgeDeliveryChainId, getSynBridgeStatus } from './synStatus'

const sourceTx = '0x' + '11'.repeat(32)
const composeTx = '0x' + '22'.repeat(32)
const recipient = '0x1234567890123456789012345678901234567890'
const coreWriter = '0x3333333333333333333333333333333333333333'
const coreWriterInterface = new Interface([
  'event RawAction(address indexed user, bytes data)',
])
const tokenInterface = new Interface([
  'event Transfer(address indexed from, address indexed to, uint256 value)',
])

const fallbackLog = (
  overrides: {
    token?: string
    from?: string
    to?: string
    amount?: string
  } = {}
) => ({
  address: overrides.token ?? SYN_ADDRESS_MAP[SupportedChainId.HYPEREVM],
  ...tokenInterface.encodeEventLog(tokenInterface.getEvent('Transfer'), [
    overrides.from ?? SYN_COMPOSER_ADDRESS,
    overrides.to ?? recipient,
    overrides.amount ?? utils.parseEther('1').toString(),
  ]),
})

const payload = utils.solidityPack(
  ['bytes32', 'uint64', 'bytes32', 'bytes'],
  [
    utils.hexZeroPad(SYN_COMPOSER_ADDRESS, 32),
    1000000,
    utils.hexZeroPad(recipient, 32),
    utils.defaultAbiCoder.encode(['uint256', 'address'], [0, recipient]),
  ]
)

const coreMessage = (composeStatus = 'SUCCEEDED') => ({
  pathway: { dstEid: 30367 },
  status: { name: 'DELIVERED' },
  source: { tx: { payload } },
  destination: {
    status: 'SUCCEEDED',
    lzCompose: { status: composeStatus, txs: [{ txHash: composeTx }] },
  },
})

const rawActionLog = (amount = 100000000) => {
  const action = utils.hexConcat([
    '0x01000006',
    utils.defaultAbiCoder.encode(
      ['address', 'uint64', 'uint64'],
      [recipient, 873, amount]
    ),
  ])
  return {
    address: coreWriter,
    ...coreWriterInterface.encodeEventLog(
      coreWriterInterface.getEvent('RawAction'),
      [SYN_COMPOSER_ADDRESS, action]
    ),
  }
}

describe('SYN bridge status', () => {
  let provider: Provider

  beforeEach(() => {
    global.fetch = jest.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ data: [coreMessage()] }),
    }) as any
    provider = {
      getTransactionReceipt: jest.fn().mockResolvedValue({
        status: 1,
        logs: [rawActionLog()],
      }),
    } as unknown as Provider
  })

  afterEach(() => jest.restoreAllMocks())

  it('waits for a finalized ordinary OFT receive', async () => {
    expect(
      await getSynBridgeDeliveryChainId(SupportedChainId.HYPEREVM, sourceTx)
    ).toBe(SupportedChainId.HYPEREVM)
    expect(await getSynBridgeStatus(SupportedChainId.HYPEREVM, sourceTx)).toBe(
      true
    )
    ;(global.fetch as jest.Mock).mockResolvedValueOnce({
      ok: true,
      json: async () => ({
        data: [{ ...coreMessage(), status: { name: 'CONFIRMING' } }],
      }),
    })
    expect(await getSynBridgeStatus(SupportedChainId.HYPEREVM, sourceTx)).toBe(
      false
    )
  })

  it('requires a finalized compose for the HyperCore route', async () => {
    ;(global.fetch as jest.Mock).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ data: [coreMessage('WAITING')] }),
    })
    expect(
      await getSynBridgeDeliveryChainId(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBeUndefined()
    expect(provider.getTransactionReceipt).not.toHaveBeenCalled()
  })

  it('requires a HyperEVM provider to inspect the compose receipt', async () => {
    expect(await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx)).toBe(false)
  })

  it('recognizes the matching CoreWriter SpotSend submission', async () => {
    expect(
      await getSynBridgeDeliveryChainId(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBe(HYPERCORE_CHAIN_ID)
    expect(global.fetch).toHaveBeenCalledTimes(1)
    expect(provider.getTransactionReceipt).toHaveBeenCalledTimes(1)
    expect(
      await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBe(true)
    expect(provider.getTransactionReceipt).toHaveBeenCalledWith(composeTx)
  })

  it('does not complete a compose with no verified settlement', async () => {
    ;(provider.getTransactionReceipt as jest.Mock).mockResolvedValueOnce({
      status: 1,
      logs: [],
    })
    expect(
      await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBe(false)
  })

  it('completes the exact automatic SYN fallback on HyperEVM', async () => {
    ;(provider.getTransactionReceipt as jest.Mock).mockResolvedValue({
      status: 1,
      logs: [fallbackLog()],
    })
    expect(
      await getSynBridgeDeliveryChainId(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBe(SupportedChainId.HYPEREVM)
    expect(global.fetch).toHaveBeenCalledTimes(1)
    expect(provider.getTransactionReceipt).toHaveBeenCalledTimes(1)
    expect(
      await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBe(true)
  })

  it.each([
    { token: recipient },
    { from: recipient },
    { to: SYN_COMPOSER_ADDRESS },
    { amount: '999999999999999999' },
  ])('rejects an unrelated or partial fallback %j', async (overrides) => {
    ;(provider.getTransactionReceipt as jest.Mock).mockResolvedValueOnce({
      status: 1,
      logs: [fallbackLog(overrides)],
    })
    expect(
      await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBe(false)
  })

  it('does not complete a failed compose receipt even with a matching transfer', async () => {
    ;(provider.getTransactionReceipt as jest.Mock).mockResolvedValueOnce({
      status: 0,
      logs: [fallbackLog()],
    })
    expect(
      await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBe(false)
  })

  it.each([
    utils.hexDataSlice(payload, 0, 135),
    utils.hexConcat([
      utils.hexZeroPad(recipient, 32),
      utils.hexDataSlice(payload, 32),
    ]),
    utils.hexConcat([
      utils.hexDataSlice(payload, 0, 32),
      '0x0000000000000000',
      utils.hexDataSlice(payload, 40),
    ]),
    utils.hexConcat([
      utils.hexDataSlice(payload, 0, 72),
      utils.defaultAbiCoder.encode(['uint256', 'address'], [1, recipient]),
    ]),
    utils.hexConcat([
      utils.hexDataSlice(payload, 0, 104),
      '0x01',
      utils.hexDataSlice(payload, 105),
    ]),
  ])(
    'rejects an invalid source payload before accepting a fallback',
    async (invalidPayload) => {
      ;(global.fetch as jest.Mock).mockResolvedValueOnce({
        ok: true,
        json: async () => ({
          data: [
            { ...coreMessage(), source: { tx: { payload: invalidPayload } } },
          ],
        }),
      })
      ;(provider.getTransactionReceipt as jest.Mock).mockResolvedValueOnce({
        status: 1,
        logs: [fallbackLog()],
      })
      expect(
        await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx, provider)
      ).toBe(false)
      expect(provider.getTransactionReceipt).not.toHaveBeenCalled()
    }
  )

  it('rejects an unrelated CoreWriter action in the same receipt', async () => {
    ;(provider.getTransactionReceipt as jest.Mock).mockResolvedValueOnce({
      status: 1,
      logs: [rawActionLog(999)],
    })
    expect(
      await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBe(false)
  })
})
