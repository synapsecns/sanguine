import { useSelector } from 'react-redux'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import { useEffect, useState } from 'react'
import { RootState } from '@/store/store'

import LoadingSpinner from '@/components/ui/tailwind/LoadingSpinner'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { DEFAULT_DEPOSIT_QUOTE } from './Deposit'
import { stringToBigInt } from '@/utils/bigint/format'

const DepositButton = ({ approveTxn, depositTxn }) => {
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

  const { depositQuote, inputValue, isLoading, inputSum } = useSelector(
    (state: RootState) => state.poolDeposit
  )
  const { poolUserData } = useSelector((state: RootState) => state.poolUserData)

  const isBalanceEnough = Object.entries(inputValue.bi).every(
    ([tokenAddr, amount]) =>
      poolUserData.tokens.every((tokenObj) => {
        if (tokenObj.token.addresses[pool.chainId] !== tokenAddr) return true
        const rawBalanceBigInt = stringToBigInt(
          `${tokenObj.balance}`,
          tokenObj.token.decimals[pool.chainId]
        )

        return amount <= rawBalanceBigInt
      })
  )

  const isApprovalNeeded = Object.entries(inputValue.bi).some(
    ([tokenAddr, amount]) => {
      return (
        typeof amount !== 'undefined' &&
        Object.keys(depositQuote.allowances).length > 0 &&
        amount !== 0n &&
        typeof depositQuote.allowances[tokenAddr] !== 'undefined' &&
        amount > BigInt(depositQuote.allowances[tokenAddr])
      )
    }
  )

  const isButtonDisabled =
    isLoading ||
    !isBalanceEnough ||
    depositQuote === DEFAULT_DEPOSIT_QUOTE ||
    inputSum === 0 ||
    inputSum === 0n

  let buttonProperties

  if (!isBalanceEnough) {
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
  } else if (isApprovalNeeded) {
    buttonProperties = {
      onClick: approveTxn,
      label: `Approve Token(s)`,
      pendingLabel: 'Approving',
    }
  } else {
    buttonProperties = {
      onClick: depositTxn,
      label: `Deposit`,
      pendingLabel: 'Depositing',
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

export default DepositButton
