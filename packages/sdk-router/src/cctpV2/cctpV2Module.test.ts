import { hexDataSlice } from '@ethersproject/bytes'
import { AddressZero } from '@ethersproject/constants'
import { BigNumber, utils } from 'ethers'

import {
  CCTP_V2_DOMAIN_MAP,
  CCTP_V2_FORWARD_HOOK_DATA,
  SupportedChainId,
} from '../constants'
import { getMessages } from './api'
import { CctpV2BurnParams, CCTPv2Module } from './cctpV2Module'

jest.mock('./api', () => ({
  getMessages: jest.fn(),
}))

const mockGetMessages = getMessages as jest.MockedFunction<typeof getMessages>

const MODULE_ADDRESS = '0x1111111111111111111111111111111111111111'
const MINT_RECIPIENT = '0x2222222222222222222222222222222222222222'
const BURN_TOKEN = '0x3333333333333333333333333333333333333333'
const ORIGIN_CHAIN_ID = SupportedChainId.ETH
const ORIGIN_DOMAIN = CCTP_V2_DOMAIN_MAP[ORIGIN_CHAIN_ID]
const ORIGIN_TX_HASH =
  '0xcd593dc11f7607e2e48c1cc70236c0a993cf54b37ad398d14e485087b4508d34'
const CCTP_V2_SYNAPSE_TX_ID = `${ORIGIN_TX_HASH}:${ORIGIN_CHAIN_ID}`

describe('CCTPv2Module', () => {
  let module: CCTPv2Module

  const burnParams: CctpV2BurnParams = {
    amount: 1_000_000,
    destinationDomain: CCTP_V2_DOMAIN_MAP[SupportedChainId.ARBITRUM],
    mintRecipient: MINT_RECIPIENT,
    burnToken: BURN_TOKEN,
    maxFee: 25_000,
    minFinalityThreshold: 2_000,
  }

  beforeEach(() => {
    module = new CCTPv2Module(SupportedChainId.ETH, MODULE_ADDRESS)
    mockGetMessages.mockReset()
  })

  it('formats synapseTxId as txHash:originChainId', async () => {
    await expect(module.getSynapseTxId(ORIGIN_TX_HASH)).resolves.toBe(
      CCTP_V2_SYNAPSE_TX_ID
    )
  })

  it('encodes depositForBurnWithHook calldata', () => {
    const tx = module.populateDepositForBurnWithHook(burnParams)
    const expectedData =
      CCTPv2Module.tokenMessengerV2Interface.encodeFunctionData(
        'depositForBurnWithHook',
        [
          burnParams.amount,
          burnParams.destinationDomain,
          utils.hexZeroPad(burnParams.mintRecipient, 32),
          burnParams.burnToken,
          utils.hexZeroPad(AddressZero, 32),
          burnParams.maxFee,
          burnParams.minFinalityThreshold,
          CCTP_V2_FORWARD_HOOK_DATA,
        ]
      )

    expect(tx).toEqual({
      to: MODULE_ADDRESS,
      data: expectedData,
    })
  })

  it('uses bytes hookData in TokenMessengerV2 ABI', () => {
    const fn = CCTPv2Module.tokenMessengerV2Interface.getFunction(
      'depositForBurnWithHook'
    )
    expect(fn.inputs[7].type).toBe('bytes')
  })

  it('detects amount position within encoded calldata', () => {
    const burnParamsWithoutAmount = {
      destinationDomain: burnParams.destinationDomain,
      mintRecipient: burnParams.mintRecipient,
      burnToken: burnParams.burnToken,
      maxFee: burnParams.maxFee,
      minFinalityThreshold: burnParams.minFinalityThreshold,
    }

    const amountPosition = module.getAmountPosition(burnParamsWithoutAmount)
    const tx = module.populateDepositForBurnWithHook({
      ...burnParamsWithoutAmount,
      amount: 42,
    })
    const encodedAmount = utils.hexZeroPad(BigNumber.from(42).toHexString(), 32)

    expect(amountPosition).toBe(4)
    expect(hexDataSlice(tx.data!, amountPosition, amountPosition + 32)).toBe(
      encodedAmount
    )
    expect(module.getAmountPosition(burnParamsWithoutAmount)).toBe(
      amountPosition
    )
  })

  it('uses the forwarding hook constant in calldata', () => {
    const tx = module.populateDepositForBurnWithHook(burnParams)
    const decoded = CCTPv2Module.tokenMessengerV2Interface.decodeFunctionData(
      'depositForBurnWithHook',
      tx.data!
    )

    expect(decoded.hookData).toBe(CCTP_V2_FORWARD_HOOK_DATA)
  })

  it.each<[string, { status: string; forwardState: string }, boolean]>([
    [
      'terminal success + forwarding success',
      { status: 'complete', forwardState: 'succeeded' },
      true,
    ],
    [
      'terminal success + forwarding confirmed',
      { status: 'complete', forwardState: 'confirmed' },
      true,
    ],
    [
      'pending message status',
      { status: 'pending', forwardState: 'succeeded' },
      false,
    ],
    [
      'pending forwarding status',
      { status: 'complete', forwardState: 'pending' },
      false,
    ],
    [
      'error message status',
      { status: 'error', forwardState: 'succeeded' },
      false,
    ],
  ])(
    'returns expected bridge status for %s',
    async (_caseName, message, expectedStatus) => {
      mockGetMessages.mockResolvedValueOnce([message])

      await expect(
        module.getBridgeTxStatus(CCTP_V2_SYNAPSE_TX_ID)
      ).resolves.toBe(expectedStatus)
      expect(mockGetMessages).toHaveBeenCalledWith(
        ORIGIN_DOMAIN,
        ORIGIN_TX_HASH
      )
    }
  )

  it('returns false for malformed CCTPv2 synapseTxId', async () => {
    await expect(module.getBridgeTxStatus(ORIGIN_TX_HASH)).resolves.toBe(false)
    expect(mockGetMessages).not.toHaveBeenCalled()
  })

  it('returns false for malformed txHash segment', async () => {
    await expect(
      module.getBridgeTxStatus(`0xabc:${ORIGIN_CHAIN_ID}`)
    ).resolves.toBe(false)
    expect(mockGetMessages).not.toHaveBeenCalled()
  })

  it('returns false for unsupported origin chain id segment', async () => {
    await expect(
      module.getBridgeTxStatus(`${ORIGIN_TX_HASH}:999999`)
    ).resolves.toBe(false)
    expect(mockGetMessages).not.toHaveBeenCalled()
  })

  it('returns false when status API errors', async () => {
    mockGetMessages.mockRejectedValueOnce(new Error('Circle API unavailable'))

    await expect(module.getBridgeTxStatus(CCTP_V2_SYNAPSE_TX_ID)).resolves.toBe(
      false
    )
  })
})
