import { useContext } from 'react'
import { NetworkContext } from '@store/NetworkStore'

export function useNetworkController() {
  return useContext(NetworkContext)
}