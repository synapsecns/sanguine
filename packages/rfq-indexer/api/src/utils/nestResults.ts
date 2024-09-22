export const nest_results = (sqlResults: any[]) => {
  return sqlResults.map((transaction: any) => {
    const bridgeRequest: { [key: string]: any } = {}
    const bridgeRelay: { [key: string]: any } = {}
    const bridgeProof: { [key: string]: any } = {}
    const bridgeClaim: { [key: string]: any } = {}
    const bridgeRefund: { [key: string]: any } = {}
    const bridgeDispute: { [key: string]: any } = {}
    const transactionFields: { [key: string]: any } = {}

    let transactionIdSet = false

    for (const [key, value] of Object.entries(transaction)) {
      if (key.startsWith('transactionId')) {
        if (!transactionIdSet) {
          transactionFields[key.replace(/_.+$/, '')] = value
          transactionIdSet = true
        }
        // Ignore other transactionId fields
      } else if (key.endsWith('_deposit')) {
        bridgeRequest[key.replace('_deposit', '')] = value
      } else if (key.endsWith('_relay')) {
        bridgeRelay[key.replace('_relay', '')] = value
      } else if (key.endsWith('_proof')) {
        bridgeProof[key.replace('_proof', '')] = value
      } else if (key.endsWith('_claim')) {
        bridgeClaim[key.replace('_claim', '')] = value
      } else if (key.endsWith('_refund')) {
        bridgeRefund[key.replace('_refund', '')] = value
      } else if (key.endsWith('_dispute')) {
        bridgeDispute[key.replace('_dispute', '')] = value
      } else {
        transactionFields[key] = value
      }
    }

    const result: { [key: string]: any } = { Bridge: transactionFields }
    if (Object.keys(bridgeRequest).length) {
      result.BridgeRequest = bridgeRequest
    }
    if (Object.keys(bridgeRelay).length) {
      result.BridgeRelay = bridgeRelay
    }
    if (Object.keys(bridgeProof).length) {
      result.BridgeProof = bridgeProof
    }
    if (Object.keys(bridgeClaim).length) {
      result.BridgeClaim = bridgeClaim
    }
    if (Object.keys(bridgeRefund).length) {
      result.BridgeRefund = bridgeRefund
    }
    if (Object.keys(bridgeDispute).length) {
      result.BridgeDispute = bridgeDispute
    }
    return result
  })
}
