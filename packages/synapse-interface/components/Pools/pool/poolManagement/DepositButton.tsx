
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import { useEffect, useMemo, useState } from 'react'
import {
  usePoolDataState,
  usePoolUserDataState,
  usePoolDepositState,
} from '@/slices/pool/hooks'
import { fetchPoolUserData } from '@/slices/poolUserDataSlice'
import LoadingDots from '@tw/LoadingDots'
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

  const { pool, poolData } = usePoolDataState()

  const { depositQuote, inputValue, isLoading, inputSum } = usePoolDepositState()
  const { poolUserData } = usePoolUserDataState()

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

  const isEmptyPool = useMemo(() => {
    return poolData.totalLocked === 0
  }, [poolData])

  const isButtonDisabled =
    (isLoading ||
      !isBalanceEnough ||
      depositQuote === DEFAULT_DEPOSIT_QUOTE ||
      inputSum === 0 ||
      inputSum === 0n) &&
    !isEmptyPool

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

export default DepositButton
