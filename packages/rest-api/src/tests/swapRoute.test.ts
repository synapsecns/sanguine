import request from 'supertest'
import express from 'express'

import swapRoute from '../routes/swapRoute'

const app = express()
app.use('/swap', swapRoute)

describe('Swap Route with Real Synapse Service', () => {
  it('should return a real swap quote for valid input, 1000 USDC', async () => {
    const response = await request(app).get('/swap').query({
      chain: '1',
      fromToken: 'USDC',
      toToken: 'DAI',
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
      fromToken: 'USDC',
      toToken: 'DAI',
      amount: '1000',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('message', 'Unsupported chain')
    expect(response.body.error).toHaveProperty('field', 'chain')
  }, 10_000)

  it('should return 400 for missing toToken, with error message', async () => {
    const response = await request(app).get('/swap').query({
      chain: '1',
      fromToken: 'USDC',
      amount: '1000',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'toToken')
  }, 10_000)

  it('should return 400 for missing amount, with error message', async () => {
    const response = await request(app).get('/swap').query({
      chain: '1',
      fromToken: 'USDC',
      toToken: 'DAI',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'amount')
  }, 10_000)
})
