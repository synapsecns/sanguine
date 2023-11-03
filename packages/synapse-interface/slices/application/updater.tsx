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

  useEffect(() => {
    if (isValidAddress(address)) {
      if (
        isValidAddress(lastConnectedAddress) &&
        getValidAddress(address) !== getValidAddress(lastConnectedAddress)
      ) {
        dispatch(resetReduxCache())
        dispatch(updateLastConnectedAddress(address))
      }
      dispatch(updateLastConnectedTime(getTimeMinutesBeforeNow(0)))
    }
  }, [address, lastConnectedAddress, lastConnectedTimestamp])

  return null
}
