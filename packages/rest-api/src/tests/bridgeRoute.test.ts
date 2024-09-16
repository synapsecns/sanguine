import request from 'supertest'
import express from 'express'

import bridgeRoute from '../routes/bridgeRoute'

const app = express()
app.use('/bridge', bridgeRoute)

describe('Bridge Route with Real Synapse Service', () => {
  it('should return bridge quotes for valid input, 1000 USDC from Ethereum to Polygon', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '137',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48', // USDC on Ethereum
      toToken: '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174', // USDC.e on Polygon
      amount: '1000',
    })
    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBeGreaterThan(0)
    expect(response.body[0]).toHaveProperty('maxAmountOutStr')
    expect(response.body[0]).toHaveProperty('bridgeFeeFormatted')
  }, 15000)

  it('should return 400 for unsupported fromChain, with error message', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '999',
      toChain: '137',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174',
      amount: '1000',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Unsupported fromChain'
    )
  }, 10000)

  it('should return 400 for unsupported toChain, with error message', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '999',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174',
      amount: '1000',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('message', 'Unsupported toChain')
  }, 10000)

  it('should return 400 for invalid fromToken address, with error message', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '137',
      fromToken: 'invalid_address',
      toToken: '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174',
      amount: '1000',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Invalid fromToken address'
    )
  }, 10000)

  it('should return 400 for missing amount, with error message', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '137',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'amount')
  }, 10000)
})
