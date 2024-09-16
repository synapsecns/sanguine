import request from 'supertest'
import express from 'express'

import swapTxInfoRoute from '../routes/swapTxInfoRoute'

const app = express()
app.use('/swapTxInfo', swapTxInfoRoute)

describe('Swap TX Info Route with Real Synapse Service', () => {
  it('should return transaction info for valid input, 1000 USDC to DAI', async () => {
    const response = await request(app).get('/swapTxInfo').query({
      chain: '1',
      fromToken: 'USDC',
      toToken: 'DAI',
      amount: '1000',
    })
    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('data')
    expect(response.body).toHaveProperty('to')
    expect(response.body).toHaveProperty('value')
  }, 10_000)

  it('should return 400 for unsupported chain, with error message', async () => {
    const response = await request(app).get('/swapTxInfo').query({
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
    const response = await request(app).get('/swapTxInfo').query({
      chain: '1',
      fromToken: 'USDC',
      amount: '1000',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'toToken')
  }, 10_000)

  it('should return 400 for missing amount, with error message', async () => {
    const response = await request(app).get('/swapTxInfo').query({
      chain: '1',
      fromToken: 'USDC',
      toToken: 'DAI',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'amount')
  }, 10_000)
})
