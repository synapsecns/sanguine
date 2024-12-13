import { useEffect, useState } from 'react'
import { useAccount, useAccountEffect, useSwitchChain } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { useTranslations } from 'next-intl'

import { useAppDispatch } from '@/store/hooks'
import { useWalletState } from '@/slices/wallet/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { useBridgeValidations } from './hooks/useBridgeValidations'
import { USDC } from '@/constants/tokens/bridgeable'
import { ARBITRUM, HYPERLIQUID } from '@/constants/chains/master'
import { Address, erc20Abi } from 'viem'
import { MAX_UINT256 } from '@/constants'
import { wagmiConfig } from '@/wagmiConfig'
import {
  simulateContract,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'
import { stringToBigInt } from '@/utils/bigint/format'
import { set } from 'lodash'
import { arbitrum } from 'viem/chains'
import { ArrowRightIcon, CheckCircleIcon } from '@heroicons/react/outline'
import { ArrowUpRightIcon } from '../icons/ArrowUpRightIcon'
import { shortenAddress } from '@/utils/shortenAddress'
import Image from 'next/image'
import { CopyButton } from '../ui/CopyButton'
import { fetchAndStoreSingleNetworkPortfolioBalances } from '@/slices/portfolio/hooks'

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

export const HyperliquidTransactionButton = ({ isTyping }) => {
  const [isApproved, setIsApproved] = useState(false)
  const [isApproving, setIsApproving] = useState(false)
  const [isDepositing, setIsDepositing] = useState(false)
  const [hasDeposited, setHasDeposited] = useState(false)
  const [depositHash, setDepositHash] = useState('')

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
    try {
      const txReceipt = await deposit(amount)

      setDepositHash(txReceipt.transactionHash)
      setHasDeposited(true)
      setIsApproved(false)
      dispatch(
        fetchAndStoreSingleNetworkPortfolioBalances({
          address,
          chainId: ARBITRUM.id,
        })
      )
    } catch (error) {
      console.error('Deposit error:', error)
    } finally {
      setIsDepositing(false)
    }
  }

  useEffect(() => {
    // refresh balances
  }, [])

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
      label: t('Connect Wallet to Deposit'),
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
          {hasDeposited && (
            <div className="flex flex-col p-2 mb-2 text-sm border rounded border-zinc-300 dark:border-separator">
              <div className="flex items-center mb-2 space-x-2 ">
                <Image
                  loading="lazy"
                  src={HYPERLIQUID.chainImg}
                  alt="Switch Network"
                  width="16"
                  height="16"
                  className="w-4 h-4 max-w-fit"
                />

                <div>Hyperliquid Deposit</div>
              </div>
              <div className="flex items-center justify-between space-x-4 text-sm ">
                <div className="">
                  <span className="text-white/65">Receipt: </span>{' '}
                  {shortenAddress(depositHash, 8)}{' '}
                  <span className="text-white/65">/ </span>
                  <CopyButton text={depositHash} />
                </div>
                <a
                  href={`${arbitrum.blockExplorers.default.url}/tx/${depositHash}`}
                  target="_blank"
                  rel="noreferrer"
                  className="text-cortex-yellow hover:underline"
                >
                  <div className="flex items-center space-x-1">
                    <Image
                      src={ARBITRUM.explorerImg}
                      alt="Arbiscan"
                      width="12"
                      height="12"
                      className="w-3 h-3"
                    />
                    <div>{arbitrum.blockExplorers.default.name} </div>
                    <ArrowUpRightIcon className="w-2.5 h-2.5" />
                  </div>
                </a>
              </div>
            </div>
          )}

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
