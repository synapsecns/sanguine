import { ChainId } from '@constants/networks'
import { createContext, useState } from 'react'


const initialState = {
  blockNumber: {
    [ChainId.ETH]:       1,
    [ChainId.BSC]:       1,
    [ChainId.POLYGON]:   1,
    [ChainId.FANTOM]:    1,
    [ChainId.BOBA]:      1,
    [ChainId.MOONRIVER]: 1,
    [ChainId.MOONBEAM]:  1,
    [ChainId.ARBITRUM]:  1,
    [ChainId.OPTIMISM]:  1,
    [ChainId.AVALANCHE]: 1,
    [ChainId.DFK]:       1,
    [ChainId.AURORA]:    1,
    [ChainId.HARMONY]:   1,
    [ChainId.CRONOS]:    1,
    [ChainId.METIS]:     1,
    [ChainId.KLAYTN]:    1,
    [ChainId.DOGECHAIN]: 1,
    [ChainId.CANTO]: 1,
  },
  multicall: {
    callResults: {}
  }
}

export function ChainStore({ children }) {
  const [chainState, setChainState] = useState(initialState)

  return (
    <ChainContext.Provider value={[chainState, setChainState]}>
      {children}
    </ChainContext.Provider>
  )
}

export const ChainContext = createContext([])

