import { AddressZero, Zero } from '@ethersproject/constants'
import {
  hexConcat,
  hexDataSlice,
  hexDataLength,
  hexZeroPad,
} from '@ethersproject/bytes'
import { BigNumber } from 'ethers'

import { Prettify } from '../../utils/types'

export const ZAP_DATA_VERSION = 1
export const AMOUNT_NOT_PRESENT = 0xffff

const OFFSET_AMOUNT_POSITION = 2
const OFFSET_FINAL_TOKEN = 4
const OFFSET_FORWARD_TO = 24
const OFFSET_MIN_FWD_AMOUNT = 44
const OFFSET_TARGET = 76
const OFFSET_PAYLOAD = 96

export type ZapDataV1 = {
  target: string
  payload: string
  amountPosition: number
  finalToken: string
  forwardTo: string
  minFinalAmount: BigNumber
}

export type PartialZapDataV1 = Prettify<Partial<ZapDataV1>>

export const encodeZapData = (zapData: PartialZapDataV1): string => {
  if (!zapData.target) {
    return '0x'
  }
  const {
    target,
    payload,
    amountPosition,
    finalToken,
    forwardTo,
    minFinalAmount,
  } = applyDefaultValues(zapData)
  return hexConcat([
    encodeUint16(ZAP_DATA_VERSION),
    encodeUint16(amountPosition),
    finalToken,
    forwardTo,
    encodeUint256(minFinalAmount),
    target,
    payload,
  ])
}

export const decodeZapData = (zapData: string): PartialZapDataV1 => {
  if (zapData === '0x') {
    return {}
  }
  if (hexDataLength(zapData) < OFFSET_PAYLOAD) {
    throw new Error('decodeZapData: zapData too short')
  }
  // Offsets of the fields in the packed ZapData struct
  // uint16   version                 [000 .. 002)
  // uint16   amountPosition          [002 .. 004)
  // address  finalToken              [004 .. 024)
  // address  forwardTo               [024 .. 044)
  // uint256  minFinalAmount          [044 .. 076)
  // address  target                  [076 .. 096)
  // bytes    payload                 [096 .. ***)
  const version = parseInt(hexDataSlice(zapData, 0, 2), 16)
  if (version !== ZAP_DATA_VERSION) {
    throw new Error('decodeZapData: unsupported version')
  }
  return {
    amountPosition: parseInt(
      hexDataSlice(zapData, OFFSET_AMOUNT_POSITION, OFFSET_FINAL_TOKEN),
      16
    ),
    finalToken: hexDataSlice(zapData, OFFSET_FINAL_TOKEN, OFFSET_FORWARD_TO),
    forwardTo: hexDataSlice(zapData, OFFSET_FORWARD_TO, OFFSET_MIN_FWD_AMOUNT),
    minFinalAmount: BigNumber.from(
      hexDataSlice(zapData, OFFSET_MIN_FWD_AMOUNT, OFFSET_TARGET)
    ),
    target: hexDataSlice(zapData, OFFSET_TARGET, OFFSET_PAYLOAD),
    payload: hexDataSlice(zapData, OFFSET_PAYLOAD),
  }
}

export const modifyMinFinalAmount = (
  zapData: string,
  newMinFinalAmount: BigNumber
): string => {
  const decoded = decodeZapData(zapData)
  return encodeZapData({
    ...decoded,
    minFinalAmount: newMinFinalAmount,
  })
}

export const applyDefaultValues = (zapData: PartialZapDataV1): ZapDataV1 => {
  return {
    target: zapData.target || AddressZero,
    payload: zapData.payload || '0x',
    amountPosition: zapData.amountPosition || AMOUNT_NOT_PRESENT,
    finalToken: zapData.finalToken || AddressZero,
    forwardTo: zapData.forwardTo || AddressZero,
    minFinalAmount: zapData.minFinalAmount || Zero,
  }
}

const encodeUint16 = (n: number): string => {
  return hexZeroPad('0x' + n.toString(16), 2)
}

const encodeUint256 = (n: BigNumber): string => {
  return hexZeroPad(n.toHexString(), 32)
}
