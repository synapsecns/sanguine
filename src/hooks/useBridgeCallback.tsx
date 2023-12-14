import { stringToBigInt } from '@/utils/stringToBigInt'
import { ZeroAddress } from 'ethers'
import { useState, useCallback, useMemo } from 'react'

export enum UseBridgeCallbackError {
  REQUIRE_TOKEN_ADDRESS = 'Bridge: Missing Token Address',
  REQUIRE_ORIGIN_ROUTER_ADDRESS = 'Bridge: Missing Router Address',
  REQUIRE_DESTINATION_ADDRESS = 'Bridge: Missing Destination Address',
  REQUIRE_ORIGIN_CHAIN = 'Bridge: Missing Origin ChainId',
  REQUIRE_DESTINATION_CHAIN = 'Bridge: Missing Destination ChainId',
  REQUIRE_ORIGIN_QUERY = 'Bridge: Missing Quote Origin Query',
  REQUIRE_DESTINATION_QUERY = 'Bridge: Missing Quote Destination Query',
  REQUIRE_SYNAPSE_SDK = 'Bridge: Missing Synapse SDK',
  REQUIRE_SIGNER = 'Bridge: Missing Signer',
}

export enum BridgeCallbackState {
  PENDING = 'pending',
  SUCCESS = 'success',
  IDLE = 'idle',
}

export interface UseBridgeCallbackArgs {
  destinationAddress: string
  originRouterAddress: string
  originChainId: number
  destinationChainId: number
  tokenAddress: string
  amount: bigint
  originQuery: any
  destinationQuery: any
  synapseSDK?: any
  signer?: any
}

export function useBridgeCallback({
  destinationAddress,
  originRouterAddress,
  originChainId,
  destinationChainId,
  tokenAddress,
  amount,
  originQuery,
  destinationQuery,
  synapseSDK,
  signer,
}: UseBridgeCallbackArgs) {
  const [bridgeState, setBridgeState] = useState<BridgeCallbackState>(
    BridgeCallbackState.IDLE
  )

  const BridgeStateCallback = {
    initiateBridge: async () => {
      setBridgeState(BridgeCallbackState.PENDING)
    },
    successBridge: async () => {
      setBridgeState(BridgeCallbackState.SUCCESS)
    },
    resetState: () => {
      setBridgeState(BridgeCallbackState.IDLE)
    },
  }

  const bridgeCallback = useCallback(async () => {
    try {
      /**
       * Dependency Guardrail checks
       * Bail on executing Bridge Callback if dependencies missing
       * Set respective error when dependency missing
       */
      if (!tokenAddress) {
        throw new Error(UseBridgeCallbackError.REQUIRE_TOKEN_ADDRESS)
      } else if (!originRouterAddress) {
        throw new Error(UseBridgeCallbackError.REQUIRE_ORIGIN_ROUTER_ADDRESS)
      } else if (!destinationAddress) {
        throw new Error(UseBridgeCallbackError.REQUIRE_DESTINATION_ADDRESS)
      } else if (!originChainId) {
        throw new Error(UseBridgeCallbackError.REQUIRE_ORIGIN_CHAIN)
      } else if (!destinationChainId) {
        throw new Error(UseBridgeCallbackError.REQUIRE_DESTINATION_CHAIN)
      } else if (!originQuery) {
        throw new Error(UseBridgeCallbackError.REQUIRE_ORIGIN_QUERY)
      } else if (!destinationQuery) {
        throw new Error(UseBridgeCallbackError.REQUIRE_DESTINATION_QUERY)
      } else if (!synapseSDK) {
        throw new Error(UseBridgeCallbackError.REQUIRE_SYNAPSE_SDK)
      } else if (!signer) {
        throw new Error(UseBridgeCallbackError.REQUIRE_SIGNER)
      }

      BridgeStateCallback.initiateBridge()

      const data = await synapseSDK.bridge(
        destinationAddress,
        originRouterAddress,
        originChainId,
        destinationChainId,
        tokenAddress,
        amount,
        originQuery,
        destinationQuery
      )

      const payload =
        tokenAddress === ZeroAddress
          ? {
              data: data.data,
              to: data.to,
              value: amount,
            }
          : data

      const transactionHash = await signer.sendTransaction(payload)

      BridgeStateCallback.successBridge()
      return transactionHash
    } catch (error) {
      BridgeStateCallback.resetState()
      console.error(error)
    }
  }, [
    destinationAddress,
    originRouterAddress,
    originChainId,
    destinationChainId,
    tokenAddress,
    amount,
    originQuery,
    destinationQuery,
    synapseSDK,
    signer,
  ])

  return useMemo(() => {
    /**
     * Return error state based on current hook props
     * Callback will throw error when invoked with missing dependencies
     */
    if (!tokenAddress) {
      return {
        state: bridgeState,
        callback: () => bridgeCallback(),
        error: UseBridgeCallbackError.REQUIRE_TOKEN_ADDRESS,
      }
    } else if (!originRouterAddress) {
      return {
        state: bridgeState,
        callback: () => bridgeCallback(),
        error: UseBridgeCallbackError.REQUIRE_ORIGIN_ROUTER_ADDRESS,
      }
    } else if (!destinationAddress) {
      return {
        state: bridgeState,
        callback: () => bridgeCallback(),
        error: UseBridgeCallbackError.REQUIRE_DESTINATION_ADDRESS,
      }
    } else if (!originChainId) {
      return {
        state: bridgeState,
        callback: () => bridgeCallback(),
        error: UseBridgeCallbackError.REQUIRE_ORIGIN_CHAIN,
      }
    } else if (!destinationChainId) {
      return {
        state: bridgeState,
        callback: () => bridgeCallback(),
        error: UseBridgeCallbackError.REQUIRE_DESTINATION_CHAIN,
      }
    } else if (!originQuery) {
      return {
        state: bridgeState,
        callback: () => bridgeCallback(),
        error: UseBridgeCallbackError.REQUIRE_ORIGIN_QUERY,
      }
    } else if (!destinationQuery) {
      return {
        state: bridgeState,
        callback: () => bridgeCallback(),
        error: UseBridgeCallbackError.REQUIRE_DESTINATION_QUERY,
      }
    } else if (!synapseSDK) {
      return {
        state: bridgeState,
        callback: () => bridgeCallback(),
        error: UseBridgeCallbackError.REQUIRE_SYNAPSE_SDK,
      }
    } else if (!signer) {
      return {
        state: bridgeState,
        callback: () => bridgeCallback(),
        error: UseBridgeCallbackError.REQUIRE_SIGNER,
      }
    }

    return {
      state: bridgeState,
      callback: () => bridgeCallback(),
      error: null,
    }
  }, [
    destinationAddress,
    originRouterAddress,
    originChainId,
    destinationChainId,
    tokenAddress,
    amount,
    originQuery,
    destinationQuery,
    synapseSDK,
    signer,
    bridgeCallback,
    bridgeState,
  ])
}
