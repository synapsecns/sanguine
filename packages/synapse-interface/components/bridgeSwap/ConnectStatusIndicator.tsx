import { useEffect, useMemo, useState } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import {
  ConnectToNetworkButton,
  ConnectWalletButton,
  ConnectedIndicator,
} from '@/components/ConnectionIndicators'



export const ConnectStatusIndicator = ({ targetChainId }) => {
  const { isConnected } = useAccount()
  const { chain } = useNetwork()
  const [hasMounted, setHasMounted] = useState(false)

  useEffect(() => {
    setHasMounted(true)
  }, [])

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
