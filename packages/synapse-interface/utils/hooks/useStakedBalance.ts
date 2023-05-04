import { Interface } from '@ethersproject/abi'
import { useState, useEffect } from 'react'
import { readContracts, Address } from '@wagmi/core'
import { useBlockNumber } from 'wagmi'
import MINICHEF_ABI from '@abis/miniChef.json'

import { useMiniChefContract } from '../contracts/useMiniChefContract'

export const useStakedBalance = ({ poolId }: { poolId: number }) => {
  const [balance, setBalance] = useState<any>()
  const blockNumber = useBlockNumber()
  const miniChefContract = useMiniChefContract()
  const miniChefAddress:Address = miniChefContract && miniChefContract.address

  useEffect(() => {
    if (!miniChefContract || !miniChefAddress || !blockNumber) return

    (async () => {
      try{
        setBalance(await readContracts({
          contracts: [
            {
              address: miniChefAddress,
              abi: MINICHEF_ABI
              functionName: 'pendingSynapse'
            }
          ]
        }))
      } catch (error) {
        console.error('Error from useStakedBalance: ', error)
      }
    })()
  }, [])

  return poolId
}
