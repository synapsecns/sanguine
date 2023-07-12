import { useState } from 'react'

import LiquidityManagementTabs from '../components/LiquidityManagementTabs'
import Deposit from './Deposit'
import Withdraw from './Withdraw'
import { useSelector } from 'react-redux'
import { RootState } from '@/store/store'
import LoadingSpinner from '@/components/ui/tailwind/LoadingSpinner'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { TransactionButton } from '@/components/buttons/TransactionButton'

const PoolManagement = ({
  address,
  chainId,
}: {
  address: string
  chainId: number
}) => {
  const [cardNav, setCardNav] = useState(getLiquidityMode('#addLiquidity')) // 'addLiquidity'
  const { isConnected } = useAccount()
  const { chain } = useNetwork()

  const { pool } = useSelector((state: RootState) => state.poolData)
  const { poolUserData, isLoading } = useSelector(
    (state: RootState) => state.poolUserData
  )
  const { openConnectModal } = useConnectModal()
  const { chains, error, pendingChainId, switchNetwork } = useSwitchNetwork()

  if (isConnected && chain.id !== pool.chainId) {
    return (
      <div className="flex flex-col justify-center h-full">
        <TransactionButton
          label={`Switch to ${chains.find((c) => c.id === pool.chainId).name}`}
          pendingLabel="Switching chains"
          onClick={() =>
            new Promise((resolve, reject) => {
              try {
                switchNetwork(pool.chainId)
                resolve(true)
              } catch (e) {
                reject(e)
              }
            })
          }
        />
      </div>
    )
  }

  if (!isConnected) {
    return (
      <div className="flex flex-col justify-center h-full">
        <TransactionButton
          label="Connect wallet"
          pendingLabel="Connecting"
          onClick={() =>
            new Promise((resolve, reject) => {
              try {
                openConnectModal()
                resolve(true)
              } catch (e) {
                reject(e)
              }
            })
          }
        />
      </div>
    )
  }

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <LoadingSpinner />
      </div>
    )
  }

  return (
    <div>
      <div className="rounded-lg text-default">
        <LiquidityManagementTabs
          cardNav={cardNav}
          setCardNav={(val) => {
            setCardNav(val)
          }}
        />
        <div className="mt-4">
          {cardNav === 'addLiquidity' && poolUserData.tokens && (
            <Deposit address={address} chainId={chainId} />
          )}
          {cardNav === 'removeLiquidity' && (
            <Withdraw chainId={chainId} address={address} />
          )}
        </div>
      </div>
    </div>
  )
}

function getLiquidityMode(browserHash) {
  switch (browserHash) {
    case '#addLiquidity':
      return 'addLiquidity'
    case '#removeLiquidity':
      return 'removeLiquidity'
    default:
      return 'addLiquidity'
  }
}
export default PoolManagement
