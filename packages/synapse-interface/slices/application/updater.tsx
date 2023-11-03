import { useAccount } from 'wagmi'
import { updateLastConnectedAddress, updateLastConnectedTime } from './actions'
import { useAppDispatch } from '@/store/hooks'

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const { address } = useAccount()
}
