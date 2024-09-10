import { useEffect, useMemo, useState } from 'react'
import { useAccount, useAccountEffect, useSwitchChain } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { useTranslations } from 'next-intl'

import {
  usePoolDataState,
  usePoolDepositState,
  usePoolUserDataState,
} from '@/slices/pools/hooks'
import { stringToBigInt } from '@/utils/bigint/format'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { DEFAULT_DEPOSIT_QUOTE } from './Deposit'
import LoadingDots from '@/components/ui/tailwind/LoadingDots'

const DepositButton = ({ approveTxn, depositTxn }) => {
  const t = useTranslations('Pools.DepositButton')
  const [isConnected, setIsConnected] = useState(false)
  const { openConnectModal } = useConnectModal()

  const { chain, isConnected: isConnectedInit } = useAccount()
  const { chains, switchChain } = useSwitchChain()

  useAccountEffect({
    onDisconnect() {
      setIsConnected(false)
    },
  })

  useEffect(() => {
    setIsConnected(isConnectedInit)
  }, [isConnectedInit])

  const { pool, poolData } = usePoolDataState()
  const { depositQuote, inputValue, isLoading, inputSum } =
    usePoolDepositState()
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
      label: t('insufficientBalance'),
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
      label: t('connectWallet'),
      onClick: openConnectModal,
    }
  } else if (chain?.id !== pool.chainId) {
    const targetChain = chains.find((c) => c.id === pool.chainId)
    buttonProperties = {
      label: t('switchChain', { chainName: targetChain.name }),
      onClick: () => switchChain({ chainId: pool.chainId }),
      pendingLabel: t('switchingChains'),
    }
  } else if (isApprovalNeeded) {
    buttonProperties = {
      onClick: approveTxn,
      label: t('approveTokens'),
      pendingLabel: t('approving'),
    }
  } else {
    buttonProperties = {
      onClick: depositTxn,
      label: t('deposit'),
      pendingLabel: t('depositing'),
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
