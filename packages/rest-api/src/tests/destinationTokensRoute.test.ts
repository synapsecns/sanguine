import request from 'supertest'
import express from 'express'

import destinationTokensRoute from '../routes/destinationTokensRoute'
import { NativeGasAddress, ZeroAddress } from '../constants'
import { USDC, USDT } from '../constants/bridgeable'

const app = express()
app.use('/destinationTokens', destinationTokensRoute)

describe('destinatonTokens Route', () => {
  it('should return destination tokens for valid input', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromChain: '1',
      fromToken: USDC.addresses[1],
    })

    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBeGreaterThan(0)
    expect(response.body[0]).toHaveProperty('symbol')
    expect(response.body[0]).toHaveProperty('address')
    expect(response.body[0]).toHaveProperty('chainId')
  })

  it('should return destination tokens for valid gas Tokens', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromChain: '1',
      fromToken: NativeGasAddress,
    })

    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBeGreaterThan(0)
    expect(response.body[0]).toHaveProperty('symbol')
    expect(response.body[0]).toHaveProperty('address')
    expect(response.body[0]).toHaveProperty('chainId')
  })

  it('should return destination tokens for valid gas Tokens, ZeroAddress', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromChain: '1',
      fromToken: ZeroAddress,
    })

    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBeGreaterThan(0)
    expect(response.body[0]).toHaveProperty('symbol')
    expect(response.body[0]).toHaveProperty('address')
    expect(response.body[0]).toHaveProperty('chainId')
  })

  it('should return precisely the number of destination tokens', async () => {
    // 'USDC-534352': [ 'USDC-1', 'USDC-10', 'USDC-8453', 'USDC-42161', 'USDC-59144' ]

    const response = await request(app).get('/destinationTokens').query({
      fromChain: '534352',
      fromToken: USDC.addresses[534352],
    })

    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBe(5)
    expect(response.body[0]).toHaveProperty('symbol')
    expect(response.body[0]).toHaveProperty('address')
    expect(response.body[0]).toHaveProperty('chainId')
  })

  it('should return destination tokens for non-checksummed address', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromChain: '43114',
      fromToken: USDT.addresses[43114].toLowerCase(),
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
      fromToken: USDC.addresses[1],
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Token not supported on specified chain'
    )
  })

  it('should return 400 for missing fromChain', async () => {
    const response = await request(app).get('/destinationTokens').query({
      fromToken: USDC.addresses[1],
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
