import { useEffect, useState } from 'react'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { DEFAULT_WITHDRAW_QUOTE } from '@/slices/poolWithdrawSlice'
import {
  usePoolDataState,
  usePoolWithdrawState,
  usePoolUserDataState,
} from '@/slices/pools/hooks'
import { stringToBigInt } from '@/utils/bigint/format'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import LoadingDots from '@/components/ui/tailwind/LoadingDots'

const WithdrawButton = ({ approveTxn, withdrawTxn, isApproved }) => {
  const { chain } = useNetwork()
  const { chains, switchNetwork } = useSwitchNetwork()
  const { openConnectModal } = useConnectModal()
  const [isConnected, setIsConnected] = useState(false) // Initialize to false

  const { isConnected: isConnectedInit } = useAccount({
    onDisconnect() {
      setIsConnected(false)
    },
  })

  useEffect(() => {
    setIsConnected(isConnectedInit)
  }, [isConnectedInit])

  const { pool } = usePoolDataState()
  const { poolUserData } = usePoolUserDataState()
  const { withdrawQuote, inputValue, isLoading } = usePoolWithdrawState()

  const poolDecimals = pool?.decimals[pool?.chainId]

  const isBalanceEnough =
    stringToBigInt(inputValue, poolDecimals) !== 0n &&
    stringToBigInt(inputValue, poolDecimals) <= poolUserData.lpTokenBalance

  const isValidInput = stringToBigInt(inputValue, poolDecimals) !== 0n
  const isValidQuote = withdrawQuote !== DEFAULT_WITHDRAW_QUOTE
  const isButtonDisabled = isLoading || !isBalanceEnough || !isValidQuote

  let buttonProperties

  if (!isBalanceEnough && isValidQuote && isValidInput) {
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
  } else if (!isApproved && isValidQuote && isValidInput) {
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
        style={
          isButtonDisabled
            ? {
                border: '1px solid #453F47',
                borderRadius: '4px',
              }
            : {
                background:
                  'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
                border: '1px solid #9B6DD7',
                borderRadius: '4px',
              }
        }
        {...buttonProperties}
        disabled={isButtonDisabled}
        chainId={pool.chainId}
      />
    )
  )
}

export default WithdrawButton
