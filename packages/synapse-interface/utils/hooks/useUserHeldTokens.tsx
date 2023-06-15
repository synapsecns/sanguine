import { useState, useEffect, useMemo } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { AddressZero } from '@ethersproject/constants'
import { multicall, Address } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import multicallABI from '@/constants/abis/multicall.json'
import erc20ABI from '@/constants/abis/erc20.json'
import { Token } from '../types'
import { Contract } from 'ethers'

export function useUserHeldTokens() {
  const [tokens, setTokens] = useState([])
  const { address } = useAccount()
  const { chain } = useNetwork()

  return useMemo(async () => {
    if (address === undefined || chain === undefined) return []

    const currentChainBridgableTokens: Token[] = BRIDGABLE_TOKENS[chain.id]
    let multicallInputs = []
    let multicallData: any

    currentChainBridgableTokens.map((token) => {
      const tokenAddress = token.addresses[chain.id as keyof Token['addresses']]
      const multicallAddress: Address = `0xcA11bde05977b3631167028862bE2a173976CA11` //deterministic multicall3 ethereum address

      if (tokenAddress === undefined) return
      else if (tokenAddress === AddressZero) {
        multicallInputs.push({
          address: multicallAddress,
          abi: multicallABI,
          functionName: 'getEthBalance',
        } as Partial<Contract>)
      } else {
        const formattedTokenAddress: Address = `0x${tokenAddress.slice(2)}`
        multicallInputs.push({
          address: formattedTokenAddress,
          abi: erc20ABI,
          functionName: 'balanceOf',
          chainId: chain.id as number,
          args: [address],
        } as Partial<Contract>)
      }
    })

    if (multicallInputs.length > 0) {
      multicallData = await multicall({ contracts: multicallInputs })
    }
  }, [address, chain])
}
