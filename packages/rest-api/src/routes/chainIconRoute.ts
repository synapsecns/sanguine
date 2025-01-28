import express from 'express'
import { CHAINS, Chain } from '@synapsecns/synapse-constants'
import fetch from 'cross-fetch'

import { addSvgHeaderIfMissing } from '../utils/svgUtils'

const router: express.Router = express.Router()

router.get('/:chainId.svg', async (req, res) => {
  const chainId = parseInt(req.params.chainId, 10)

  // Find the chain with matching ID
  const chain = Object.values(CHAINS).find(
    (c): c is Chain =>
      typeof c === 'object' && c !== null && 'id' in c && c.id === chainId
  )

  if (!chain || !chain.chainImg) {
    res.status(404).json({ error: 'Chain icon not found' })
    return
  }

  try {
    // Fetch the image from the URL
    const response = await fetch(chain.chainImg)
    if (!response.ok) {
      throw new Error(`Failed to fetch image: ${response.statusText}`)
    }

    const buffer = await response.arrayBuffer()
    const contentType = response.headers.get('content-type') || 'image/svg+xml'

    // Only process SVG files
    const processedBuffer =
      contentType === 'image/svg+xml'
        ? addSvgHeaderIfMissing(buffer)
        : Buffer.from(buffer)

    // Set cache headers (cache for 1 week)
    res.set({
      'Cache-Control': 'public, max-age=604800',
      'Content-Type': contentType,
    })

    res.send(processedBuffer)
  } catch (error) {
    console.error('Error fetching chain icon:', error)
    res.status(500).json({ error: 'Failed to fetch chain icon' })
  }
  return
})

export default router
