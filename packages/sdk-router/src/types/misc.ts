export { FeeConfig, Query } from '../module'
export { PoolToken } from '../router'
export { Slippage } from '../swap'

/**
 * Populated transaction.
 *
 * @property {string} to - Address of the transaction recipient, in 0x format
 * @property {string} data - Data to be attached to the transaction, in 0x format
 * @property {string} value - Value to be sent with the transaction, in wei
 */
export type PopulatedTx = {
  to: string
  data: string
  value: string
}
