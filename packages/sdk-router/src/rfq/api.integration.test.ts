import { getAllQuotes } from './api'

global.fetch = require('node-fetch')

// Live quotes should not run in the default test suite.
describe.skip('getAllQuotes', () => {
  it('Integration test', async () => {
    const result = await getAllQuotes()
    // console.log('Current quotes: ' + JSON.stringify(result, null, 2))
    expect(result.length).toBeGreaterThan(0)
  })
})
