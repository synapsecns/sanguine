import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { Address, multicall, erc20ABI } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import { Token } from '../types'
import { AddressZero } from '@ethersproject/constants'
import multicallABI from '@/constants/abis/multicall.json'
import { getSortedBridgableTokens } from '../actions/getSortedBridgableTokens'

//move to constants file later
const MULTICALL3_ADDRESS: Address = '0xcA11bde05977b3631167028862bE2a173976CA11'

export const usePortfolioBalances = () => {
  const availableChains = Object.keys(BRIDGABLE_TOKENS)

  const foo = getSortedBridgableTokens(1)
  // const bar = getSortedBridgableTokens(42161)

  console.log('foo: ', foo)
  // console.log('bar: ', bar)
}

const useTokenApprovals = () => {}
