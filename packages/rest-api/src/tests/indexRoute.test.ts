import request from 'supertest'
import express from 'express'

import indexRoute from '../routes/indexRoute'

const app = express()
app.use('/', indexRoute)

describe('Index Route', () => {
  it('should return welcome message', async () => {
    const response = await request(app).get('/')

    expect(response.status).toBe(200)

    expect(response.body).toHaveProperty('message')
    expect(response.body.message).toBe(
      'Welcome to the Synapse REST API for swap and bridge quotes'
    )
  })

  it('should return available chains and available tokens', async () => {
    const response = await request(app).get('/')

    expect(response.status).toBe(200)

    expect(response.body).toHaveProperty('availableChains')
    expect(response.body.availableChains.length).toBe(26)

    expect(response.body).toHaveProperty('availableTokens')
    expect(response.body.availableTokens.length).toBe(68)
  })
})
