import { useMemo } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import {
  ConnectToNetworkButton,
  ConnectWalletButton,
  ConnectedIndicator,
} from '@/components/ConnectionIndicators'
import { useHasMounted } from '@/utils/hooks/useHasMounted'



export const ConnectStatusIndicator = ({ targetChainId }) => {
  const { isConnected } = useAccount()
  const { chain } = useNetwork()
  const hasMounted = useHasMounted()


  const connectedStatus = useMemo(() => {
    if (hasMounted) {
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
  }, [chain, targetChainId, isConnected, hasMounted])

  return <>{connectedStatus}</>
}
