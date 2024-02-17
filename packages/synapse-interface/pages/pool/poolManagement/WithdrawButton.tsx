import { useSelector } from 'react-redux'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { RootState } from '@/store/store'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import { useEffect, useState } from 'react'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { stringToBigInt } from '@/utils/bigint/format'
import LoadingDots from '@/components/ui/tailwind/LoadingDots'
import { DEFAULT_WITHDRAW_QUOTE } from '@/slices/poolWithdrawSlice'

const WithdrawButton = ({ approveTxn, withdrawTxn, isApproved }) => {
  const [isConnected, setIsConnected] = useState(false) // Initialize to false
  const { openConnectModal } = useConnectModal()

  const { chain } = useNetwork()
  const { chains, switchNetwork } = useSwitchNetwork()

  const { isConnected: isConnectedInit } = useAccount({
    onDisconnect() {
      setIsConnected(false)
    },
  })

  useEffect(() => {
    setIsConnected(isConnectedInit)
  }, [isConnectedInit])

  const { pool } = useSelector((state: RootState) => state.poolData)

  const poolDecimals = pool?.decimals[pool?.chainId]

  const { withdrawQuote, inputValue, isLoading } = useSelector(
    (state: RootState) => state.poolWithdraw
  )
  const { poolUserData } = useSelector((state: RootState) => state.poolUserData)

  const needsInput = stringToBigInt(inputValue, poolDecimals) === 0n

  const isBalanceEnough =
    stringToBigInt(inputValue, poolDecimals) !== 0n &&
    stringToBigInt(inputValue, poolDecimals) <= poolUserData.lpTokenBalance

  const isButtonDisabled =
    isLoading || !isBalanceEnough || withdrawQuote === DEFAULT_WITHDRAW_QUOTE

  let buttonProperties

  if (needsInput) {
    buttonProperties = {
      label: 'Enter amount',
      onClick: null,
    }
  } else if (!isBalanceEnough) {
    buttonProperties = {
      label: 'Insufficient Balance',
      onClick: null,
    }
  } else if (isLoading) {
    buttonProperties = {
      label: (
        <div className="flex items-center justify-center h-[24px]">
          <LoadingDots />
        </div>
      ),
      onClick: null,
    }
  } else if (!isConnected) {
    buttonProperties = {
      label: `Connect Wallet to Bridge`,
      onClick: openConnectModal,
    }
  } else if (chain?.id !== pool.chainId) {
    buttonProperties = {
      label: `Switch to ${chains.find((c) => c.id === pool.chainId).name}`,
      onClick: () => switchNetwork(pool.chainId),
      pendingLabel: 'Switching chains',
    }
  } else if (!isApproved) {
    buttonProperties = {
      onClick: approveTxn,
      label: `Approve Token`,
      pendingLabel: 'Approving',
    }
  } else {
    buttonProperties = {
      onClick: withdrawTxn,
      label: `Withdraw`,
      pendingLabel: 'Withdrawing...',
    }
  }

  return (
    pool &&
    buttonProperties && (
      <TransactionButton
        {...buttonProperties}
        disabled={isButtonDisabled}
        chainId={pool.chainId}
      />
    )
  )
}

export default WithdrawButton
