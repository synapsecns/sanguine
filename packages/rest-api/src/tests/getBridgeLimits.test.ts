import request from 'supertest'
import express from 'express'

import getBridgeLimitsRoute from '../routes/getBridgeLimitsRoute'

const app = express()
app.use('/getBridgeLimits', getBridgeLimitsRoute)

describe('Get Bridge Limits Route', () => {
  it('should return min/max origin amounts for valid input', async () => {
    const response = await request(app).get('/getBridgeLimits').query({
      fromChain: 1,
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toChain: 10,
      toToken: '0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85',
    })

    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('maxOriginAmount')
    expect(response.body).toHaveProperty('minOriginAmount')
  }, 10_000)

  it('should return 400 for unsupported fromChain', async () => {
    const response = await request(app).get('/getBridgeLimits').query({
      fromChain: '999',
      toChain: '137',
      fromToken: '0x176211869cA2b568f2A7D4EE941E073a821EE1ff',
      toToken: '0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Unsupported fromChain'
    )
  }, 10_000)

  it('should return 400 for unsupported ', async () => {
    const response = await request(app).get('/getBridgeLimits').query({
      fromChain: '999',
      toChain: '137',
      fromToken: '0x176211869cA2b568f2A7D4EE941E073a821EE1ff',
      toToken: '0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Unsupported fromChain'
    )
  }, 10_000)

  it('should return 400 for unsupported toChain', async () => {
    const response = await request(app).get('/getBridgeLimits').query({
      fromChain: '137',
      toChain: '999',
      fromToken: '0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359',
      toToken: '0x176211869cA2b568f2A7D4EE941E073a821EE1ff',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('message', 'Unsupported toChain')
  }, 10_000)

  it('should return 400 for missing fromToken', async () => {
    const response = await request(app).get('/getBridgeLimits').query({
      fromChain: '1',
      toChain: '137',
      toToken: '0x176211869cA2b568f2A7D4EE941E073a821EE1ff',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'fromToken')
  }, 10_000)

  it('should return 400 for missing toToken', async () => {
    const response = await request(app).get('/getBridgeLimits').query({
      fromChain: '1',
      toChain: '137',
      fromToken: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'toToken')
  }, 10_000)
})
