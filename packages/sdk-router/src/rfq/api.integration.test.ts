import { getAllQuotes } from './api'

global.fetch = require('node-fetch')

// Retry the flaky tests up to 3 times
jest.retryTimes(3)

describe('getAllQuotes', () => {
  it('Integration test', async () => {
    const result = await getAllQuotes()
    // console.log('Current quotes: ' + JSON.stringify(result, null, 2))
    expect(result.length).toBeGreaterThan(0)
  })
})
