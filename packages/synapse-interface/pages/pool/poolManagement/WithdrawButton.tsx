import { useDispatch, useSelector } from 'react-redux'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { RootState } from '@/store/store'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import { useEffect, useState } from 'react'
import {
  useConnectModal,
  useAccountModal,
  useChainModal,
} from '@rainbow-me/rainbowkit'
import { stringToBigInt } from '@/utils/stringToBigNum'
import LoadingSpinner from '@/components/ui/tailwind/LoadingSpinner'
import {
  DEFAULT_WITHDRAW_QUOTE,
  resetPoolWithdraw,
} from '@/slices/poolWithdrawSlice'
import { fetchPoolUserData } from '@/slices/poolUserDataSlice'

const WithdrawButton = ({ approveTxn, withdrawTxn, isApproved }) => {
  const dispatch: any = useDispatch()
  const [isConnected, setIsConnected] = useState(false) // Initialize to false
  const { openConnectModal } = useConnectModal()

  const { chain } = useNetwork()
  const { chains, error, pendingChainId, switchNetwork } = useSwitchNetwork()

  const { address, isConnected: isConnectedInit } = useAccount({
    onDisconnect() {
      setIsConnected(false)
    },
  })

  useEffect(() => {
    setIsConnected(isConnectedInit)
  }, [isConnectedInit])

  const { pool } = useSelector((state: RootState) => state.poolData)

  const { withdrawQuote, inputValue, isLoading } = useSelector(
    (state: RootState) => state.poolWithdraw
  )
  const { poolUserData } = useSelector((state: RootState) => state.poolUserData)

  const needsInput = inputValue.bi === 0n

  const isBalanceEnough =
    inputValue.bi !== 0n && inputValue.bi <= poolUserData.lpTokenBalance

  const isButtonDisabled =
    isLoading || !isBalanceEnough || withdrawQuote === DEFAULT_WITHDRAW_QUOTE

  let buttonProperties

  if (!pool) {
    return
  }

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
          <LoadingSpinner />
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
