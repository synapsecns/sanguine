import { useState, useEffect } from 'react'
import { readContracts, Address } from '@wagmi/core'
import { useBlockNumber, useAccount } from 'wagmi'
import { Contract } from 'ethers'
import MINICHEF_ABI from '@abis/miniChef.json'
import { Zero } from '@ethersproject/constants'

import { useMiniChefContract } from '../contracts/useMiniChefContract'

export const useStakedBalance = ({ poolId }: { poolId: number }) => {
  console.log('poolId: ', poolId)
  const [balance, setBalance] = useState<[any, any]>([undefined, undefined])
  const blockNumber = useBlockNumber()
  const { address } = useAccount()

  const [miniChefContract, miniChefAddress]: [Contract, Address] =
    useMiniChefContract()

  useEffect(() => {
    if (!miniChefContract || !miniChefAddress || !blockNumber || !address) {
      return
    }

    ;(async () => {
      try {
        setBalance(
          await readContracts({
            contracts: [
              {
                address: miniChefAddress,
                abi: MINICHEF_ABI,
                functionName: 'userInfo',
                args: [poolId, address],
              },
              {
                address: miniChefAddress,
                abi: MINICHEF_ABI,
                functionName: 'pendingSynapse',
                args: [poolId, address],
              },
            ],
          })
        )
      } catch (error) {
        console.error('Error from useStakedBalance: ', error)
      }
    })()
  }, [address])

  return { amount: balance[0] ?? Zero, reward: balance[1] ?? Zero }
}
