import { ArrowNarrowRightIcon } from '@heroicons/react/outline'

import Grid from '@tw/Grid'

import { ChainImage } from '@components/misc/ChainImage'
import { StyleAddress } from '@components/misc/StyleAddress'
import { StyleHash } from '@components/misc/StyleHash'
import { IconAndAmount } from '@components/misc/IconAndAmount'
import { timeAgo } from '@utils/timeAgo'

export function InfoRow({ txLabel, addrLabel, info }) {
  const isDestination = txLabel === 'Destination'
  let margin
  let textClassName

  if (isDestination) {
    textClassName = 'text-sm opacity-70'
  }

  if (!isDestination) {
    margin = 'mb-1'
  }

  return (
    <Grid
      gap={2}
      cols={{ sm: 1, lg: 4 }}
      className={`pb-1 text-white whitespace-nowrap ${margin} ${textClassName}`}
    >
      <InfoItem isDestination={isDestination} showArrow={true}>
        <ChainImage chainId={info.chainId} imgSize="w-5 h-5" />
        <span className="mr-1">{txLabel}:</span>
        <StyleHash sourceInfo={info} />
      </InfoItem>
      <InfoItem isDestination={isDestination}>
        <IconAndAmount
          formattedValue={info.formattedValue}
          tokenAddress={info.tokenAddress}
          chainId={info.chainId}
          tokenSymbol={info.tokenSymbol}
          iconSize="w-4 h-4"
          textSize={isDestination ? 'text-sm' : 'text-base'}
        />
      </InfoItem>
      <InfoItem isDestination={isDestination}>
        <StyleAddress sourceInfo={info} />
      </InfoItem>
      <InfoItem isDestination={isDestination}>
        {addrLabel}: {timeAgo({ timestamp: info.time })}
      </InfoItem>
    </Grid>
  )
}

function InfoItem({ isDestination, showArrow = false, children }) {
  let style
  if (showArrow) {
    style = 'w-6 h-6 mr-2 text-gray-400'
  } else {
    style = 'w-6 h-6 mr-2 text-transparent'
  }
  return (
    <div className="flex items-center">
      {isDestination ? (
        <ArrowNarrowRightIcon className={`w-6 h-6 mr-2 ${style}`} />
      ) : (
        ''
      )}
      {children}
    </div>
  )
}
