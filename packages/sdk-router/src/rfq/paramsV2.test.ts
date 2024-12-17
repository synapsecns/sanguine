import { BigNumber } from 'ethers'

import {
  BridgeParamsV2,
  SavedParamsV1,
  encodeSavedBridgeParams,
  decodeSavedBridgeParams,
} from './paramsV2'
import { ETH_USDC, ETH_USDT, ARB_USDC, ARB_USDT } from '../constants/testValues'

describe('paramsV2', () => {
  const ether = BigNumber.from(10).pow(18)

  const paramsV1: SavedParamsV1 = {
    originSender: ARB_USDC,
    destRecipient: ARB_USDT,
    destChainId: 1234,
    destToken: ETH_USDC,
    destAmount: ether.mul(2),
  }

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
    const encoded = encodeSavedBridgeParams(paramsV1, paramsV2)
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.paramsV1).toEqual(paramsV1)
    expect(decoded.paramsV2).toEqual(paramsV2)
  })

  it('roundtrip encoding with empty quoteId', () => {
    const encoded = encodeSavedBridgeParams(paramsV1, paramsV2QuoteIdEmpty)
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.paramsV1).toEqual(paramsV1)
    expect(decoded.paramsV2).toEqual(paramsV2QuoteIdEmpty)
  })

  it('roundtrip encoding with empty zapData', () => {
    const encoded = encodeSavedBridgeParams(paramsV1, paramsV2ZapDataEmpty)
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.paramsV1).toEqual(paramsV1)
    expect(decoded.paramsV2).toEqual(paramsV2ZapDataEmpty)
  })

  it('roundtrip encoding with empty quoteId and zapData', () => {
    const encoded = encodeSavedBridgeParams(paramsV1, paramsV2AllBytesEmpty)
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.paramsV1).toEqual(paramsV1)
    expect(decoded.paramsV2).toEqual(paramsV2AllBytesEmpty)
  })

  it('roundtrip encoding with negative exclusivity', () => {
    const encoded = encodeSavedBridgeParams(
      paramsV1,
      paramsV2NegativeExclusivity
    )
    const decoded = decodeSavedBridgeParams(encoded)
    expect(decoded.paramsV1).toEqual(paramsV1)
    expect(decoded.paramsV2).toEqual(paramsV2NegativeExclusivity)
  })
})
