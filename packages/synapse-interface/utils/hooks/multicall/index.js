import _ from 'lodash'
import { useEffect, useMemo } from 'react'

import { BigNumber } from '@ethersproject/bignumber'
import { Contract } from '@ethersproject/contracts'

import { FunctionFragment, Interface } from '@ethersproject/abi'
import { addMulticallListeners, removeMulticallListeners } from '@utils/multicall/actions'

import { useBlockNumber } from '@hooks/useBlockNumber'
import { useMulticallState } from '@hooks/multicall/useMulticallState'

import { parseCallKey, toCallKey } from '@utils/multicall/utils'



function isMethodArg(x) {
  return BigNumber.isBigNumber(x) || ['string', 'number'].indexOf(typeof x) !== -1
}

function isValidMethodArgs(x) {
  return (
    x === undefined ||
    (Array.isArray(x) && x.every((xi) => isMethodArg(xi) || (Array.isArray(xi) && xi.every(isMethodArg))))
  )
}


const INVALID_RESULT = {
  valid:       false,
  blockNumber: undefined,
  data:        undefined
}

// use this options object
export const NEVER_RELOAD = {
  blocksPerFetch: Infinity,
}

// the lowest level call for subscribing to contract data
function useCallsData(chainId, calls, options) {
  const { blocksPerFetch=1 } = options ?? {}
  // const { chainId } = useActiveWeb3React()
  // console.log(chainId)
  const [multicallState, setMulticallState ] = useMulticallState()

  const callResults = multicallState.callResults

  const serializedCallKeys = useMemo(
    () => {
      return JSON.stringify(
        calls
          ?.filter((c) => Boolean(c))
          ?.map(toCallKey)
          ?.sort() ?? []
      )
    },
    [calls, chainId]
  )

  // update listeners when there is an actual change that persists for at least 100ms
  useEffect(
    () => {
      const callKeys = JSON.parse(serializedCallKeys)

      if (!chainId || callKeys.length === 0) return undefined
      const calls = callKeys.map((key) => parseCallKey(key))

      setMulticallState(
        addMulticallListeners(
          multicallState,
          {
            chainId,
            calls,
            options: { blocksPerFetch },
          }
        )
      )

      return () => {
        setMulticallState(
          removeMulticallListeners(multicallState, {
            chainId,
            calls,
            options: { blocksPerFetch },
          })
        )
      }
    },
    [chainId, blocksPerFetch, serializedCallKeys] // ,setMulticallState
  )
  return useMemo(
    () => {
      const callResultsArr = calls.map((call) => {
        if (!chainId || !call) return INVALID_RESULT

        const result = callResults[chainId]?.[toCallKey(call)]

        let data
        if (result?.data && result?.data !== '0x') {
          data = result.data
        }

        return {
          valid: true,
          data,
          blockNumber: result?.blockNumber
        }
      })

      return callResultsArr
    },
    [chainId, callResults, calls]
  )
}


const INVALID_CALL_STATE = {
  valid:   false,
  result:  undefined,
  loading: false,
  syncing: false,
  error:   false
}
const LOADING_CALL_STATE = {
  valid:   true,
  result:  undefined,
  loading: true,
  syncing: true,
  error:   false
}
/**
 * @param {CallResult} callResult
 * @param {Interface} contractInterface
 * @param {FunctionFragment} fragment
 * @param {number} latestBlockNumber
 */
function toCallState(callResult, contractInterface, fragment, latestBlockNumber) {
  // console.log("%ctoCallState", "color:blue")
  if (!callResult) return INVALID_CALL_STATE
  const { valid, data, blockNumber } = callResult

  if (!valid) return INVALID_CALL_STATE
  if (valid && !blockNumber) return LOADING_CALL_STATE
  if (!contractInterface || !fragment || !latestBlockNumber) return LOADING_CALL_STATE
  const success = data && data.length > 2
  const syncing = (blockNumber ?? 0) < latestBlockNumber
  let result = undefined
  if (success && data) {
    try {
      result = contractInterface.decodeFunctionResult(fragment, data)
    } catch (error) {
      console.debug('Result data parsing failed', fragment, data)
      return {
        valid: true,
        loading: false,
        error: true,
        syncing,
        result,
      }
    }
  }
  return {
    valid: true,
    loading: false,
    syncing,
    result: result,
    error: !success,
  }
}

/**
 * @param {number} chainId
 * @param {Contract} contract
 * @param {string} methodName
 * @param callInputs
 * @param options
 * @param {number} gasRequired
 */
export function useSingleContractMultipleData(
  chainId,
  contract,
  methodName,
  callInputs,
  options={},
  gasRequired
) {
  const fragment = useMemo(
    () => contract?.interface?.getFunction(methodName),
    [contract, methodName]
  )

  const calls = useMemo(
    () => {
      if (contract && fragment && callInputs?.length > 0 && callInputs.every((inputs) => isValidMethodArgs(inputs)) ) {
        return callInputs.map((inputs) => {
          return {
            address: contract.address,
            callData: contract.interface.encodeFunctionData(fragment, inputs),
            ...(gasRequired ? { gasRequired } : {}),
          }
        })
      } else {
        return []
      }
    },
    [contract, fragment, callInputs, gasRequired]
  )

  const results = useCallsData(chainId, calls, options)

  const [latestBlockNumber, setBlockNumber] = useBlockNumber(chainId)

  return useMemo(() => {
    return results.map(result => {
      const item = toCallState(result, contract?.interface, fragment, latestBlockNumber)
      if (options.resultOnly) {
        return item.result
      } else {
        return item
      }
    })
  }, [fragment, contract, results, latestBlockNumber, chainId])
}


/**
 * @param {number} chainId
 * @param {string[]} addresses
 * @param {Interface} contractInterface
 * @param {string} methodName
 * @param callInputs
 * @param options
 * @param {number} gasRequired
 */
export function useMultipleContractSingleData(
  chainId,
  addresses,
  contractInterface,
  methodName,
  callInputs,
  options={},
  gasRequired
) {
  const fragment = useMemo(
    () => contractInterface.getFunction(methodName),
    [contractInterface, methodName]
  )
  const callData = useMemo(
    () =>
      fragment && isValidMethodArgs(callInputs)
        ? contractInterface.encodeFunctionData(fragment, callInputs)
        : undefined,
    [callInputs, contractInterface, fragment, chainId]
  )

  const calls = useMemo(
    () => {
      if (fragment && addresses && addresses?.length > 0 && callData) {
        return addresses.map((address) => {
          if (address && callData) {
            return {
              address,
              callData,
              ...(gasRequired ? { gasRequired } : {}),
            }
          } else {
            return undefined
          }
        })
      } else {
        return []
      }
    },
    [addresses, callData, fragment, gasRequired]
  )

  const results = useCallsData(chainId, calls, options)

  const [latestBlockNumber, setBlockNumber] = useBlockNumber(chainId)

  return useMemo(() => {
    return results.map(result => {
      const item = toCallState(result, contractInterface, fragment, latestBlockNumber)
      if (options.resultOnly) {
        return item.result
      } else {
        return item
      }
    })
  }, [fragment, results, contractInterface, latestBlockNumber, chainId])
}


/**
 * @param {Contract} contract
 * @param {string} methodName
 * @param inputs
 * @param options
 * @param {number} gasRequired
 */
export function useSingleCallResult(
  chainId,
  contract,
  methodName,
  inputs,
  options={},
  gasRequired
) {
  const fragment = useMemo(
    () => contract?.interface?.getFunction(methodName),
    [contract, methodName]
  )

  const calls = useMemo(
    () => {
      if (contract && fragment && isValidMethodArgs(inputs)) {
        return [
          {
            address: contract.address,
            callData: contract.interface.encodeFunctionData(fragment, inputs),
            ...(gasRequired ? { gasRequired } : {}),
          },
        ]
      }
      return []
    },
    [contract, fragment, inputs, gasRequired]
  )

  const result = useCallsData(chainId, calls, options)[0]

  const [latestBlockNumber, setBlockNumber] = useBlockNumber(chainId)

  return useMemo(() => {
    const item =  toCallState(result, contract?.interface, fragment, latestBlockNumber)
    if (options.resultOnly) {
      return item.result
    } else {
      return item
    }
  }, [result, contract?.interface, fragment, latestBlockNumber, chainId])
}


/**
 * @param {number} chainId
 * @param {Contract} contract
 * @param {Object.<string, string[]>} methodsAndCalls
 * @param callInputs
 * @param options
 * @param {number} gasRequired
 */
export function useSingleContractMultipleMethods(
  chainId,
  contract,
  methodsAndCalls,
  options={},
  gasRequired
) {
  const methodsAndCallsPaired = useMemo(
    () => Object.entries(methodsAndCalls),
    [methodsAndCalls]
  )
  const methods = useMemo(
    () => methodsAndCallsPaired.map( mc => mc[0]),
    [methodsAndCallsPaired]
  )
  const callsByMethods = useMemo(
    () => methodsAndCallsPaired.map(mc => mc[1]),
    [methodsAndCallsPaired]
  )

  const fragments = useMemo(
    () => methods.map( methodName => contract?.interface?.getFunction(methodName)),
    [contract, methods]
  )

  const calls = useMemo(
    () => {
      if (
        contract &&
        fragments?.length > 0 &&
        callsByMethods?.length > 0 &&
        callsByMethods.every(inputs => isValidMethodArgs(inputs))
      ) {
        return _.zip(fragments, callsByMethods).map( ([fragment, inputs]) => {
          return {
            address: contract.address,
            callData: contract.interface.encodeFunctionData(fragment, inputs),
            ...(gasRequired ? { gasRequired } : {}),
          }
        })
      } else {
        return []
      }
    },
    [contract, fragments, callsByMethods, gasRequired]
  )

  const results = useCallsData(chainId, calls, options)

  const [latestBlockNumber, setBlockNumber] = useBlockNumber(chainId)

  return useMemo(() => {
    return _.zip(fragments, results).map(([fragment, result]) => {
      const item = toCallState(result, contract?.interface, fragment, latestBlockNumber)
      if (options.resultOnly) {
        return item.result
      } else {
        return item
      }
    })
  }, [fragments, contract, callsByMethods, results, latestBlockNumber, chainId])
}