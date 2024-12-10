import { BigNumber } from 'ethers'

import {
  BridgeParamsV2,
  encodeSavedBridgeParams,
  decodeSavedBridgeParams,
} from './paramsV2'
import { ZapDataV1, applyDefaultValues } from './zapData'
import {
  ETH_USDC,
  ETH_USDT,
  ETH_DAI,
  ETH_SYN,
  ETH_NUSD,
} from '../constants/testValues'

describe('paramsV2', () => {
  const ether = BigNumber.from(10).pow(18)
  const sender = ETH_USDC

  const paramsV2: BridgeParamsV2 = {
    quoteRelayer: ETH_USDT,
    quoteExclusivitySeconds: BigNumber.from(30),
    quoteId: '0xdead',
    zapNative: ether,
    zapData: '0xbeef',
  }

  const paramsV2QuoteIdEmpty: BridgeParamsV2 = {
    quoteRelayer: ETH_USDT,
    quoteExclusivitySeconds: BigNumber.from(30),
    quoteId: '0x',
    zapNative: ether,
    zapData: '0xdeadbeef',
  }

  const paramsV2ZapDataEmpty: BridgeParamsV2 = {
    quoteRelayer: ETH_USDT,
    quoteExclusivitySeconds: BigNumber.from(30),
    quoteId: '0xdeadbeef',
    zapNative: ether,
    zapData: '0x',
  }

  const paramsV2AllBytesEmpty: BridgeParamsV2 = {
    quoteRelayer: ETH_USDT,
    quoteExclusivitySeconds: BigNumber.from(30),
    quoteId: '0x',
    zapNative: ether,
    zapData: '0x',
  }

  const paramsV2NegativeExclusivity: BridgeParamsV2 = {
    quoteRelayer: ETH_USDT,
    quoteExclusivitySeconds: BigNumber.from(-30),
    quoteId: '0xdeadbeef',
    zapNative: ether,
    zapData: '0x',
  }

  const zapData: ZapDataV1 = {
    target: ETH_DAI,
    payload: '0xb00b1e55',
    amountPosition: BigNumber.from(420),
    finalToken: ETH_SYN,
    forwardTo: ETH_NUSD,
  }

  const zapDataDefault = applyDefaultValues({})

  it('roundtrip encoding', () => {
    const encoded = encodeSavedBridgeParams(sender, paramsV2, zapData)
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.sender).toEqual(sender)
    expect(decoded.paramsV2).toEqual(paramsV2)
    expect(decoded.zapData).toEqual(zapData)
  })

  it('roundtrip encoding with empty quoteId', () => {
    const encoded = encodeSavedBridgeParams(
      sender,
      paramsV2QuoteIdEmpty,
      zapData
    )
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.sender).toEqual(sender)
    expect(decoded.paramsV2).toEqual(paramsV2QuoteIdEmpty)
    expect(decoded.zapData).toEqual(zapData)
  })

  it('roundtrip encoding with empty zapData', () => {
    const encoded = encodeSavedBridgeParams(
      sender,
      paramsV2ZapDataEmpty,
      zapDataDefault
    )
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.sender).toEqual(sender)
    expect(decoded.paramsV2).toEqual(paramsV2ZapDataEmpty)
    expect(decoded.zapData).toEqual(zapDataDefault)
  })

  it('roundtrip encoding with empty quoteId and zapData', () => {
    const encoded = encodeSavedBridgeParams(
      sender,
      paramsV2AllBytesEmpty,
      zapDataDefault
    )
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.sender).toEqual(sender)
    expect(decoded.paramsV2).toEqual(paramsV2AllBytesEmpty)
    expect(decoded.zapData).toEqual(zapDataDefault)
  })

  it('roundtrip encoding with negative exclusivity', () => {
    const encoded = encodeSavedBridgeParams(
      sender,
      paramsV2NegativeExclusivity,
      zapData
    )
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.sender).toEqual(sender)
    expect(decoded.paramsV2).toEqual(paramsV2NegativeExclusivity)
    expect(decoded.zapData).toEqual(zapData)
  })
})
