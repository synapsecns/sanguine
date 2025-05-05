import {
  handleNativeToken,
  ETH_NATIVE_TOKEN_ADDRESS,
} from './handleNativeToken'

describe('test empty string', () => {
  it('is correct for zero', () => {
    expect(handleNativeToken('')).toEqual(ETH_NATIVE_TOKEN_ADDRESS)
  })
  it('test random addr', () => {
    const testAddr = '0x123456789'
    expect(handleNativeToken(testAddr)).toEqual(testAddr)
  })
})
