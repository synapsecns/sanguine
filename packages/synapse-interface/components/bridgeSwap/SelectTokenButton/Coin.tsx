import _ from 'lodash'

import { ELIGIBILITY_DEFAULT_TEXT } from '@/utils/hooks/useStipEligibility'
import { AvailableChains } from '@/components/bridgeSwap/SelectTokenButton/AvailableChains'

export const Coin = ({
  token,
  showAllChains,
  isOrigin,
  pausedChainIds,
  isEligible=false,

}: {
  token,
  showAllChains: boolean
  isOrigin: boolean
  pausedChainIds?: any
  isEligible?: boolean
}) => {

  return (
    <div className="flex-col text-left">
      <div className="text-lg text-primaryTextColor">{token?.symbol}</div>
      <div className="flex items-center space-x-2 text-xs text-secondaryTextColor">
        {isOrigin && isEligible ? (
          <div className="text-greenText">{ELIGIBILITY_DEFAULT_TEXT}</div>
        ) : (
          <div>{token?.name}</div>
        )}
        {showAllChains &&
          <AvailableChains
            token={token}
            pausedChainIds={pausedChainIds}
          />
        }
      </div>
    </div>
  )
}
