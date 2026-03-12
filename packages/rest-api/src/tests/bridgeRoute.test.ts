import request from 'supertest'
import express from 'express'

import bridgeRoute from '../routes/bridgeRoute'
import { NativeGasAddress, ZeroAddress } from '../constants'
import { USDC } from '../constants/bridgeable'
import { UNSUPPORTED_CHAIN } from './testConstants'

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

  it('should return bridge quotes for valid originUserAddress', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '10',
      fromToken: USDC.addresses[1],
      toToken: USDC.addresses[10],
      amount: '1000',
      originUserAddress: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
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

  it('should return 400 for invalid originUserAddress', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '10',
      fromToken: USDC.addresses[1],
      toToken: USDC.addresses[10],
      amount: '1000',
      originUserAddress: 'invalid_address',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Invalid originUserAddress address'
    )
  }, 15000)

  it('should return 400 for unsupported route', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '10',
      fromToken: NativeGasAddress,
      toToken: USDC.addresses[10],
      amount: '10',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'No valid route exists for the chain/token combination'
    )
  })

  it('should return 400 for unsupported fromChain, with error message', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: UNSUPPORTED_CHAIN,
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
  })

  it('should return 400 for unsupported toChain, with error message', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: UNSUPPORTED_CHAIN,
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: '0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85',
      amount: '1000',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('message', 'Unsupported toChain')
  })

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
  })

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
  })

  it('should return 400 for missing amount, with error message', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '10',
      fromToken: USDC.addresses[1],
      toToken: USDC.addresses[10],
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'amount')
  })

  it('should return bridge quotes with callData when originUserAddress and destAddress are provided', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '10',
      fromToken: USDC.addresses[1],
      toToken: USDC.addresses[10],
      amount: '1000',
      destAddress: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
      originUserAddress: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })

    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBeGreaterThan(0)
    expect(response.body[0]).toHaveProperty('callData')
    expect(response.body[0].callData).toHaveProperty('to')
    expect(response.body[0].callData).toHaveProperty('data')
    expect(response.body[0].callData).toHaveProperty('value')
  }, 15000)

  it('should return bridge quotes without callData when destAddress is not provided', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '10',
      fromToken: USDC.addresses[1],
      toToken: USDC.addresses[10],
      amount: '1000',
      originUserAddress: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })

    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBeGreaterThan(0)
    expect(response.body[0].callData).toBeNull()
  }, 15000)

  it('should return bridge quotes without callData when originUserAddress is not provided', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '10',
      fromToken: USDC.addresses[1],
      toToken: USDC.addresses[10],
      amount: '1000',
      destAddress: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })

    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBeGreaterThan(0)
    expect(response.body[0].callData).toBeNull()
  }, 15000)

  it('should return 400 for invalid destAddress', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '1',
      toChain: '10',
      fromToken: USDC.addresses[1],
      toToken: USDC.addresses[10],
      amount: '1000',
      destAddress: 'invalid_address',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('message', 'Invalid destAddress')
  }, 15000)

  it('should return 404 when no bridge routes found for small amount (0.001 USDC from Arbitrum to Ethereum)', async () => {
    const response = await request(app).get('/bridge').query({
      fromChain: '42161', // Arbitrum
      toChain: '1', // Ethereum
      fromToken: USDC.addresses[42161],
      toToken: USDC.addresses[1],
      amount: '0.001',
    })

    expect(response.status).toBe(404)
    expect(response.body).toHaveProperty('error', 'No bridge routes found')
  }, 15000)
})
