import { useCallback, useContext } from 'react'
import { ChainContext } from '@store/ChainStore'


export function useMulticallState() {
  let [chainState, setChainState] = useContext(ChainContext)
  // console.log("%cuseMulticallState", "color: teal")

  const setMulticallState = useCallback(
    (multicallStateVal) => {
      // console.log({ multicallStateVal, chainStateMulticall: chainState.multicall})
      setChainState({
        ...chainState,
        multicall: {
          ...chainState.multicall,
          ...multicallStateVal
        }
      })
    },
    [
      chainState.multicall
    ]
  )

  return [chainState.multicall, setMulticallState]
}
