import request from 'supertest'
import express from 'express'

import bridgeLimitsRoute from '../routes/bridgeLimitsRoute'
import { USDC, ETH } from '../constants/bridgeable'
import { NativeGasAddress } from '../constants'

const app = express()
app.use('/bridgeLimits', bridgeLimitsRoute)

describe('Get Bridge Limits Route', () => {
  it('should return min/max origin amounts bridging USDC', async () => {
    const response = await request(app).get('/bridgeLimits').query({
      fromChain: 1,
      fromToken: USDC.addresses[1],
      toChain: 10,
      toToken: USDC.addresses[10],
    })

    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('maxOriginAmount')
    expect(response.body).toHaveProperty('minOriginAmount')
  }, 10_000)

  it('should return min/max origin amounts bridging ETH', async () => {
    const response = await request(app).get('/bridgeLimits').query({
      fromChain: 1,
      fromToken: ETH.addresses[1],
      toChain: 10,
      toToken: ETH.addresses[10],
    })

    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('maxOriginAmount')
    expect(response.body).toHaveProperty('minOriginAmount')
  }, 10_000)

  it('should return 400 for unsupported route', async () => {
    const response = await request(app).get('/bridgeLimits').query({
      fromChain: '1',
      toChain: '10',
      fromToken: NativeGasAddress,
      toToken: USDC.addresses[10],
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'No valid route exists for the chain/token combination'
    )
  })

  it('should return 400 for unsupported fromChain', async () => {
    const response = await request(app).get('/bridgeLimits').query({
      fromChain: '999',
      toChain: '137',
      fromToken: '0x176211869cA2b568f2A7D4EE941E073a821EE1ff',
      toToken: USDC.addresses[137],
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Unsupported fromChain'
    )
  })

  it('should return 400 for unsupported toChain', async () => {
    const response = await request(app).get('/bridgeLimits').query({
      fromChain: '137',
      toChain: '999',
      fromToken: USDC.addresses[137],
      toToken: '0x176211869cA2b568f2A7D4EE941E073a821EE1ff',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('message', 'Unsupported toChain')
  })

  it('should return 400 for missing fromToken', async () => {
    const response = await request(app).get('/bridgeLimits').query({
      fromChain: '1',
      toChain: '137',
      toToken: USDC.addresses[137],
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'fromToken')
  })

  it('should return 400 for missing toToken', async () => {
    const response = await request(app).get('/bridgeLimits').query({
      fromChain: '1',
      toChain: '137',
      fromToken: USDC.addresses[1],
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'toToken')
  })
})
