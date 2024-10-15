import * as dotenv from 'dotenv'

dotenv.config()

export const getOmniRpcUrl = (chainId: number) => {
  if (!process.env.OMNIRPC_BASE_URL) {
    return null
  }
  return `${process.env.OMNIRPC_BASE_URL}/${chainId}`
}
