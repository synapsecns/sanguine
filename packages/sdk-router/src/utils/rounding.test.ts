import { roundDown, roundUp } from './rounding'

describe('rounding', () => {
  describe('roundDown', () => {
    it('rounds down to the nearest multiple of precision', () => {
      expect(roundDown(17, 5)).toBe(15)
      expect(roundDown(23, 10)).toBe(20)
      expect(roundDown(99, 25)).toBe(75)
    })

    it('returns the same value if already a multiple of precision', () => {
      expect(roundDown(20, 5)).toBe(20)
      expect(roundDown(100, 10)).toBe(100)
      expect(roundDown(0, 5)).toBe(0)
    })

    it('handles decimal precision', () => {
      expect(roundDown(2.7, 0.5)).toBe(2.5)
      expect(roundDown(3.14, 0.1)).toBe(3.1)
    })

    it('handles values smaller than precision', () => {
      expect(roundDown(3, 5)).toBe(0)
      expect(roundDown(9, 10)).toBe(0)
    })
  })

  describe('roundUp', () => {
    it('rounds up to the nearest multiple of precision', () => {
      expect(roundUp(17, 5)).toBe(20)
      expect(roundUp(23, 10)).toBe(30)
      expect(roundUp(76, 25)).toBe(100)
    })

    it('returns the same value if already a multiple of precision', () => {
      expect(roundUp(20, 5)).toBe(20)
      expect(roundUp(100, 10)).toBe(100)
      expect(roundUp(0, 5)).toBe(0)
    })

    it('handles decimal precision', () => {
      expect(roundUp(2.3, 0.5)).toBe(2.5)
      expect(roundUp(3.11, 0.1)).toBe(3.2)
    })

    it('handles values smaller than precision', () => {
      expect(roundUp(3, 5)).toBe(5)
      expect(roundUp(1, 10)).toBe(10)
    })
  })
})
