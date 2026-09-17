import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
import { utils } from 'ethers'

import {
  HYPERCORE_CHAIN_ID,
  SYN_COMPOSER_ADDRESS,
  SupportedChainId,
} from '../constants'
import { getSynBridgeStatus } from './synStatus'

const sourceTx = '0x' + '11'.repeat(32)
const composeTx = '0x' + '22'.repeat(32)
const recipient = '0x1234567890123456789012345678901234567890'
const coreWriter = '0x3333333333333333333333333333333333333333'
const coreWriterInterface = new Interface([
  'event RawAction(address indexed user, bytes data)',
])

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
      await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBe(false)
    expect(provider.getTransactionReceipt).not.toHaveBeenCalled()
  })

  it('requires a HyperEVM provider to inspect the compose receipt', async () => {
    expect(await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx)).toBe(false)
  })

  it('recognizes the matching CoreWriter SpotSend submission', async () => {
    expect(
      await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBe(true)
    expect(provider.getTransactionReceipt).toHaveBeenCalledWith(composeTx)
  })

  it('does not mistake a successful compose with a HyperEVM refund for Core delivery', async () => {
    ;(provider.getTransactionReceipt as jest.Mock).mockResolvedValueOnce({
      status: 1,
      logs: [],
    })
    expect(
      await getSynBridgeStatus(HYPERCORE_CHAIN_ID, sourceTx, provider)
    ).toBe(false)
  })

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
