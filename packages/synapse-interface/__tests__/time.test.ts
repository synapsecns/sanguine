import { formatCompactDuration } from '@/utils/time'

describe('formatCompactDuration', () => {
  it.each([
    [0, '0s'],
    [1, '1s'],
    [59, '59s'],
    [60, '1m0s'],
    [61, '1m1s'],
    [90, '1m30s'],
    [1290, '21m30s'],
  ])('formats %s seconds as %s', (durationInSeconds, expectedOutput) => {
    expect(formatCompactDuration(durationInSeconds)).toBe(expectedOutput)
  })

  it('supports localized compact suffixes', () => {
    expect(
      formatCompactDuration(90, {
        minute: '分',
        second: '秒',
      })
    ).toBe('1分30秒')
  })
})
