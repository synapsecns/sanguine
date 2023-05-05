import { useState, useEffect } from 'react'
import { readContracts, ReadContractResult, Address } from '@wagmi/core'
import { useBlockNumber, useAccount } from 'wagmi'
import { Contract, BigNumber } from 'ethers'
import MINICHEF_ABI from '@abis/miniChef.json'
import { Zero } from '@ethersproject/constants'

import { useMiniChefContract } from '../contracts/useMiniChefContract'

export const useStakedBalance = ({ poolId }: { poolId: number }) => {
  const [balance, setBalance] = useState<[BigNumber, BigNumber]>([Zero, Zero])
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
        const data: ReadContractResult = await readContracts({
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
        if (data[0]?.amount && data[0]?.reward) {
          setBalance([data[0].amount, data[0].reward])
        }
      } catch (error) {
        console.error('Error from useStakedBalance: ', error)
      }
    })()
  }, [address])

  return { amount: balance[0] ?? Zero, reward: balance[1] ?? Zero }
}
