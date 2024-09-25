import request from 'supertest'
import express from 'express'

import bridgeRoute from '../routes/bridgeRoute'
import { NativeGasAddress, ZeroAddress } from '../constants'
import { USDC } from '../constants/bridgeable'

const app = express()
app.use('/bridge', bridgeRoute)

describe('Bridge Route with Real Synapse Service', () => {
  it('should return bridge quotes for valid input, 1000 USDC from Ethereum to Optimism', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '10',
      fromToken: USDC.addresses[1],
      toToken: USDC.addresses[10],
      amount: '1000',
    })
    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBeGreaterThan(0)
    expect(response.body[0]).toHaveProperty('maxAmountOutStr')
    expect(response.body[0]).toHaveProperty('bridgeFeeFormatted')
  }, 15000)

  it('should return bridge quotes for ZeroAddress', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '10',
      fromToken: ZeroAddress,
      toToken: ZeroAddress,
      amount: '10',
    })
    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBeGreaterThan(0)
    expect(response.body[0]).toHaveProperty('maxAmountOutStr')
    expect(response.body[0]).toHaveProperty('bridgeFeeFormatted')
  }, 15000)

  it('should return bridge quotes for NativeGasAddress', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '10',
      fromToken: NativeGasAddress,
      toToken: NativeGasAddress,
      amount: '10',
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
      toToken: '0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85',
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

  it('should return 400 for token not supported on specified chain, with error message', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '137',
      fromToken: '0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F', // SNX on Ethereum (Not supported)
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
      toChain: '10',
      fromToken: USDC.addresses[1],
      toToken: USDC.addresses[10],
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'amount')
  }, 10000)
})
