import { useEffect } from 'react'
import { useAccount } from 'wagmi'
import { useApplicationState } from './hooks'
import { ApplicationState } from './reducer'
import { updateLastConnectedAddress, updateLastConnectedTime } from './actions'
import { useAppDispatch } from '@/store/hooks'
import { isValidAddress, getValidAddress } from '@/utils/isValidAddress'
import { getTimeMinutesBeforeNow } from '@/utils/time'
import { resetReduxCache } from './actions'

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const { address } = useAccount()

  const { lastConnectedAddress, lastConnectedTimestamp }: ApplicationState =
    useApplicationState()

  /**
   * Record last connected address and timestamp into persisted state
   * Clear redux cache after 7 days
   * Clear redux cache if new address connects
   */
  useEffect(() => {
    if (isValidAddress(address)) {
      if (
        isValidAddress(lastConnectedAddress) &&
        getValidAddress(address) !== getValidAddress(lastConnectedAddress)
      ) {
        console.log('reset redux cache')
        dispatch(resetReduxCache())
        dispatch(updateLastConnectedAddress(address))
      }

      const currentTime: number = getTimeMinutesBeforeNow(0)
      const sevenDaysAgo: number = getTimeMinutesBeforeNow(10080)

      if (lastConnectedTimestamp < sevenDaysAgo) {
        console.log('reset timestamp cache')
        dispatch(resetReduxCache())
      }
      dispatch(updateLastConnectedTime(getTimeMinutesBeforeNow(0)))
    }
  }, [address, lastConnectedAddress, lastConnectedTimestamp])

  return null
}
