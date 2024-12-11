import { BigNumber } from 'ethers'

import { ETH_USDC, ETH_USDT, NATIVE_ADDRESS } from '../constants/testValues'
import {
  StepParams,
  decodeStepParams,
  encodeStepParams,
  extractSingleZapData,
} from './steps'

describe('Steps', () => {
  const ether = BigNumber.from(10).pow(18)

  const param0: StepParams = {
    token: NATIVE_ADDRESS,
    amount: ether.mul(1),
    msgValue: ether.mul(2),
    zapData: '0x',
  }

  const param1: StepParams = {
    token: ETH_USDC,
    amount: ether.mul(3),
    msgValue: ether.mul(4),
    zapData: '0xdeadbeef',
  }

  const param2: StepParams = {
    token: ETH_USDT,
    amount: ether.mul(5),
    msgValue: ether.mul(6),
    zapData: '0x00112233445566778899aabbccddeeff',
  }

  it('roundtrip with one step', () => {
    const data = encodeStepParams([param0])
    expect(decodeStepParams(data)).toEqual([param0])
  })

  it('roundtrip with two steps', () => {
    const data = encodeStepParams([param0, param1])
    expect(decodeStepParams(data)).toEqual([param0, param1])
  })

  it('roundtrip with three steps', () => {
    const data = encodeStepParams([param0, param1, param2])
    expect(decodeStepParams(data)).toEqual([param0, param1, param2])
  })

  describe('extractSingleZapData', () => {
    it('no steps', () => {
      expect(extractSingleZapData([])).toEqual('0x')
    })

    it('single step', () => {
      expect(extractSingleZapData([param0])).toEqual('0x')
      expect(extractSingleZapData([param1])).toEqual('0xdeadbeef')
      expect(extractSingleZapData([param2])).toEqual(
        '0x00112233445566778899aabbccddeeff'
      )
    })

    it('multiple steps', () => {
      expect(() => extractSingleZapData([param0, param1])).toThrowError(
        'extractSingleZapData: more than one step'
      )
      expect(() => extractSingleZapData([param0, param2])).toThrowError(
        'extractSingleZapData: more than one step'
      )
      expect(() => extractSingleZapData([param1, param2])).toThrowError(
        'extractSingleZapData: more than one step'
      )
    })
  })
})
