import { useAccount, useNetwork } from 'wagmi'
import {
  ConnectToNetworkButton,
  ConnectWalletButton,
  ConnectedIndicator,
} from '@/components/ConnectionIndicators'




export const ConnectStatusIndicator = ({ targetChainId }) => {
  const { isConnected } = useAccount()
  const { chain } = useNetwork()

  if (isConnected) {
    if (targetChainId === chain?.id) {
      return <ConnectedIndicator />
    } else {
      return <ConnectToNetworkButton chainId={targetChainId} />
    }
  } else {
    return <ConnectWalletButton />
  }
}
