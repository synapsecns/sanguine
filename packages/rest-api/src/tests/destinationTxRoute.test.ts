import request from 'supertest'
import express from 'express'

import destinationTxRoute from '../routes/destinationTxRoute'

const app = express()
app.use('/destinationTx', destinationTxRoute)

describe('Get Destination TX Route', () => {
  it('should return destination transaction info for valid input', async () => {
    const response = await request(app).get('/destinationTx').query({
      originChainId: '8453',
      txHash:
        '0x13486d9eaefd68de6a20b704d70deb8436effbac1f77fddfc0c7ef14f08e96c3',
    })

    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('status')
    expect(response.body.status).toBe('completed')
    expect(response.body).toHaveProperty('toInfo')
    if (response.body.toInfo) {
      expect(response.body.toInfo).toHaveProperty('chainID')
      expect(response.body.toInfo).toHaveProperty('address')
    }
  }, 10000)

  it('should return a refunded response for refunded transaction', async () => {
    const response = await request(app).get('/destinationTx').query({
      originChainId: '8453',
      txHash:
        '0x019f84bbb9999e3d34f8c636ddb6b4852bfeeaed423fd70607047f393cbfd070',
    })

    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('status')
    expect(response.body.status).toBe('refunded')
    expect(response.body).toHaveProperty('fromInfo')
    if (response.body.fromInfo) {
      expect(response.body.fromInfo).toHaveProperty('chainID')
      expect(response.body.fromInfo).toHaveProperty('address')
    }
    expect(response.body).toHaveProperty('toInfo')
    expect(response.body.toInfo).toBeNull()
  }, 10000)

  it('should return 404 for non-existent txHash', async () => {
    const response = await request(app).get('/destinationTx').query({
      originChainId: '8453',
      txHash:
        '0x12411d1beafd68de6a20b704d70deb8436effbac1f77fddfc0c7ef14f08e96c3',
    })

    expect(response.status).toBe(404)
    expect(response.body.status).toBe('not found')
  }, 10000)

  it('should return 400 for missing originChainId', async () => {
    const response = await request(app).get('/destinationTx').query({
      txHash:
        '0x13486d9eaefd68de6a20b704d70deb8436effbac1f77fddfc0c7ef14f08e96c3',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'originChainId')
  }, 10000)

  it('should return 400 for missing txHash', async () => {
    const response = await request(app).get('/destinationTx').query({
      originChainId: '1',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'txHash')
  }, 10000)

  it('should return 400 for non-numeric originChainId', async () => {
    const response = await request(app).get('/destinationTx').query({
      originChainId: 'not-a-number',
      txHash:
        '0x13486d9eaefd68de6a20b704d70deb8436effbac1f77fddfc0c7ef14f08e96c3',
    })
    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'originChainId')
  }, 10000)
})
