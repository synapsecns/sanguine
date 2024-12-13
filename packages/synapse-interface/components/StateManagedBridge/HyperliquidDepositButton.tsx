import { useEffect, useState } from 'react'
import { useAccount, useAccountEffect, useSwitchChain } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { useTranslations } from 'next-intl'
import { Address, erc20Abi } from 'viem'
import {
  simulateContract,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'

import { wagmiConfig } from '@/wagmiConfig'
import { useAppDispatch } from '@/store/hooks'
import { useWalletState } from '@/slices/wallet/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { useBridgeValidations } from './hooks/useBridgeValidations'
import { USDC } from '@/constants/tokens/bridgeable'
import { ARBITRUM, HYPERLIQUID } from '@/constants/chains/master'
import { stringToBigInt } from '@/utils/bigint/format'
import { fetchAndStoreSingleNetworkPortfolioBalances } from '@/slices/portfolio/hooks'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { addPendingBridgeTransaction } from '@/slices/transactions/actions'
import { getUnixTimeMinutesFromNow } from '@/utils/time'

const HYPERLIQUID_DEPOSIT_ADDRESS = '0x2Df1c51E09aECF9cacB7bc98cB1742757f163dF7'

const approve = async (address: Address, amount: bigint) => {
  const { request } = await simulateContract(wagmiConfig, {
    chainId: ARBITRUM.id,
    address: USDC.addresses[ARBITRUM.id],
    abi: erc20Abi,
    functionName: 'approve',
    args: [address, amount],
  })

  const hash = await writeContract(wagmiConfig, request)

  const txReceipt = await waitForTransactionReceipt(wagmiConfig, { hash })

  return txReceipt
}

const deposit = async (amount: bigint) => {
  try {
    const { request } = await simulateContract(wagmiConfig, {
      chainId: ARBITRUM.id,
      address: USDC.addresses[ARBITRUM.id],
      abi: erc20Abi,
      functionName: 'transfer',
      args: [HYPERLIQUID_DEPOSIT_ADDRESS, amount],
    })

    const hash = await writeContract(wagmiConfig, request)

    const txReceipt = await waitForTransactionReceipt(wagmiConfig, { hash })

    return txReceipt
  } catch (error) {
    console.error('Confirmation error:', error)
    throw error
  }
}

export const HyperliquidTransactionButton = ({
  isTyping,
  hasDepositedOnHyperliquid,
  setHasDepositedOnHyperliquid,
}) => {
  const [isApproved, setIsApproved] = useState(false)
  const [isApproving, setIsApproving] = useState(false)
  const [isDepositing, setIsDepositing] = useState(false)

  const { address } = useAccount()

  const dispatch = useAppDispatch()
  const { openConnectModal } = useConnectModal()
  const [isConnected, setIsConnected] = useState(false)

  const { isConnected: isConnectedInit } = useAccount()
  const { chains, switchChain } = useSwitchChain()

  const { fromToken, fromChainId, debouncedFromValue } = useBridgeState()

  const { isWalletPending } = useWalletState()

  const { hasValidInput, hasSufficientBalance, onSelectedChain } =
    useBridgeValidations()

  const depositingMinimumAmount = Number(debouncedFromValue) >= 5

  const t = useTranslations('Bridge')

  const amount = stringToBigInt(
    debouncedFromValue,
    fromToken.decimals[fromChainId]
  )

  const handleApprove = async () => {
    setIsApproving(true)

    try {
      await approve(address, amount)
      setIsApproved(true)
    } catch (error) {
      console.error('Approval error:', error)
    } finally {
      setIsApproving(false)
    }
  }

  const handleDeposit = async () => {
    setIsDepositing(true)
    const currentTimestamp: number = getUnixTimeMinutesFromNow(0)
    try {
      const txReceipt = await deposit(amount)

      setHasDepositedOnHyperliquid(true)
      setIsApproved(false)
      segmentAnalyticsEvent(`[Hyperliquid Deposit]`, {
        inputAmount: debouncedFromValue,
      })
      dispatch(
        fetchAndStoreSingleNetworkPortfolioBalances({
          address,
          chainId: ARBITRUM.id,
        })
      )
      dispatch(
        addPendingBridgeTransaction({
          id: currentTimestamp,
          originChain: ARBITRUM,
          originToken: fromToken,
          originValue: debouncedFromValue,
          destinationChain: HYPERLIQUID,
          destinationToken: undefined,
          transactionHash: txReceipt.transactionHash,
          timestamp: undefined,
          isSubmitted: false,
          estimatedTime: undefined,
          bridgeModuleName: undefined,
          destinationAddress: undefined,
          routerAddress: undefined,
        })
      )
    } catch (error) {
      console.error('Deposit error:', error)
    } finally {
      setIsDepositing(false)
    }
  }

  useAccountEffect({
    onDisconnect() {
      setIsConnected(false)
    },
  })

  useEffect(() => {
    setIsConnected(isConnectedInit)
  }, [isConnectedInit])

  const isButtonDisabled =
    isTyping ||
    isApproving ||
    isDepositing ||
    !depositingMinimumAmount ||
    isWalletPending ||
    !hasValidInput ||
    (isConnected && !hasSufficientBalance)

  let buttonProperties

  if (isConnected && !hasSufficientBalance) {
    buttonProperties = {
      label: t('Insufficient balance'),
      onClick: null,
    }
  } else if (!depositingMinimumAmount) {
    buttonProperties = {
      label: '5 USDC Minimum',
      onClick: null,
    }
  } else if (!isConnected && hasValidInput) {
    buttonProperties = {
      label: t('Connect Wallet to Bridge'),
      onClick: openConnectModal,
    }
  } else if (!onSelectedChain && hasValidInput) {
    buttonProperties = {
      label: t('Switch to {chainName}', {
        chainName: chains.find((c) => c.id === fromChainId)?.name,
      }),
      onClick: () => switchChain({ chainId: fromChainId }),
      pendingLabel: t('Switching chains'),
    }
  } else if (!isApproved && hasValidInput) {
    buttonProperties = {
      onClick: handleApprove,
      label: t('Approve {symbol}', { symbol: fromToken?.symbol }),
      pendingLabel: t('Approving'),
    }
  } else {
    buttonProperties = {
      onClick: handleDeposit,
      label: t('Deposit {symbol}', { symbol: fromToken?.symbol }),
      pendingLabel: t('Depositing'),
    }
  }

  return (
    buttonProperties && (
      <>
        <div className="flex flex-col w-full">
          <TransactionButton
            {...buttonProperties}
            disabled={isButtonDisabled}
            chainId={fromChainId}
          />
        </div>
      </>
    )
  )
}
