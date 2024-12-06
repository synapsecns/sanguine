import { BigNumber } from 'ethers'

import {
  BridgeParamsV2,
  encodeSavedBridgeParams,
  decodeSavedBridgeParams,
} from './paramsV2'
import { ETH_USDC, ETH_USDT } from '../constants/testValues'

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

  it('roundtrip encoding', () => {
    const encoded = encodeSavedBridgeParams(sender, paramsV2)
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.sender).toEqual(sender)
    expect(decoded.paramsV2).toEqual(paramsV2)
  })

  it('roundtrip encoding with empty quoteId', () => {
    const encoded = encodeSavedBridgeParams(sender, paramsV2QuoteIdEmpty)
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.sender).toEqual(sender)
    expect(decoded.paramsV2).toEqual(paramsV2QuoteIdEmpty)
  })

  it('roundtrip encoding with empty zapData', () => {
    const encoded = encodeSavedBridgeParams(sender, paramsV2ZapDataEmpty)
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.sender).toEqual(sender)
    expect(decoded.paramsV2).toEqual(paramsV2ZapDataEmpty)
  })

  it('roundtrip encoding with empty quoteId and zapData', () => {
    const encoded = encodeSavedBridgeParams(sender, paramsV2AllBytesEmpty)
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.sender).toEqual(sender)
    expect(decoded.paramsV2).toEqual(paramsV2AllBytesEmpty)
  })

  it('roundtrip encoding with negative exclusivity', () => {
    const encoded = encodeSavedBridgeParams(sender, paramsV2NegativeExclusivity)
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.sender).toEqual(sender)
    expect(decoded.paramsV2).toEqual(paramsV2NegativeExclusivity)
  })
})
