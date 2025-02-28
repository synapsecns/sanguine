import request from 'supertest'
import express from 'express'

import tokenListRoute from '../routes/tokenListRoute'

const app = express()
app.use('/tokenList', tokenListRoute)

describe('Index Route', () => {
  it('should return a list of tokens with chain address mappings', async () => {
    const response = await request(app).get('/tokenList')

    expect(response.status).toBe(200)

    const keys = Object.keys(response.body)

    expect(keys.length).toBe(68)
    expect(response.body['ETH']['addresses']['1']).toBe(
      '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'
    )
    expect(response.body['SYN']['addresses']['1']).toBe(
      '0x0f2d719407fdbeff09d87557abb7232601fd9f29'
    )
  })
})
