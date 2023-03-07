import { useGenericMulticall2Contract } from "@hooks/contracts/useMulticallContract"
import { useSingleCallResult } from "@hooks/multicall"


export function useBlockHeight(chainId) {
  const multicallContract = useGenericMulticall2Contract(chainId)
  const blockHeightResult = useSingleCallResult(
    chainId,
    multicallContract,
    'getBlockNumber',
    [],
    { resultOnly: true }
  )

  return blockHeightResult?.blockNumber?.toNumber() ?? 0
}
