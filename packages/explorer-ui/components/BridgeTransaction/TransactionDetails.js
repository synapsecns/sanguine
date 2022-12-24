import {ClockIcon, IdentificationIcon} from '@heroicons/react/outline'

import {CopyButtonIcon} from '@components/CopyButtonIcon'
import {ContainerCard} from '@components/ContainerCard'
import {ChainInfo} from '@components/misc/ChainInfo'
import {IconAndAmount} from '@components/misc/IconAndAmount'
import {StyleHash} from '@components/misc/StyleHash'
import {StyleAddress} from '@components/misc/StyleAddress'

import {getNetworkShadow, getNetworkTextHoverColor} from '@utils/styles/networks'
import {unixTimestampToUTCString} from '@utils/unixTimestampToUTCString'
import {getChainUrl} from '@urls'
import Link from 'next/link'

export function TransactionDetails({ info, subtitle }) {
  return (
    <div className="items-center">
      <div className="flex items-center mb-2 sm:text-xs md:text-lg ">
        <div className="font-mono text-slate-400 ">
          <StyleHash sourceInfo={info} limiter={12} />
        </div>
        <span className="ml-2 sm:mt-1">
          <CopyButtonIcon
            text={info.hash}
            className="text-slate-600 hover:text-slate-300"
            tooltipText="transaction hash"
          />
        </span>
      </div>
      <ContainerCard
        className={`shadow-lg ${getNetworkShadow(
          info.chainId
        )} border-none text-2xl`}
        title={
          <Link
            className={`${getNetworkTextHoverColor(
              info.chainId
            )} hover:underline`}
            href={getChainUrl({ chainId: info.chainId })}
          >
            <ChainInfo chainId={info.chainId} imgClassName="w-7 h-7" />
          </Link>
        }
        subtitle={subtitle}
      >
        <div className="mt-10 mb-10 text-center">
          {info.formattedValue ? (
            <IconAndAmount
              formattedValue={info.formattedValue}
              tokenAddress={info.tokenAddress}
              chainId={info.chainId}
              tokenSymbol={info.tokenSymbol}
              iconSize="w-8 h-8"
              textSize="sm:text-3xl md:text-2xl lg:text-4xl"
            />
          ) : (
            <div className="w-full h-12 animate-pulse bg-slate-500" />
          )}
        </div>
        <div className="flex items-center text-sm">
          <IdentificationIcon className={`w-3 h-3 rounded-md mr-2 inline`} />
          <span>
            <StyleAddress sourceInfo={info} limiter={14} />
          </span>
        </div>
        <div className="flex items-center">
          <ClockIcon className={`w-3 h-3 rounded-md mr-2 inline`} />
          {info.time ? (
            <span className="text-sm">
              {unixTimestampToUTCString(info.time)}
            </span>
          ) : (
            <div className="w-40 h-3 animate-pulse bg-slate-500" />
          )}
        </div>
      </ContainerCard>
    </div>
  )
}
