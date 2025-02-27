import { isSameAddress } from './addressUtils'

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

    it('both empty', () => {
      expect(isSameAddress('', '')).toBe(false)
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
