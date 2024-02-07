import { BigNumber } from '@ethersproject/bignumber'

import { applyOptionalDeadline, calculateDeadline } from './deadlines'

describe('deadlines', () => {
  // Something good happened on this day
  Date.now = jest.fn(() => Date.parse('2021-08-29'))

  describe('calculateDeadline', () => {
    it('calculates correct deadlines', () => {
      const seconds = 1337
      const deadline = calculateDeadline(seconds)
      const now = Math.floor(Date.now() / 1000)
      expect(deadline.toNumber()).toBe(now + seconds)
    })
  })

  describe('applyOptionalDeadline', () => {
    it('returns the deadline if it is defined', () => {
      const deadline = BigNumber.from(1337)
      expect(applyOptionalDeadline(deadline, 1234)).toBe(deadline)
    })

    it('applies the default period if the deadline is undefined', () => {
      const deadline = applyOptionalDeadline(undefined, 1234)
      const now = Math.floor(Date.now() / 1000)
      expect(deadline.toNumber()).toBe(now + 1234)
    })
  })
})
