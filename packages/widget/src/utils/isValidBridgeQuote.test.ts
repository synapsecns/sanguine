import { isValidBridgeQuote } from '@/utils/isValidBridgeQuote'

const createQuote = (nativeFee: unknown) => ({
  nativeFee,
})

describe('isValidBridgeQuote', () => {
  it.each([
    ['zero string nativeFee', '0', true],
    ['positive string nativeFee', '77', true],
    ['zero bigint nativeFee', 0n, true],
    ['positive bigint nativeFee', 77n, true],
    ['missing nativeFee', undefined, false],
    ['malformed nativeFee', 'abc', false],
    ['decimal nativeFee', '1.5', false],
    ['negative string nativeFee', '-1', false],
    ['negative bigint nativeFee', -1n, false],
  ])('returns %s', (_scenario, nativeFee, expected) => {
    expect(isValidBridgeQuote(createQuote(nativeFee))).toBe(expected)
  })
})
