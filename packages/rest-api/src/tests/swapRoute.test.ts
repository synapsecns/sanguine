import request from 'supertest'
import express from 'express'

import swapRoute from '../routes/swapRoute'
import { NativeGasAddress, ZeroAddress } from '../constants'
import { DAI, ETH, NETH, USDC } from '../constants/bridgeable'

const app = express()
app.use('/swap', swapRoute)

describe('Swap Route with Real Synapse Service', () => {
  it('should return a real swap quote for valid input, 1000 USDC', async () => {
    const response = await request(app).get('/swap').query({
      chain: '1',
      fromToken: USDC.addresses[1],
      toToken: DAI.addresses[1],
      amount: '1000',
    })

    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('maxAmountOut')
    expect(response.body).toHaveProperty('routerAddress')
    expect(response.body).toHaveProperty('query')
  }, 10_000)

  it('should return a real swap quote for valid input, Eth ZeroAddress', async () => {
    const response = await request(app).get('/swap').query({
      chain: '10',
      fromToken: ZeroAddress,
      toToken: NETH.addresses[10],
      amount: '1',
    })

    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('maxAmountOut')
    expect(response.body).toHaveProperty('routerAddress')
    expect(response.body).toHaveProperty('query')
  }, 10_000)

  it('should return a real swap quote for valid input, Eth NativeGasAddress', async () => {
    const response = await request(app).get('/swap').query({
      chain: '10',
      fromToken: NativeGasAddress,
      toToken: NETH.addresses[10],
      amount: '1',
    })

    expect(response.status).toBe(200)
    expect(response.body).toHaveProperty('maxAmountOut')
    expect(response.body).toHaveProperty('routerAddress')
    expect(response.body).toHaveProperty('query')
  }, 10_000)

  it('should return 400 for unsupported chain, with error message', async () => {
    const response = await request(app).get('/swap').query({
      chain: '111',
      fromToken: USDC.addresses[1],
      toToken: DAI.addresses[1],
      amount: '1000',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('message', 'Unsupported chain')
  }, 10_000)

  it('should return 400 for invalid toToken address, with error message', async () => {
    const response = await request(app).get('/swap').query({
      chain: '1',
      fromToken: USDC.addresses[1],
      toToken: 'invalid_address',
      amount: '1000',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Invalid toToken address'
    )
  }, 10_000)

  it('should return 400 for invalid fromToken + toToken combo', async () => {
    const response = await request(app).get('/swap').query({
      chain: '1',
      fromToken: ETH.addresses[1],
      toToken: USDC.addresses[1],
      amount: '1000',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Swap route not supported'
    )
  })

  it('should return 400 for token not supported on specified chain', async () => {
    const response = await request(app).get('/swap').query({
      chain: '1',
      fromToken: USDC.addresses[1],
      toToken: '0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F', // SNX on Ethereum (Not supported)
      amount: '1000',
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty(
      'message',
      'Invalid toToken address'
    )
  }, 10_000)

  it('should return 400 for missing amount, with error message', async () => {
    const response = await request(app).get('/swap').query({
      chain: '1',
      fromToken: USDC.addresses[1],
      toToken: DAI.addresses[1],
    })

    expect(response.status).toBe(400)
    expect(response.body.error).toHaveProperty('field', 'amount')
  }, 10_000)
})
