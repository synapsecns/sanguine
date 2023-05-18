import _ from 'lodash'

import {CHAIN_INFO_MAP} from '@constants/networks'
import {TOKEN_HASH_MAP} from '@constants/tokens/basic'

import {getCoinTextColor} from '@utils/styles/coins'
import {getNetworkTextColor} from '@utils/styles/networks'
import {getChainUrl} from '@urls'

import {AssetImage} from './AssetImage'
import {ChainImage} from './ChainImage'

export function TokenOnChain({ tokenAddress, chainId }) {
  const { chainName } = CHAIN_INFO_MAP[chainId]
  const token = TOKEN_HASH_MAP[chainId][_.toLower(tokenAddress)] ?? {}
  const { name } = token

  return (
    <>
      // @ts-expect-error TS(2749): 'AssetImage' refers to a value, but is being used ... Remove this comment to see the full error message
      <AssetImage
        tokenAddress={tokenAddress}
        chainId={chainId}
        // @ts-expect-error TS(2304): Cannot find name 'className'.
        className="w-8 h-8"
      />
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="z-10 mt-3 -ml-3">
        // @ts-expect-error TS(2749): 'ChainImage' refers to a value, but is being used ... Remove this comment to see the full error message
        <ChainImage chainId={chainId} className="mr-0" />
      </div>
      // @ts-expect-error TS(2304): Cannot find name 'span'.
      <span className={`${getCoinTextColor(token)}} ml-2 mr-2`}>{name}</span>
      // @ts-expect-error TS(2304): Cannot find name 'span'.
      <span className="mr-2 text-white font-extralight">on</span>
      // @ts-expect-error TS(2304): Cannot find name 'span'.
      <span className={getNetworkTextColor(chainId)}>
        // @ts-expect-error TS(2304): Cannot find name 'a'.
        <a className="hover:underline" href={getChainUrl({ chainId })}>
          // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
          {chainName}
        </a>
      </span>
    </>
  )
}
