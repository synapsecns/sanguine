import React, { useMemo, useCallback, useState } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { useAccount } from 'wagmi'
import {
  setFromChainId,
  setFromToken,
  updateFromValue,
} from '@/slices/bridge/reducer'
import { Token } from '@/utils/types'
import { formatBigIntToString } from '@/utils/bigint/format'
import { CHAINS_BY_ID } from '@/constants/chains'
import { inputRef } from '../../../StateManagedBridge/InputContainer'
import { approveToken } from '@/utils/approveToken'
import { Address, switchNetwork } from '@wagmi/core'
import Image from 'next/image'
import { toast } from 'react-hot-toast'
import {
  ROUTER_ADDRESS,
  CCTP_ROUTER_ADDRESS,
  Allowances,
} from '@/utils/actions/fetchPortfolioBalances'
import { useBridgeState } from '@/slices/bridge/hooks'
import { fetchAndStoreSingleTokenAllowance } from '@/slices/portfolio/hooks'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { isUndefined } from 'lodash'

const handleFocusOnInput = () => {
  inputRef.current.focus()
}

type PortfolioTokenAssetProps = {
  token: Token
  balance: bigint
  allowances?: Allowances
  portfolioChainId: number
  connectedChainId: number
}

export const PortfolioTokenAsset = ({
  token,
  balance,
  allowances,
  portfolioChainId,
  connectedChainId,
}: PortfolioTokenAssetProps) => {
  const dispatch = useAppDispatch()
  const { fromChainId, fromToken, toChainId, toToken, bridgeQuote } =
    useBridgeState()
  const { address } = useAccount()
  const { icon, symbol, decimals, addresses } = token as Token

  const parsedBalance: string = useMemo(() => {
    const formattedBalance = formatBigIntToString(
      balance,
      decimals[portfolioChainId],
      3
    )
    return balance > 0n && hasOnlyZeroes(formattedBalance)
      ? '< 0.001'
      : formattedBalance
  }, [balance, portfolioChainId])

  const parsedBalanceLong: string = useMemo(() => {
    const formattedBalance = formatBigIntToString(
      balance,
      decimals[portfolioChainId],
      8
    )
    return balance > 0n && hasOnlyZeroes(formattedBalance)
      ? '< 0.001'
      : formattedBalance
  }, [balance, portfolioChainId])

  const isCCTP: boolean = bridgeQuote?.bridgeModuleName === 'SynapseCCTP'

  const tokenRouterAddress: string = bridgeQuote?.routerAddress

  const bridgeAllowance: bigint = allowances?.[tokenRouterAddress]

  const parsedAllowance: string =
    bridgeAllowance &&
    formatBigIntToString(bridgeAllowance, decimals[portfolioChainId], 3)

  const currentChainName: string = CHAINS_BY_ID[portfolioChainId].name

  const tokenAddress: string = addresses[portfolioChainId]

  const isCurrentlyConnected: boolean = portfolioChainId === connectedChainId

  const isTokenSelected: boolean = useMemo(() => {
    return fromToken === token && fromChainId === portfolioChainId
  }, [fromChainId, fromToken, token, portfolioChainId])

  const hasAllowanceButLessThanBalance: boolean =
    bridgeAllowance && balance > bridgeAllowance

  const isDisabled: boolean = false

  // /** Fetch allowances for selected token via current router if not already stored */
  if (isTokenSelected && tokenRouterAddress && isUndefined(bridgeAllowance)) {
    ;(async () => {
      await dispatch(
        fetchAndStoreSingleTokenAllowance({
          routerAddress: tokenRouterAddress as Address,
          tokenAddress: tokenAddress as Address,
          address: address as Address,
          chainId: portfolioChainId as number,
        })
      )
    })()
  }

  const handleTotalBalanceInputCallback = useCallback(async () => {
    await dispatch(setFromChainId(portfolioChainId as number))
    await dispatch(setFromToken(token as Token))
    await dispatch(
      await updateFromValue(
        formatBigIntToString(
          balance,
          token.decimals[portfolioChainId]
        ) as string
      )
    )
    handleFocusOnInput()
  }, [isDisabled, token, balance, portfolioChainId])

  const handleSelectFromTokenCallback = useCallback(() => {
    dispatch(setFromChainId(portfolioChainId as number))
    dispatch(setFromToken(token as Token))
    handleFocusOnInput()
  }, [token, isDisabled, portfolioChainId])

  const handleApproveCallback = useCallback(async () => {
    if (isCurrentlyConnected) {
      dispatch(setFromChainId(portfolioChainId as number))
      dispatch(setFromToken(token as Token))
      try {
        await approveToken(
          tokenRouterAddress,
          connectedChainId,
          tokenAddress
        ).then((success) => {
          dispatch(
            fetchAndStoreSingleTokenAllowance({
              routerAddress: tokenRouterAddress as Address,
              tokenAddress: tokenAddress as Address,
              address: address as Address,
              chainId: portfolioChainId as number,
            })
          )
        })
      } catch (error) {
        toast.error(
          `Failed to approve ${token.symbol} token on ${currentChainName} network`,
          {
            id: 'approve-in-progress-popup',
            duration: 5000,
          }
        )
      }
    } else {
      try {
        await switchNetwork({ chainId: portfolioChainId })
        await approveToken(
          tokenRouterAddress,
          portfolioChainId,
          tokenAddress
        ).then((success) => {
          success &&
            dispatch(
              fetchAndStoreSingleTokenAllowance({
                routerAddress: tokenRouterAddress as Address,
                tokenAddress: tokenAddress as Address,
                address: address as Address,
                chainId: portfolioChainId as number,
              })
            )
        })
      } catch (error) {
        toast.error(
          `Failed to approve ${token.symbol} token on ${currentChainName} network`,
          {
            id: 'approve-in-progress-popup',
            duration: 5000,
          }
        )
      }
    }
  }, [
    token,
    address,
    tokenAddress,
    connectedChainId,
    portfolioChainId,
    isCurrentlyConnected,
    isDisabled,
    tokenRouterAddress,
  ])

  return (
    <div
      data-test-id="portfolio-token-asset"
      className={`
        p-2 flex items-center border-y text-white justify-between last:rounded-b-md
        ${
          isTokenSelected
            ? 'bg-tint border-surface'
            : 'border-transparent'
        }
      `}
    >
      <div
        onClick={handleTotalBalanceInputCallback}
        className={`
          flex items-center gap-2
          pl-2 pr-3 py-2 cursor-pointer rounded
          hover:bg-tint active:opacity-70
        `}
        title={`${parsedBalanceLong} ${symbol}`}
      >
        <Image
          loading="lazy"
          alt={`${symbol} img`}
          className="w-6 h-6 rounded-md"
          src={icon}
        />
        {parsedBalance} {symbol}
      </div>
      <PortfolioAssetActionButton
        selectCallback={handleSelectFromTokenCallback}
        isDisabled={isDisabled || isTokenSelected}
        isSelected={isTokenSelected}
      />
    </div>
  )
}

export const HoverClickableText = ({
  defaultText,
  hoverText,
  callback,
}: {
  defaultText: string
  hoverText: string
  callback: () => void
}) => {
  const [isHovered, setIsHovered] = useState<boolean>(false)
  return (
    <div
      data-test-id="hover-clickable-text"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
      onClick={callback}
      className={`
        group px-2
        text-[#A3A3C2]
        hover:text-[#75E6F0]
        hover:underline
        hover:cursor-pointer
        active:opacity-70
      `}
    >
      <div className="text-sm">{isHovered ? hoverText : defaultText}</div>
    </div>
  )
}

type PortfolioAssetActionButtonProps = {
  selectCallback: () => void
  isDisabled: boolean
  isSelected: boolean
}

const PortfolioAssetActionButton = ({
  selectCallback,
  isDisabled,
  isSelected,
}: PortfolioAssetActionButtonProps) => {
  return (
    <React.Fragment>
      <button
        data-test-id="portfolio-asset-action-button"
        className={`
          py-1 px-6 rounded-sm
          border border-fuchsia-500
          ${!isDisabled && 'cursor-pointer hover:bg-zinc-800 active:opacity-70'}
        `}
        onClick={selectCallback}
        disabled={isDisabled}
      >
        Select{isSelected && 'ed'}
      </button>
    </React.Fragment>
  )
}
