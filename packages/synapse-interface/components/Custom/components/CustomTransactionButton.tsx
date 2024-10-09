import { useMemo, useState } from 'react'
import { useAccount, useSwitchChain } from 'wagmi'

import { CHAINS_BY_ID } from '@/constants/chains'
import { stringToBigInt } from '@/utils/bigint/format'
import { useConnectModal } from '@rainbow-me/rainbowkit'

export const CustomTransactionButton = ({
  fromChainId,
  toChainId,
  fromToken,
  bridgeQuote,
  fromValue,
  fromTokenBalance,
  isLoading,
  approveTxn,
  executeBridge,
}) => {
  const { chain, isConnected } = useAccount()
  const { openConnectModal } = useConnectModal()
  const { switchChain } = useSwitchChain()

  const [isApproving, setIsApproving] = useState(false)
  const [isBridging, setIsBridging] = useState(false)

  const buttonClassName = `
    p-2 mb-2
    text-lg font-sans font-medium tracking-wide w-full
    shadow-[0_0_0_2px_#00C185,0_0_0_4px_#FF8736,0_0_0_6px_#FFC100] 
  `

  const comparableFromTokenBalance = stringToBigInt(
    fromTokenBalance,
    fromToken.decimals[fromChainId]
  )
  const comparableFromValue = stringToBigInt(
    fromValue,
    fromToken.decimals[fromChainId]
  )

  const isApproved = useMemo(() => {
    return (
      fromToken &&
      bridgeQuote?.allowance &&
      stringToBigInt(fromValue, fromToken.decimals[fromChainId]) <=
        bridgeQuote.allowance
    )
  }, [bridgeQuote, fromToken, fromValue, fromChainId])

  const handleApproveTxn = async () => {
    setIsApproving(true)
    try {
      await approveTxn()
    } catch (error) {
      console.error('Approval failed', error)
    } finally {
      setIsApproving(false)
    }
  }

  const handleBridgeTxn = async () => {
    setIsBridging(true)
    try {
      await executeBridge()
    } catch (error) {
      console.error('Bridge failed', error)
    } finally {
      setIsBridging(false)
    }
  }

  if (isLoading) {
    return <button className={buttonClassName}>Loading quote...</button>
  }

  if (!isConnected) {
    return (
      <button className={buttonClassName} onClick={openConnectModal}>
        Connect Wallet
      </button>
    )
  }

  if (isConnected && chain.id !== fromChainId) {
    return (
      <button
        className={buttonClassName}
        onClick={() => switchChain({ chainId: fromChainId })}
      >
        Switch to {CHAINS_BY_ID[fromChainId].name}
      </button>
    )
  }

  if (fromValue === '' || fromValue === '0') {
    return <button className={buttonClassName}>Enter an amount</button>
  }

  if (comparableFromValue > comparableFromTokenBalance) {
    return <button className={buttonClassName}>Insufficient balance</button>
  }

  if (!isApproved) {
    return (
      <button
        className={buttonClassName}
        onClick={handleApproveTxn}
        disabled={isApproving}
      >
        {isApproving ? 'Approving...' : 'Approve'}
      </button>
    )
  }

  return (
    <button
      className={buttonClassName}
      onClick={handleBridgeTxn}
      disabled={isBridging}
    >
      {isBridging ? 'Bridging...' : 'Bridge'}
    </button>
  )
}
