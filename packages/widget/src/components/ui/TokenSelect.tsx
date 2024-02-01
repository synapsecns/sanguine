import _ from 'lodash'
import { BridgeableToken } from 'types'

import { useBridgeState } from '@/state/slices/bridge/hooks'
import { BridgeState } from '@/state/slices/bridge/reducer'
import { TokenPopoverSelect } from '@/components/ui/TokenPopoverSelect'
import { useWalletState } from '@/state/slices/wallet/hooks'
import { getRoutePossibilities } from '@/utils/routeMaker/generateRoutePossibilities'

type Props = {
  label: 'In' | 'Out'
  isOrigin: boolean
  onChange: (newToken: BridgeableToken) => void
  token: BridgeableToken
}

export const TokenSelect = ({ label, isOrigin, token, onChange }: Props) => {
  const {
    originChainId,
    destinationChainId,
    originTokens,
    destinationTokens,
    targetTokens,
    targetChainIds,
  }: BridgeState = useBridgeState()

  const { balances } = useWalletState()

  let options
  let remaining

  if (isOrigin) {
    options = originTokens

    if (targetChainIds?.length > 0 && targetChainIds.includes(originChainId)) {
      remaining = generateOriginRemainingTokens(
        originChainId,
        destinationChainId,
        targetTokens
      )

      remaining = _.differenceWith(
        remaining,
        options,
        (a: { routeSymbol: string }, b: { routeSymbol: string }) =>
          a.routeSymbol === b.routeSymbol
      )
    } else {
      remaining = generateOriginRemainingTokens(
        originChainId,
        destinationChainId,
        []
      )

      remaining = _.differenceWith(
        remaining,
        options,
        (a: { routeSymbol: string }, b: { routeSymbol: string }) =>
          a.routeSymbol === b.routeSymbol
      )
    }
  } else {
    if (
      targetChainIds?.length > 0 &&
      targetChainIds.includes(destinationChainId)
    ) {
      /* If consumer provides no target tokens, then we just show the destination tokens */
      /* If consumer does provide target tokens, then we only the destination tokens that are included in the target tokens */
      options = _.isEmpty(targetTokens)
        ? destinationTokens
        : _.filter(destinationTokens, (dt) =>
            _.some(targetTokens, (tt) => tt.routeSymbol === dt.routeSymbol)
          )

      /* determine possible remaining tokens and filter out options */
      remaining = generateDestinationRemainingTokens(
        originChainId,
        destinationChainId,
        targetTokens
      )

      remaining = _.differenceWith(
        remaining,
        options,
        (a: { routeSymbol: string }, b: { routeSymbol: string }) =>
          a.routeSymbol === b.routeSymbol
      )
    } else {
      options = destinationTokens
      remaining = generateDestinationRemainingTokens(
        originChainId,
        destinationChainId,
        []
      )

      remaining = _.differenceWith(
        remaining,
        options,
        (a: { routeSymbol: string }, b: { routeSymbol: string }) =>
          a.routeSymbol === b.routeSymbol
      )
    }
  }

  return (
    <TokenPopoverSelect
      options={options}
      remaining={remaining}
      balances={isOrigin ? balances : []}
      onSelect={(selected) => {
        onChange(selected)
      }}
      selected={token}
    />
  )
}

const generateOriginRemainingTokens = (
  originChainId,
  destinationChainId,
  targetTokens
) => {
  if (targetTokens && targetTokens.length > 0) {
    const fromTokensArrays = targetTokens.map((targetToken) => {
      const { fromTokens } = getRoutePossibilities({
        fromChainId: originChainId,
        fromToken: null,
        toChainId: destinationChainId,
        toToken: targetToken,
      })
      return fromTokens
    })

    const combinedFromTokens = fromTokensArrays.flat()

    return _.uniq(combinedFromTokens)
  } else {
    const { fromTokens } = getRoutePossibilities({
      fromChainId: originChainId,
      fromToken: null,
      toChainId: destinationChainId,
      toToken: null,
    })

    return fromTokens
  }
}

const generateDestinationRemainingTokens = (
  originChainId,
  destinationChainId,
  targetTokens
) => {
  if (targetTokens && targetTokens.length > 0) {
    const toTokensArrays = targetTokens.map((targetToken) => {
      const { toTokens } = getRoutePossibilities({
        fromChainId: originChainId,
        fromToken: null,
        toChainId: destinationChainId,
        toToken: targetToken,
      })
      return toTokens
    })

    const combinedToTokens = toTokensArrays.flat()

    return _.uniq(combinedToTokens)
  } else {
    const { toTokens } = getRoutePossibilities({
      fromChainId: originChainId,
      fromToken: null,
      toChainId: destinationChainId,
      toToken: null,
    })

    return toTokens
  }
}
