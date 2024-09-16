import request from 'supertest'
import express from 'express'

import swapRoute from '../routes/swapRoute'

const app = express()
app.use('/swap', swapRoute)

describe('Swap Route with Real Synapse Service', () => {
  it('should return a real swap quote for valid input, 1000 USDC', async () => {
    const response = await request(app).get('/swap').query({
      chain: '1',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48', // USDC on Ethereum
      toToken: '0x6B175474E89094C44Da98b954EedeAC495271d0F', // DAI on Ethereum
      amount: '1000',
    })

    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('maxAmountOut')
    expect(response.body).toHaveProperty('routerAddress')
    expect(response.body).toHaveProperty('query')
  }, 10_000)

  it('should return 400 for unsupported chain, with error message', async () => {
    const response = await request(app).get('/swap').query({
      chain: '111',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: '0x6B175474E89094C44Da98b954EedeAC495271d0F',
      amount: '1000',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('message', 'Unsupported chain')
  }, 10_000)

  it('should return 400 for invalid toToken address, with error message', async () => {
    const response = await request(app).get('/swap').query({
      chain: '1',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: 'invalid_address',
      amount: '1000',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Invalid toToken address'
    )
  }, 10_000)

  it('should return 400 for missing amount, with error message', async () => {
    const response = await request(app).get('/swap').query({
      chain: '1',
      fromToken: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      toToken: '0x6B175474E89094C44Da98b954EedeAC495271d0F',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'amount')
  }, 10_000)
})
