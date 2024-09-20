import request from 'supertest'
import express from 'express'

import bridgeTxStatusRoute from '../routes/bridgeTxStatusRoute'

const app = express()
app.use('/bridgeTxStatus', bridgeTxStatusRoute)

describe('Get Bridge TX Status Route', () => {
  it('should return bridge transaction status for valid input', async () => {
    const response = await request(app).get('/bridgeTxStatus').query({
      destChainId: '42161',
      bridgeModule: 'SynapseRFQ',
      synapseTxId:
        '0x9beb59b36ff4570d6b823b075dcd4fa9acd82dc4a28bf93a456ab8c93990604a',
    })

    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('status')
    expect(response.body).toHaveProperty('toInfo')
    expect(response.body.status).toBe(true)
    if (response.body.toInfo) {
      expect(response.body.toInfo).toHaveProperty('chainID')
      expect(response.body.toInfo).toHaveProperty('address')
      expect(response.body.toInfo).toHaveProperty('txnHash')
      expect(response.body.toInfo).toHaveProperty('formattedValue')
    }
  }, 10000)

  it('should return 400 for unsupported destChainId', async () => {
    const response = await request(app).get('/bridgeTxStatus').query({
      destChainId: '999',
      bridgeModule: 'bridge',
      synapseTxId:
        '0x9beb59b36ff4570d6b823b075dcd4fa9acd82dc4a28bf93a456ab8c93990604a',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Unsupported destChainId'
    )
  }, 10000)

  it('should return 400 for invalid bridgeModule', async () => {
    const response = await request(app).get('/bridgeTxStatus').query({
      destChainId: '1',
      bridgeModule: 'invalidModule',
      synapseTxId:
        '0x9beb59b36ff4570d6b823b075dcd4fa9acd82dc4a28bf93a456ab8c93990604a',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Invalid bridge module. Must be one of: SynapseBridge, SynapseCCTP, SynapseRFQ'
    )
  }, 10000)

  it('should return 400 for missing synapseTxId', async () => {
    const response = await request(app).get('/bridgeTxStatus').query({
      destChainId: '1',
      bridgeModule: 'SynapseRFQ',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'synapseTxId')
  }, 10000)

  it('should return 400 for missing destChainId', async () => {
    const response = await request(app).get('/bridgeTxStatus').query({
      bridgeModule: 'bridge',
      synapseTxId:
        '0x9beb59b36ff4570d6b823b075dcd4fa9acd82dc4a28bf93a456ab8c93990604a',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'destChainId')
  }, 10000)

  it('should return 400 for missing bridgeModule', async () => {
    const response = await request(app).get('/bridgeTxStatus').query({
      destChainId: '137',
      synapseTxId:
        '0x9beb59b36ff4570d6b823b075dcd4fa9acd82dc4a28bf93a456ab8c93990604a',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'bridgeModule')
  }, 10000)
})
