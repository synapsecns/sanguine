import request from 'supertest'
import express from 'express'

import swapTxInfoRoute from '../routes/swapTxInfoRoute'

const app = express()
app.use('/swapTxInfo', swapTxInfoRoute)

describe('Swap TX Info Route with Real Synapse Service', () => {
  it('should return transaction info for valid input, 1000 USDC to DAI', async () => {
    const response = await request(app).get('/swapTxInfo').query({
      chain: '1',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48', // USDC on Ethereum
      toToken: '0x6B175474E89094C44Da98b954EedeAC495271d0F', // DAI on Ethereum
      amount: '1000',
      address: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })
    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('data')
    expect(response.body).toHaveProperty('to')
    expect(response.body).toHaveProperty('value')
  }, 10_000)

  it('should return 400 for invalid address, with error message', async () => {
    const response = await request(app).get('/swapTxInfo').query({
      chain: '1',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: '0x6B175474E89094C44Da98b954EedeAC495271d0F',
      amount: '1000',
      address: 'invalid_address',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Invalid Ethereum address'
    )
  }, 10_000)

  it('should return 400 for unsupported chain, with error message', async () => {
    const response = await request(app).get('/swapTxInfo').query({
      chain: '111',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: '0x6B175474E89094C44Da98b954EedeAC495271d0F',
      amount: '1000',
      address: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('message', 'Unsupported chain')
  }, 10_000)

  it('should return 400 for invalid toToken address, with error message', async () => {
    const response = await request(app).get('/swapTxInfo').query({
      chain: '1',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: 'invalid_address',
      amount: '1000',
      address: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Invalid toToken address'
    )
  }, 10_000)

  it('should return 400 for token not supported on specified chain', async () => {
    const response = await request(app).get('/swapTxInfo').query({
      chain: '1',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: '0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F', // SNX on Ethereum (Not supported)
      amount: '1000',
      address: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Invalid toToken address'
    )
  }, 10_000)

  it('should return 400 for missing amount, with error message', async () => {
    const response = await request(app).get('/swapTxInfo').query({
      chain: '1',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: '0x6B175474E89094C44Da98b954EedeAC495271d0F',
      address: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'amount')
  }, 10_000)
})
