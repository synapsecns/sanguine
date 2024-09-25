import { type Abi, type Address } from 'viem'

export interface AddressConfig {
  name: string
  FastBridgeV2: {
    address: Address
    abi: Abi
    startBlock: number
  }
}