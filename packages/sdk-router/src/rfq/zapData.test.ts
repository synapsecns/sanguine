import { BigNumber } from 'ethers'

import { ETH_USDC, ETH_USDT, ETH_DAI } from '../constants/testValues'
import {
  encodeZapData,
  decodeZapData,
  ZapDataV1,
  AMOUNT_NOT_PRESENT,
} from './zapData'

describe('zapData', () => {
  const zapData: ZapDataV1 = {
    target: ETH_USDC.toLowerCase(),
    payload: '0xdeadbeef',
    amountPosition: 1,
    finalToken: ETH_USDT.toLowerCase(),
    forwardTo: ETH_DAI.toLowerCase(),
    minFwdAmount: BigNumber.from('1234567890123456789012345678901234567890'),
  }

  const zapDataEmptyPayload: ZapDataV1 = {
    target: ETH_USDC.toLowerCase(),
    payload: '0x',
    amountPosition: AMOUNT_NOT_PRESENT,
    finalToken: ETH_USDT.toLowerCase(),
    forwardTo: ETH_DAI.toLowerCase(),
    minFwdAmount: BigNumber.from('1234567890'),
  }

  it('roundtrip encoding', () => {
    const encoded = encodeZapData(zapData)
    const decoded = decodeZapData(encoded)
    expect(decoded).toEqual(zapData)
  })

  it('roundtrip encoding with empty payload', () => {
    const encoded = encodeZapData(zapDataEmptyPayload)
    const decoded = decodeZapData(encoded)
    expect(decoded).toEqual(zapDataEmptyPayload)
  })
})
