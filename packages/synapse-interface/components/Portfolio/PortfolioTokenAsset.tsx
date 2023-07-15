import React, { useMemo, useCallback } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '@/store/store'
import { useAccount } from 'wagmi'
import { BigNumber } from 'ethers'
import {
  setFromChainId,
  setFromToken,
  updateFromValue,
} from '@/slices/bridgeSlice'
import { Token } from '@/utils/types'
import { formatBNToString } from '@/utils/bignumber/format'
import { CHAINS_BY_ID } from '@/constants/chains'
import { inputRef } from '../StateManagedBridge/InputContainer'
import { approveToken } from '@/utils/approveToken'
import { switchNetwork } from '@wagmi/core'
import Image from 'next/image'
import { toast } from 'react-hot-toast'
import { ROUTER_ADDRESS } from '@/utils/hooks/usePortfolioBalances'

type PortfolioTokenAssetProps = {
  token: Token
  balance: BigNumber
  allowance?: BigNumber
  portfolioChainId: number
  connectedChainId: number
  isApproved: boolean
  fetchPortfolioBalancesCallback: () => Promise<void>
}

function hasOnlyZeros(input: string): boolean {
  return /^0+(\.0+)?$/.test(input)
}

function scrollToTop() {
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const handleFocusOnInput = () => {
  inputRef.current.focus()
}

export const PortfolioTokenAsset = ({
  token,
  balance,
  allowance,
  portfolioChainId,
  connectedChainId,
  isApproved,
  fetchPortfolioBalancesCallback,
}: PortfolioTokenAssetProps) => {
  const dispatch = useDispatch()
  const { fromChainId, fromToken } = useSelector(
    (state: RootState) => state.bridge
  )
  const { address } = useAccount()
  const { icon, symbol, decimals, addresses } = token

  const parsedBalance: string = useMemo(() => {
    const formattedBalance = formatBNToString(
      balance,
      decimals[portfolioChainId],
      3
    )
    return balance.gt(0) && hasOnlyZeros(formattedBalance)
      ? '< 0.001'
      : formattedBalance
  }, [balance, portfolioChainId])

  const parsedAllowance: string =
    allowance && formatBNToString(allowance, decimals[portfolioChainId], 3)

  const currentChainName: string = CHAINS_BY_ID[portfolioChainId].name

  const tokenAddress: string = addresses[portfolioChainId]

  const isCurrentlyConnected: boolean = portfolioChainId === connectedChainId

  const isTokenSelected: boolean = useMemo(() => {
    return fromToken === token && fromChainId === portfolioChainId
  }, [fromChainId, fromToken, token, portfolioChainId])

  const hasAllowanceButLessThanBalance: boolean =
    allowance && balance.gt(allowance)

  const isDisabled: boolean = false

  const handleTotalBalanceInputCallback = useCallback(() => {
    dispatch(setFromToken(token))
    dispatch(setFromChainId(portfolioChainId))
    dispatch(updateFromValue(balance))
  }, [isDisabled, token, balance])

  const handleSelectFromTokenCallback = useCallback(() => {
    dispatch(setFromChainId(portfolioChainId))
    dispatch(setFromToken(token))
    scrollToTop()
    handleFocusOnInput()
  }, [token, isDisabled, portfolioChainId])

  const handleApproveCallback = useCallback(async () => {
    if (isCurrentlyConnected) {
      dispatch(setFromChainId(portfolioChainId))
      dispatch(setFromToken(token))
      await approveToken(ROUTER_ADDRESS, connectedChainId, tokenAddress).then(
        (success) => {
          success && fetchPortfolioBalancesCallback()
        }
      )
    } else {
      try {
        await switchNetwork({ chainId: portfolioChainId })
        await scrollToTop()
        await approveToken(ROUTER_ADDRESS, portfolioChainId, tokenAddress).then(
          (success) => {
            success && fetchPortfolioBalancesCallback()
          }
        )
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
  ])

  return (
    <div
      data-test-id="portfolio-token-asset"
      className="flex flex-row items-center py-2 text-white"
    >
      {isTokenSelected ? (
        <div className="w-4 m-auto font-bold text-green-500"> ✓ </div>
      ) : (
        <div className="w-4" />
      )}
      <div className="flex flex-row justify-between w-2/3">
        <div
          onClick={handleSelectFromTokenCallback}
          className={`
          flex flex-row px-2 py-2 mb-auto
          hover:cursor-pointer
          hover:bg-[#272731]
        `}
        >
          <Image
            loading="lazy"
            alt={`${symbol} img`}
            className="w-6 h-6 mr-2 rounded-md"
            src={icon}
          />
          <div>{symbol}</div>
        </div>
        <div className="flex flex-col">
          <div
            onClick={handleTotalBalanceInputCallback}
            className={`
            p-2 ml-auto cursor-pointer
            hover:bg-[#272731] active:opacity-[67%]
          `}
          >
            {parsedBalance}
          </div>
          {hasAllowanceButLessThanBalance && (
            <div
              onClick={handleApproveCallback}
              className={`
              text-[#A3A3C2] text-xs pt-1 px-2
              hover:text-[#75E6F0]
              hover:underline
              hover:cursor-pointer
              active:opacity-[67%]
            `}
            >
              {parsedAllowance} approved
            </div>
          )}
        </div>
      </div>
      <div className="flex flex-row items-center w-1/3 py-2 mb-auto text-left">
        <PortfolioAssetActionButton
          selectCallback={handleSelectFromTokenCallback}
          approveCallback={handleApproveCallback}
          isApproved={isApproved}
          isDisabled={isDisabled}
        />
      </div>
    </div>
  )
}

type PortfolioAssetActionButtonProps = {
  selectCallback: () => void
  approveCallback: () => Promise<void>
  isApproved: boolean
  isDisabled: boolean
}

const PortfolioAssetActionButton = ({
  selectCallback,
  approveCallback,
  isApproved,
  isDisabled,
}: PortfolioAssetActionButtonProps) => {
  const buttonClassName = `
    flex ml-auto justify-center
    py-1 px-6 ml-2 rounded-3xl
    transform-gpu transition-all duration-75
    ${isDisabled ? 'hover:cursor-default' : 'hover:cursor-pointer'}
  `
  return (
    <React.Fragment>
      {isApproved ? (
        <button
          data-test-id="portfolio-asset-action-button"
          className={`
            ${buttonClassName}
            border border-[#D747FF] mr-1
            hover:bg-[#272731]
            active:opacity-[67%]
          `}
          onClick={selectCallback}
        >
          Select
        </button>
      ) : (
        <button
          data-test-id="portfolio-asset-action-button"
          className={`
            ${buttonClassName}
            border border-[#3D3D5C]
            hover:border-[#A3A3C2]
            hover:bg-[#272731]
            active:border-[#A3A3C2]
            active:opacity-[67%]
          `}
          onClick={approveCallback}
        >
          Approve
        </button>
      )}
    </React.Fragment>
  )
}
