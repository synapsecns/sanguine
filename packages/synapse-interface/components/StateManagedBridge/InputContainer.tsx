import React, { useEffect, useState, useRef, useCallback, useMemo } from 'react'
import { useDispatch } from 'react-redux'
import { Address, useAccount, useNetwork } from 'wagmi'

import { initialState, updateFromValue } from '@/slices/bridge/reducer'
import MiniMaxButton from '../buttons/MiniMaxButton'
import { formatBigIntToString } from '@/utils/bigint/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import {
  ConnectToNetworkButton,
  ConnectWalletButton,
  ConnectedIndicator,
} from '@/components/ConnectionIndicators'
import { FromChainSelector } from './FromChainSelector'
import { FromTokenSelector } from './FromTokenSelector'
import { useBridgeState, useBridgeStatus } from '@/slices/bridge/hooks'
import {
  fetchAndStoreSingleTokenAllowance,
  usePortfolioState,
} from '@/slices/portfolio/hooks'
import { TokenWithBalanceAndAllowances } from '@/utils/actions/fetchPortfolioBalances'
import { approveToken } from '@/utils/approveToken'
import { useAppDispatch } from '@/store/hooks'
import { txErrorHandler } from '@/utils/txErrorHandler'
import { LoaderIcon } from 'react-hot-toast'
import { EMPTY_BRIDGE_QUOTE, EMPTY_BRIDGE_QUOTE_ZERO } from '@/constants/bridge'

export const inputRef = React.createRef<HTMLInputElement>()

export const InputContainer = () => {
  const dispatch = useDispatch()
  const { chain } = useNetwork()
  const { fromChainId, fromToken, fromValue } = useBridgeState()
  const { hasInputAmount, hasValidSelections, hasValidRoute, isConnected } =
    useBridgeStatus()

  const [showValue, setShowValue] = useState('')
  const [hasMounted, setHasMounted] = useState(false)

  useEffect(() => {
    setHasMounted(true)
  }, [])

  useEffect(() => {
    if (fromToken && fromToken?.decimals[fromChainId]) {
      setShowValue(fromValue)
    }

    if (fromValue === initialState.fromValue) {
      setShowValue(initialState.fromValue)
    }
  }, [fromValue, inputRef, fromChainId, fromToken])

  useEffect(() => {
    if (hasValidSelections && inputRef.current) {
      inputRef.current.focus()
      inputRef.current.select()
    }
  }, [hasValidSelections])

  const handleFromValueChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const fromValueString: string = cleanNumberInput(event.target.value)
    try {
      dispatch(updateFromValue(fromValueString))
      setShowValue(fromValueString)
    } catch (error) {
      console.error('Invalid value for conversion to BigInteger')
      const inputValue = event.target.value
      const regex = /^[0-9]*[.,]?[0-9]*$/

      if (regex.test(inputValue) || inputValue === '') {
        dispatch(updateFromValue(inputValue))
        setShowValue(inputValue)
      }
    }
  }

  const connectedStatus = useMemo(() => {
    if (hasMounted && !isConnected) {
      return (
        <ConnectWalletButton highlight={hasValidSelections && hasValidRoute} />
      )
    } else if (hasMounted && isConnected && fromChainId === chain.id) {
      return <ConnectedIndicator />
    } else if (
      hasMounted &&
      isConnected &&
      fromChainId &&
      fromChainId !== chain.id
    ) {
      return <ConnectToNetworkButton chainId={fromChainId} />
    }
  }, [
    chain,
    fromChainId,
    isConnected,
    hasMounted,
    hasValidSelections,
    hasValidRoute,
  ])

  return (
    <div
      data-test-id="input-container"
      className="text-left rounded-md p-md bg-bgLight"
    >
      <div className="flex items-center justify-between mb-3">
        <FromChainSelector />
        {connectedStatus}
      </div>
      <div className="flex h-16 mb-2 space-x-2">
        <div
          className={`
            flex flex-grow items-center justify-between
            pl-md w-full h-16 rounded-md
            border border-white border-opacity-20
          `}
        >
          <div className="flex items-center">
            <FromTokenSelector />
            <div className="flex flex-col justify-between ml-4">
              <div style={{ display: 'table' }}>
                <input
                  ref={inputRef}
                  pattern="^[0-9]*[.,]?[0-9]*$"
                  disabled={false}
                  className={`
                  text-white text-opacity-80 text-xl font-medium
                    p-0 border-none bg-transparent max-w-[190px]
                    md:text-2xl placeholder:text-[#88818C]
                    focus:outline-none focus:ring-0 focus:border-none
                  `}
                  placeholder="0.0000"
                  onChange={handleFromValueChange}
                  value={showValue}
                  name="inputRow"
                  autoComplete="off"
                  minLength={1}
                  maxLength={79}
                  style={{ display: 'table-cell', width: '100%' }}
                  onClick={() => inputRef.current && inputRef.current.select()}
                />
              </div>
              {hasMounted && <ShowLabel />}
            </div>
          </div>
          <div>
            {hasMounted && isConnected && (
              <div className="">
                <ShowButton />
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}

const ShowLabel = () => {
  const dispatch = useAppDispatch()
  const { fromChainId, fromToken, bridgeQuote, isLoading } = useBridgeState()
  const { hasEnoughBalance, hasInputAmount, hasEnoughApproved, isConnected } =
    useBridgeStatus()
  const { balancesAndAllowances } = usePortfolioState()

  const balance = balancesAndAllowances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )?.balance
  const parsedBalance: string = balancesAndAllowances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )?.parsedBalance

  const parsedAllowance = balancesAndAllowances[fromChainId]?.find(
    (token: TokenWithBalanceAndAllowances) => token.token === fromToken
  )?.allowances[bridgeQuote?.routerAddress]

  const onMaxBalance = useCallback(() => {
    dispatch(
      updateFromValue(
        formatBigIntToString(balance, fromToken?.decimals[fromChainId])
      )
    )
  }, [balance, fromChainId, fromToken, isLoading])

  const onMaxApproved = useCallback(() => {
    dispatch(
      updateFromValue(
        formatBigIntToString(parsedAllowance, fromToken?.decimals[fromChainId])
      )
    )
  }, [balance, fromChainId, fromToken, isLoading])

  if (!fromToken) {
    return (
      <label
        htmlFor="inputRow"
        onClick={null}
        className={`
          text-sm 
        text-secondary
        `}
      >
        <span>Select token</span>
      </label>
    )
  }

  if (!isConnected) {
    return null
  }

  if (
    bridgeQuote === EMPTY_BRIDGE_QUOTE_ZERO ||
    bridgeQuote === EMPTY_BRIDGE_QUOTE ||
    !hasInputAmount ||
    !hasEnoughBalance ||
    hasEnoughApproved
  ) {
    return (
      <label
        htmlFor="inputRow"
        onClick={onMaxBalance}
        className={`
          text-sm transition-all duration-150 transform-gpu
          hover:text-opacity-70 hover:cursor-pointer
          text-secondary
          ${
            fromToken && hasInputAmount && !hasEnoughBalance
              ? 'text-synapseYellow'
              : 'text-secondary'
          }
        `}
      >
        {parsedBalance ?? '0.0'}
        <span> available</span>
      </label>
    )
  }

  if (!hasEnoughApproved) {
    return (
      <label
        htmlFor="inputRow"
        onClick={onMaxApproved}
        className={`
          text-sm transition-all duration-150 transform-gpu
          hover:text-opacity-70 hover:cursor-pointer
          text-secondary
        `}
      >
        {formatBigIntToString(
          parsedAllowance,
          fromToken?.decimals[fromChainId],
          5
        ) ?? '0.0'}
        <span> approved</span>
      </label>
    )
  }
}

const ShowButton = ({}) => {
  const dispatch = useAppDispatch()
  const { fromChainId, fromToken, bridgeQuote } = useBridgeState()
  const {
    hasValidSelections,
    hasEnoughBalance,
    hasInputAmount,
    hasEnoughApproved,
    onSelectedChain,
  } = useBridgeStatus()
  const { balancesAndAllowances } = usePortfolioState()

  const balance = balancesAndAllowances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )?.balance

  const onMaxBalance = useCallback(() => {
    dispatch(
      updateFromValue(
        formatBigIntToString(balance, fromToken?.decimals[fromChainId])
      )
    )
  }, [balance, fromChainId, fromToken])

  if (
    bridgeQuote === EMPTY_BRIDGE_QUOTE_ZERO ||
    bridgeQuote === EMPTY_BRIDGE_QUOTE ||
    !hasInputAmount ||
    !hasEnoughBalance ||
    hasEnoughApproved
  ) {
    return (
      <MiniMaxButton
        disabled={!balance || balance === 0n ? true : false}
        onClickBalance={onMaxBalance}
      />
    )
  }

  if (!hasEnoughApproved) {
    return (
      <ApproveButton
        disabled={!hasInputAmount || !hasValidSelections || !onSelectedChain}
      />
    )
  }
}

const ApproveButton = ({ disabled }) => {
  const [isApproving, setIsApproving] = useState(false)
  const dispatch = useAppDispatch()
  const { address } = useAccount()
  const { fromChainId, fromToken, bridgeQuote } = useBridgeState()
  const { onSelectedChain, hasEnoughApproved } = useBridgeStatus()

  const approveTxn = async () => {
    setIsApproving(true)
    try {
      const tx = approveToken(
        bridgeQuote?.routerAddress,
        fromChainId,
        fromToken?.addresses[fromChainId]
      ).then(() => {
        dispatch(
          fetchAndStoreSingleTokenAllowance({
            routerAddress: bridgeQuote?.routerAddress as Address,
            tokenAddress: fromToken?.addresses[fromChainId] as Address,
            address: address,
            chainId: fromChainId,
          })
        )
      })

      try {
        await tx
      } catch (error) {
        return txErrorHandler(error)
      }
    } catch (error) {
      return txErrorHandler(error)
    } finally {
      setIsApproving(false)
    }
  }

  const borderStyle: {} = useMemo(() => {
    if (isApproving) {
      return {
        borderColor: '#D747FF',
        background:
          'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
      }
    } else if (onSelectedChain && !hasEnoughApproved) {
      return {
        borderColor: '#D747FF',
        background:
          'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
      }
    } else {
      return {}
    }
  }, [onSelectedChain, hasEnoughApproved, isApproving])

  return (
    <button
      style={borderStyle}
      className={`
        w-[89px] h-[32px]
        flex items-center mr-2 py-lg px-md justify-center
        text-sm text-white
        border rounded-sm
        ${disabled ? 'opacity-50' : 'hover:border-secondary'}
      `}
      disabled={disabled}
      onClick={approveTxn}
    >
      {isApproving ? <LoaderIcon /> : 'Approve'}
    </button>
  )
}
