import { utils } from 'ethers'

import {
  handleNativeToken,
  handleParams,
  isNativeToken,
  isSameAddress,
} from './addressUtils'

describe('isSameAddress', () => {
  const lowerCaseAlice = '0x0123456789abcdef0123456789abcdef01234567'
  const checkSumdAlice = '0x0123456789abcDEF0123456789abCDef01234567'
  const upperCaseAlice = '0x0123456789ABCDEF0123456789ABCDEF01234567'

  const lowerCaseBob = '0x0123456789abcdef0123456789abcdef01234568'
  const checkSumdBob = '0x0123456789ABCDeF0123456789aBcdEF01234568'
  const upperCaseBob = '0x0123456789ABCDEF0123456789ABCDEF01234568'

  describe('True when the addresses are the same', () => {
    it('Both lowercase', () => {
      expect(isSameAddress(lowerCaseAlice, lowerCaseAlice)).toBe(true)
      expect(isSameAddress(lowerCaseBob, lowerCaseBob)).toBe(true)
    })

    it('Both checksummed', () => {
      expect(isSameAddress(checkSumdAlice, checkSumdAlice)).toBe(true)
      expect(isSameAddress(checkSumdBob, checkSumdBob)).toBe(true)
    })

    it('Both uppercase', () => {
      expect(isSameAddress(upperCaseAlice, upperCaseAlice)).toBe(true)
      expect(isSameAddress(upperCaseBob, upperCaseBob)).toBe(true)
    })

    it('Lowercase and checksummed', () => {
      expect(isSameAddress(lowerCaseAlice, checkSumdAlice)).toBe(true)
      expect(isSameAddress(checkSumdAlice, lowerCaseAlice)).toBe(true)
      expect(isSameAddress(lowerCaseBob, checkSumdBob)).toBe(true)
      expect(isSameAddress(checkSumdBob, lowerCaseBob)).toBe(true)
    })

    it('Lowercase and uppercase', () => {
      expect(isSameAddress(lowerCaseAlice, upperCaseAlice)).toBe(true)
      expect(isSameAddress(upperCaseAlice, lowerCaseAlice)).toBe(true)
      expect(isSameAddress(lowerCaseBob, upperCaseBob)).toBe(true)
      expect(isSameAddress(upperCaseBob, lowerCaseBob)).toBe(true)
    })

    it('Checksummed and uppercase', () => {
      expect(isSameAddress(checkSumdAlice, upperCaseAlice)).toBe(true)
      expect(isSameAddress(upperCaseAlice, checkSumdAlice)).toBe(true)
      expect(isSameAddress(checkSumdBob, upperCaseBob)).toBe(true)
      expect(isSameAddress(upperCaseBob, checkSumdBob)).toBe(true)
    })
  })

  describe('False when the addresses are different', () => {
    it('Both lowercase', () => {
      expect(isSameAddress(lowerCaseAlice, lowerCaseBob)).toBe(false)
      expect(isSameAddress(lowerCaseBob, lowerCaseAlice)).toBe(false)
    })

    it('Both checksummed', () => {
      expect(isSameAddress(checkSumdAlice, checkSumdBob)).toBe(false)
      expect(isSameAddress(checkSumdBob, checkSumdAlice)).toBe(false)
    })

    it('Both uppercase', () => {
      expect(isSameAddress(upperCaseAlice, upperCaseBob)).toBe(false)
      expect(isSameAddress(upperCaseBob, upperCaseAlice)).toBe(false)
    })

    it('Lowercase and checksummed', () => {
      expect(isSameAddress(lowerCaseAlice, checkSumdBob)).toBe(false)
      expect(isSameAddress(checkSumdBob, lowerCaseAlice)).toBe(false)
    })

    it('Lowercase and uppercase', () => {
      expect(isSameAddress(lowerCaseAlice, upperCaseBob)).toBe(false)
      expect(isSameAddress(upperCaseBob, lowerCaseAlice)).toBe(false)
    })

    it('Checksummed and uppercase', () => {
      expect(isSameAddress(checkSumdAlice, upperCaseBob)).toBe(false)
      expect(isSameAddress(upperCaseBob, checkSumdAlice)).toBe(false)
    })
  })

  describe('False when one of the addresses is undefined', () => {
    it('single undefined', () => {
      expect(isSameAddress(undefined, lowerCaseAlice)).toBe(false)
      expect(isSameAddress(lowerCaseAlice, undefined)).toBe(false)
    })

    it('both undefined', () => {
      expect(isSameAddress(undefined, undefined)).toBe(false)
    })
  })

  describe('False when one of the addresses is empty', () => {
    it('single empty', () => {
      expect(isSameAddress('', lowerCaseAlice)).toBe(false)
      expect(isSameAddress(lowerCaseAlice, '')).toBe(false)
    })
  })

  describe('False when one of the addresses is null', () => {
    it('single null', () => {
      expect(isSameAddress(null as any, lowerCaseAlice)).toBe(false)
      expect(isSameAddress(lowerCaseAlice, null as any)).toBe(false)
    })

    it('both null', () => {
      expect(isSameAddress(null as any, null as any)).toBe(false)
    })
  })
})

describe('Native token address utils', () => {
  const empty = ''
  const zero = '0x0000000000000000000000000000000000000000'
  const eth = '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee'
  const ethCheckSummed = utils.getAddress(eth)

  const random = '0x0000000000000000000000000000000000000001'
  const nonEth = '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef'

  describe('handleNativeToken', () => {
    it('Empty address: ETH', () => {
      expect(handleNativeToken(empty)).toEqual(ethCheckSummed)
    })

    it('Random address: preserved', () => {
      expect(handleNativeToken(random)).toEqual(random)
    })

    it('Non-ETH address: preserved', () => {
      expect(handleNativeToken(nonEth)).toEqual(nonEth)
    })

    it('Zero address: ETH', () => {
      expect(handleNativeToken(zero)).toEqual(ethCheckSummed)
    })

    it('ETH lowercase: ETH', () => {
      expect(handleNativeToken(eth)).toEqual(ethCheckSummed)
    })

    it('ETH uppercase: ETH', () => {
      expect(handleNativeToken(eth.toUpperCase())).toEqual(ethCheckSummed)
    })
  })

  describe('isNativeToken', () => {
    it('Empty address: true', () => {
      expect(isNativeToken(empty)).toBe(true)
    })

    it('Random address: false', () => {
      expect(isNativeToken(random)).toBe(false)
    })

    it('Non-ETH address: false', () => {
      expect(isNativeToken(nonEth)).toBe(false)
    })

    it('Zero address: true', () => {
      expect(isNativeToken(zero)).toBe(true)
    })

    it('ETH lowercase: true', () => {
      expect(isNativeToken(eth)).toBe(true)
    })

    it('ETH uppercase: true', () => {
      expect(isNativeToken(eth.toUpperCase())).toBe(true)
    })

    it('undefined: false', () => {
      expect(isNativeToken(undefined)).toBe(false)
    })
  })

  describe('handleParams', () => {
    const restParams = {
      abc: 'def',
      def: 1234,
    }

    const ethLike = [empty, zero, eth, ethCheckSummed]
    const nonEthLike = [random, nonEth]

    it('ETH & ETH', () => {
      ethLike.forEach((fromToken) => {
        ethLike.forEach((toToken) => {
          expect(handleParams({ ...restParams, fromToken, toToken })).toEqual({
            ...restParams,
            fromToken: ethCheckSummed,
            toToken: ethCheckSummed,
          })
        })
      })
    })

    it('ETH & non-ETH', () => {
      ethLike.forEach((fromToken) => {
        nonEthLike.forEach((toToken) => {
          expect(handleParams({ ...restParams, fromToken, toToken })).toEqual({
            ...restParams,
            fromToken: ethCheckSummed,
            toToken,
          })
        })
      })
    })

    it('non-ETH & ETH', () => {
      nonEthLike.forEach((fromToken) => {
        ethLike.forEach((toToken) => {
          expect(handleParams({ ...restParams, fromToken, toToken })).toEqual({
            ...restParams,
            fromToken,
            toToken: ethCheckSummed,
          })
        })
      })
    })

    it('non-ETH & non-ETH', () => {
      nonEthLike.forEach((fromToken) => {
        nonEthLike.forEach((toToken) => {
          expect(handleParams({ ...restParams, fromToken, toToken })).toEqual({
            ...restParams,
            fromToken,
            toToken,
          })
        })
      })
    })
  })
})
