import createPersistedState from 'use-persisted-state'

import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'


const usePassthroughDestinationInfo = createPersistedState('destinationInfo42069')


export function useDestinationInfo() {
  const { account } = useActiveWeb3React()
  const [destinationInfo, setDestinationInfo] = usePassthroughDestinationInfo({})

  const addressesForAccount = destinationInfo[account] ?? []

  function setAddressesForAccount(destAddr) {
    setDestinationInfo({
      [account]: [destAddr, ...addressesForAccount],
      ...destinationInfo,
    })
  }

  return [addressesForAccount, setAddressesForAccount]
}