import request from 'supertest'
import express from 'express'
import getSynapseTxIdRoute from '../routes/getSynapseTxIdRoute'

const app = express()
app.use('/getSynapseTxId', getSynapseTxIdRoute)

describe('Get Synapse TX ID Route', () => {
  it('should return synapse transaction ID for valid input', async () => {
    const response = await request(app).get('/getSynapseTxId').query({
      originChainId: '8453',
      bridgeModule: 'SynapseRFQ',
      txHash:
        '0x13486d9eaefd68de6a20b704d70deb8436effbac1f77fddfc0c7ef14f08e96c3',
    })
    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('synapseTxId')
  }, 10000)

  it('should return 400 for missing originChainId', async () => {
    const response = await request(app).get('/getSynapseTxId').query({
      bridgeModule: 'SynapseRFQ',
      txHash:
        '0x13486d9eaefd68de6a20b704d70deb8436effbac1f77fddfc0c7ef14f08e96c3',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'originChainId')
  }, 10000)

  it('should return 400 for missing bridgeModule', async () => {
    const response = await request(app).get('/getSynapseTxId').query({
      originChainId: '1',
      txHash:
        '0x13486d9eaefd68de6a20b704d70deb8436effbac1f77fddfc0c7ef14f08e96c3',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'bridgeModule')
  }, 10000)

  it('should return 400 for missing txHash', async () => {
    const response = await request(app).get('/getSynapseTxId').query({
      originChainId: '1',
      bridgeModule: 'SynapseRFQ',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'txHash')
  }, 10000)

  it('should return 400 for non-numeric originChainId', async () => {
    const response = await request(app).get('/getSynapseTxId').query({
      originChainId: 'not-a-number',
      bridgeModule: 'SynapseRFQ',
      txHash:
        '0x13486d9eaefd68de6a20b704d70deb8436effbac1f77fddfc0c7ef14f08e96c3',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'originChainId')
  }, 10000)

  it('should return 400 for invalid bridgeModule', async () => {
    const response = await request(app).get('/getSynapseTxId').query({
      originChainId: '1',
      bridgeModule: 'invalid_module',
      txHash:
        '0x13486d9eaefd68de6a20b704d70deb8436effbac1f77fddfc0c7ef14f08e96c3',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Invalid bridge module. Must be one of: SynapseBridge, SynapseCCTP, SynapseRFQ'
    )
  }, 10000)
})
