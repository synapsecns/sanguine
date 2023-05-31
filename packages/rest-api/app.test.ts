import request from 'supertest'

import { server } from './app'

describe('Server Endpoints', () => {
  describe('/swap endpoint', () => {
    it('should respond with Invalid Params for invalid request', async () => {
      const res = await request(server).get('/swap')
      expect(res.statusCode).toEqual(200)
      expect(res.text).toContain('<h1>Invalid Params</h1>')
    })

    // Add more tests to check valid swap
    // it('should respond with swap quote for valid request', async () => {
    //   const res = await request(app).get('/swap?...');
    //   expect(res.statusCode).toEqual(200);
    //   // expect...
    // });
  })

  describe('/bridge endpoint', () => {
    it('should respond with Invalid Request for invalid request', async () => {
      const res = await request(server).get('/bridge')
      expect(res.statusCode).toEqual(200)
      expect(res.text).toContain('<h1>Invalid Request</h1>')
    })

    // Add more tests to check valid bridge
    // it('should respond with bridge quote for valid request', async () => {
    //   const res = await request(app).get('/bridge?...');
    //   expect(res.statusCode).toEqual(200);
    //   // expect...
    // });
  })
  afterAll((done) => {
    server.close(done)
  })
})
