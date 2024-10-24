export const validateKappa = (synapseTxId: string) => {
  let hexRegex

  if (synapseTxId.startsWith('0x')) {
    hexRegex = /^0x[0-9a-fA-F]{64}$/
  } else {
    hexRegex = /^[0-9a-fA-F]{64}$/
  }

  return hexRegex.test(synapseTxId)
}
