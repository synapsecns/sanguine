import { BigNumber } from '@ethersproject/bignumber'

import {
  calculateDeadline,
  getDestinationDeadline,
  getOriginDeadline,
} from './deadlines'

describe('deadlines', () => {
  // Something good happened on this day
  Date.now = jest.fn(() => Date.parse('2021-08-29'))

  it('calculates correct deadlines', () => {
    const seconds = 1337
    const deadline = calculateDeadline(seconds)
    const now = Math.floor(Date.now() / 1000)
    expect(deadline.toNumber()).toBe(now + seconds)
  })

  describe('getOriginDeadline', () => {
    it('returns the deadline if it is defined', () => {
      const deadline = BigNumber.from(1337)
      expect(getOriginDeadline(deadline)).toBe(deadline)
    })

    it('Uses 10 minutes if deadline is undefined', () => {
      const deadline = getOriginDeadline()
      const now = Math.floor(Date.now() / 1000)
      expect(deadline.toNumber()).toBe(now + 10 * 60)
    })
  })

  describe('getDestinationDeadline', () => {
    it('returns the deadline if it is defined', () => {
      const deadline = BigNumber.from(1337)
      expect(getDestinationDeadline(deadline)).toBe(deadline)
    })

    it('Uses 1 week if deadline is undefined', () => {
      const deadline = getDestinationDeadline()
      const now = Math.floor(Date.now() / 1000)
      expect(deadline.toNumber()).toBe(now + 7 * 24 * 60 * 60)
    })
  })
})
