import request from 'supertest'
import express from 'express'
import destinationTokensRoute from '../routes/destinationTokensRoute'

const app = express()
app.use('/destinationTokens', destinationTokensRoute)

describe('destinatonTokens Route', () => {
  it('should return destination tokens for valid input', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromChain: '1',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
    })

    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBeGreaterThan(0)
    expect(response.body[0]).toHaveProperty('symbol')
    expect(response.body[0]).toHaveProperty('address')
    expect(response.body[0]).toHaveProperty('chainId')
  })

  it('should return 400 for unsupported fromChain', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromChain: '999',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Unsupported fromChain'
    )
  })

  it('should return 400 for invalid fromToken address', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromChain: '1',
      fromToken: 'invalid_address',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Invalid fromToken address'
    )
  })

  it('should return 400 for token not supported by Synapse', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromChain: '1',
      fromToken: '0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Unsupported fromToken address'
    )
  })

  it('should return 400 for token not supported on specified chain', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromChain: '10',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Token not supported on specified chain'
    )
  })

  it('should return 400 for missing fromChain', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'fromChain is required'
    )
  })

  it('should return 400 for missing fromToken', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromChain: '1',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'fromToken is required'
    )
  })
})
