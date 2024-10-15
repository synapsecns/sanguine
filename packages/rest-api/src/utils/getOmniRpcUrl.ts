import * as dotenv from 'dotenv'

dotenv.config()

export const getOmniRpcUrl = (chainId: number) => {
  if (!process.env.OMNIRPC_BASE_URL) {
    throw new Error('OMNIRPC_BASE_URL environment variable is not set')
  }
  return `${process.env.OMNIRPC_BASE_URL}/${chainId}`
}
