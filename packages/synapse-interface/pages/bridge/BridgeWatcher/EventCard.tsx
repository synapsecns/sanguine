import { BridgeWatcherTx } from '@types'
import { CHAINS_BY_ID } from '@constants/chains'
import { ETH } from '@constants/tokens/master'
import { commify } from '@ethersproject/units'
import { formatBNToString } from '@bignumber/format'
import { BigNumber } from '@ethersproject/bignumber'
import ExplorerLink from './ExplorerLink'
import { getNetworkLinkTextColor } from '@styles/chains'
import { AddToWalletMiniButton } from '@components/buttons/AddToWalletButton'
import { getCoinTextColorCombined } from '@styles/tokens'
import { memo } from 'react'
import { formatTimestampToDate } from '@utils/time'

const EventCard = memo((event: BridgeWatcherTx) => {
  const chain = CHAINS_BY_ID[event.chainId]
  let showAddBtn
  if (event.token?.symbol == ETH.symbol) {
    showAddBtn = false
  } else {
    showAddBtn = true
  }

  return (
    <>
      <div className="pb-1 text-sm text-gray-500">
        {event?.timestamp && formatTimestampToDate(event.timestamp)}
      </div>
      <div className="flex items-center p-1 border border-gray-700 rounded-lg">
        <div className="flex-shrink-0">
          <img
            className="inline w-4 h-4 ml-1 mr-2 -mt-1 rounded"
            src={chain?.chainImg?.src}
          />
        </div>
        <div className="flex-grow">
          <div>
            <div className="w-full text-sm">
              {event && (
                <ExplorerLink
                  overrideExistingClassname={true}
                  className={`
                ${getNetworkLinkTextColor(chain?.color)}
                opacity-70 hover:opacity-100
              `}
                  chainId={event.chainId}
                  transactionHash={event.txHash}
                />
              )}
            </div>
          </div>
          <div className="w-full">
            <div className="w-full text-sm">
              <span className="font-medium text-gray-400">
                {event?.amount
                  ? commify(
                      formatBNToString(
                        event.amount,
                        event.token?.decimals[event.chainId] ?? 18,
                        8
                      )
                    )
                  : '0'}
              </span>

              {event.token && (
                <>
                  <span
                    className={`font-medium ${getCoinTextColorCombined(
                      event.token.color
                    )}`}
                  >
                    {' '}
                    {event.token.symbol}{' '}
                  </span>
                  <img
                    src={event.token.icon.src}
                    className="w-4 h-4 inline -mt-0.5 rounded-sm"
                  />
                </>
              )}
            </div>
          </div>
        </div>
        <div className="flex-shrink">
          {showAddBtn && event.isFrom && (
            <AddToWalletMiniButton
              token={event.token}
              icon={event.token?.icon?.src}
              chainId={event.chainId}
              className={`float-right inline-block `}
            />
          )}
        </div>
      </div>
    </>
  )
})

export default EventCard
