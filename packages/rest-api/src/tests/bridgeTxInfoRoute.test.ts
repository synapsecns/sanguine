import request from 'supertest'
import express from 'express'

import bridgeTxInfoRoute from '../routes/bridgeTxInfoRoute'

const app = express()
app.use('/bridgeTxInfo', bridgeTxInfoRoute)

describe('Bridge TX Info Route', () => {
  it('should return bridge transaction info for valid input', async () => {
    const response = await request(app).get('/bridgeTxInfo').query({
      fromChain: '1',
      toChain: '137',
      fromToken: 'USDC',
      toToken: 'USDC',
      amount: '1000',
      destAddress: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })

    expect(response.status).toBe(200)
    expect(Array.isArray(response.body)).toBe(true)
    expect(response.body.length).toBeGreaterThan(0)
    expect(response.body[0]).toHaveProperty('data')
    expect(response.body[0]).toHaveProperty(
      'to',
      '0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48'
    )
  }, 10_000)

  it('should return 400 for unsupported fromChain', async () => {
    const response = await request(app).get('/bridgeTxInfo').query({
      fromChain: '999',
      toChain: '137',
      fromToken: 'USDC',
      toToken: 'USDC',
      amount: '1000',
      destAddress: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Unsupported fromChain'
    )
  }, 10_000)

  it('should return 400 for unsupported toChain', async () => {
    const response = await request(app).get('/bridgeTxInfo').query({
      fromChain: '1',
      toChain: '999',
      fromToken: 'USDC',
      toToken: 'USDC',
      amount: '1000',
      destAddress: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('message', 'Unsupported toChain')
  }, 10_000)

  it('should return 400 for missing fromToken', async () => {
    const response = await request(app).get('/bridgeTxInfo').query({
      fromChain: '1',
      toChain: '137',
      toToken: 'USDC',
      amount: '1000',
      destAddress: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'fromToken')
  }, 10_000)

  it('should return 400 for missing amount', async () => {
    const response = await request(app).get('/bridgeTxInfo').query({
      fromChain: '1',
      toChain: '137',
      fromToken: 'USDC',
      toToken: 'USDC',
      destAddress: '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'amount')
  }, 10_000)

  it('should return 400 for missing destAddress', async () => {
    const response = await request(app).get('/bridgeTxInfo').query({
      fromChain: '1',
      toChain: '137',
      fromToken: 'USDC',
      toToken: 'USDC',
      amount: '1000',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'destAddress')
  }, 10_000)
})
