import request from 'supertest'

import { server } from './app'

describe('Server Endpoints', () => {
  describe('/swap endpoint', () => {
    it('should respond with Invalid Params for invalid request', async () => {
      const res = await request(server).get('/swap')
      expect(res.statusCode).toEqual(200)
      expect(res.text).toContain('<h1>Invalid Params</h1>')
    })
  })
  describe('/swap passing params', () => {
    it('should respond with swap quote for valid request', async () => {
      const res = await request(server).get(
        '/swap?chain=1&fromToken=USDC&toToken=DAI&amount=100'
      )
      expect(res.statusCode).toEqual(200)
      expect(res.maxAmountOutStr.length).toBeGreaterThan(0)
      expect(res.routerAddress.length).toBeGreaterThan(0)
      expect(res.query.length).toBeGreaterThan(0)
    })
  })

  describe('/bridge endpoint', () => {
    it('should respond with Invalid Request for invalid request', async () => {
      const res = await request(server).get('/bridge')
      expect(res.statusCode).toEqual(200)
      expect(res.text).toContain('<h1>Invalid Request</h1>')
    })
  })

  describe('/bridge passing params', () => {
    it('should respond with bridge quote for valid request', async () => {
      const res = await request(server).get(
        '/bridge?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000'
      )
      expect(res.statusCode).toEqual(200)
      expect(res.maxAmountOutStr.length).toBeGreaterThan(0)
      expect(res.routerAddress.length).toBeGreaterThan(0)
      expect(res.originQuery.length).toBeGreaterThan(0)
      expect(res.destQuery.length).toBeGreaterThan(0)
    })
  })
  afterAll((done) => {
    server.close(done)
  })
})
