import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { Address, multicall, erc20ABI } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import { Token } from '../types'
import { AddressZero } from '@ethersproject/constants'
import multicallABI from '@/constants/abis/multicall.json'
import { getSortedBridgableTokens } from '../actions/getSortedBridgableTokens'
import { ChainId } from '@/constants/chains'

//move to constants file later
const MULTICALL3_ADDRESS: Address = '0xcA11bde05977b3631167028862bE2a173976CA11'

export const usePortfolioBalances = () => {
  const availableChains = Object.keys(BRIDGABLE_TOKENS)

  const test1 = getSortedBridgableTokens(ChainId.ARBITRUM)
  const test2 = getSortedBridgableTokens(ChainId.ETH)
  const test3 = getSortedBridgableTokens(ChainId.AVALANCHE)

  console.log('test1: ', test1)
  console.log('test2: ', test2)
  console.log('test3: ', test3)
}

const useTokenApprovals = () => {}
