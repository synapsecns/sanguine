import { validationResult } from 'express-validator'

import { Synapse } from '../services/synapseService'

export const synapseTxIdController = async (req, res) => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() })
  }

  try {
    const { originChainId, bridgeModule, txHash } = req.query

    const synapseTxId = await Synapse.getSynapseTxId(
      Number(originChainId),
      bridgeModule,
      txHash
    )

    res.json({ synapseTxId })
  } catch (err) {
    res.status(500).json({
      error:
        'An unexpected error occurred in /synapseTxId. Please try again later.',
      details: err.message,
    })
  }
}
