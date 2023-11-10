import { useEffect } from 'react'
import { useAccount } from 'wagmi'
import { useApplicationState } from './hooks'
import { ApplicationState } from './reducer'
import { updateLastConnectedAddress, updateLastConnectedTime } from './actions'
import { useAppDispatch } from '@/store/hooks'
import { isValidAddress, getValidAddress } from '@/utils/isValidAddress'
import { getTimeMinutesBeforeNow } from '@/utils/time'
import { resetReduxCache } from './actions'
import { resetTransactionsState } from '../transactions/actions'

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const { address } = useAccount()

  const { lastConnectedAddress, lastConnectedTimestamp }: ApplicationState =
    useApplicationState()

  /**
   * Record last connected address and timestamp into persisted state
   * Clear redux cache after 7 days
   */
  useEffect(() => {
    if (lastConnectedTimestamp) {
      const sevenDaysInSeconds = 7 * 24 * 60 * 60
      const sevenDaysAgo: number = getTimeMinutesBeforeNow(10080)

      if (sevenDaysAgo - lastConnectedTimestamp > sevenDaysInSeconds) {
        console.log('reset cache from < 7 days stale')
        dispatch(resetReduxCache())
      }

      if (
        isValidAddress(lastConnectedAddress) &&
        getValidAddress(address) !== getValidAddress(lastConnectedAddress)
      ) {
        dispatch(resetTransactionsState())
      }

      dispatch(updateLastConnectedAddress(address))
      dispatch(updateLastConnectedTime(getTimeMinutesBeforeNow(0)))
    } else {
      dispatch(updateLastConnectedTime(getTimeMinutesBeforeNow(0)))
    }
  }, [address, lastConnectedAddress, lastConnectedTimestamp])

  return null
}
