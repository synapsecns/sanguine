import express from 'express'
import { BRIDGABLE_TOKENS, Token } from '@synapsecns/synapse-constants'
import fetch from 'cross-fetch'

import { addSvgHeaderIfMissing } from '../utils/svgUtils'

const router: express.Router = express.Router()

router.get('/:chainId/:address.svg', async (req, res) => {
  const chainId = parseInt(req.params.chainId, 10)
  const address = req.params.address.toLowerCase()
  const addHeaders = req.query.headers === 'true'

  // Find the token with matching address on the specified chain
  const token = Object.values(BRIDGABLE_TOKENS[chainId] || []).find(
    (t): t is Token =>
      typeof t === 'object' &&
      t !== null &&
      'addresses' in t &&
      Object.entries(t.addresses).some(([chain, addr]) => {
        const matches =
          parseInt(chain, 10) === chainId && addr.toLowerCase() === address
        return matches
      })
  )

  if (!token || !token.icon) {
    console.log('Token not found or no icon:', { token })
    res.status(404).json({ error: 'Token icon not found' })
    return
  }

  try {
    // Fetch the image from the URL
    const response = await fetch(token.icon)
    if (!response.ok) {
      throw new Error(`Failed to fetch image: ${response.statusText}`)
    }

    const buffer = await response.arrayBuffer()
    const contentType = response.headers.get('content-type') || 'image/svg+xml'

    // Only process SVG files if headers are requested
    const processedBuffer = contentType === 'image/svg+xml' && addHeaders
      ? addSvgHeaderIfMissing(buffer)
      : Buffer.from(buffer)

    // Set cache headers (cache for 1 week)
    res.set({
      'Cache-Control': 'public, max-age=604800',
      'Content-Type': contentType,
    })

    res.send(processedBuffer)
  } catch (error) {
    console.error('Error fetching token icon:', error)
    res.status(500).json({ error: 'Failed to fetch token icon' })
  }
  return
})

export default router
