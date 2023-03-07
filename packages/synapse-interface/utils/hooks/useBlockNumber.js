import { useCallback, useContext } from 'react'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { ChainContext } from '@store/ChainStore'

/**
 * @param {number} targetChainId
 */
export function useBlockNumber(targetChainId) {

  const [chainState, setChainState] = useContext(ChainContext)

  const setBlockNumber = useCallback(
    (updatedBlockNumber) => {
      if (chainState.blockNumber[targetChainId] < updatedBlockNumber) {
        setChainState({
          ...chainState,
          blockNumber: {
            ...chainState.blockNumber,
            [targetChainId]: updatedBlockNumber
          }
        })
      }
    }, [
      targetChainId
    ]
  )
  return [chainState.blockNumber[targetChainId], setBlockNumber]
}
