import _ from 'lodash'

import {CHAIN_INFO_MAP} from '@constants/networks'
import {TOKEN_HASH_MAP} from '@constants/tokens/basic'

import {getCoinTextColor} from '@utils/styles/coins'
import {getNetworkTextColor} from '@utils/styles/networks'

import Grid from '@components/tailwind/Grid'
import Tooltip from '@components/tailwind/Tooltip'

import {AssetImage} from './AssetImage'
import {ChainImage} from './ChainImage'

export function MostActive({ data }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    <div className="flex justify-center my-2 text-sm">
      {data.slice(0, 5).map((chainTokenProps, i) => (
        // @ts-expect-error TS(2749): 'MostActiveTokenChainIcon' refers to a value, but ... Remove this comment to see the full error message
        <MostActiveTokenChainIcon {...chainTokenProps} key={i} />
      ))}
    </div>
  )
}

function MostActiveTokenChainIcon({ tokenAddress, chainId }) {
  const { chainName, chainLogo } = CHAIN_INFO_MAP[chainId]
  const token = TOKEN_HASH_MAP[chainId][_.toLower(tokenAddress)] ?? {}
  const { name } = token

  return (
    // @ts-expect-error TS(2749): 'Tooltip' refers to a value, but is being used as ... Remove this comment to see the full error message
    <Tooltip
      // @ts-expect-error TS(2304): Cannot find name 'content'.
      content={
        // @ts-expect-error TS(2304): Cannot find name 'cols'.
        <Grid cols={{ xs: 1 }}>
          // @ts-expect-error TS(2304): Cannot find name 'div'.
          <div className="opacity-100">
            // @ts-expect-error TS(2749): 'AssetImage' refers to a value, but is being used ... Remove this comment to see the full error message
            <AssetImage
              // @ts-expect-error TS(2304): Cannot find name 'tokenAddress'.
              tokenAddress={tokenAddress}
              // @ts-expect-error TS(2304): Cannot find name 'chainId'.
              chainId={chainId}
              // @ts-expect-error TS(2304): Cannot find name 'className'.
              className="!w-4 !h-4 !opacity-100"
            />
            // @ts-expect-error TS(2304): Cannot find name 'span'.
            <span className={getCoinTextColor(token)}>{name}</span> on{' '}
            // @ts-expect-error TS(2304): Cannot find name 'chainId'.
            <ChainImage chainId={chainId} className="ml-1 mr-1" />
            // @ts-expect-error TS(2304): Cannot find name 'span'.
            <span className={getNetworkTextColor(chainId)}>{chainName}</span>
          </div>
        </Grid>
      }
      // @ts-expect-error TS(2304): Cannot find name 'className'.
      className="!opacity-100"
      // @ts-expect-error TS(2304): Cannot find name 'tooltipClassName'.
      tooltipClassName="!-mt-16 !-ml-16"
    >
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="flex mr-3 hover:opacity-50">
        // @ts-expect-error TS(2749): 'AssetImage' refers to a value, but is being used ... Remove this comment to see the full error message
        <AssetImage
          // @ts-expect-error TS(2304): Cannot find name 'tokenAddress'.
          tokenAddress={tokenAddress}
          // @ts-expect-error TS(2304): Cannot find name 'chainId'.
          chainId={chainId}
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className="w-8 h-8"
        />
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="z-10 mt-3 -ml-3">
          // @ts-expect-error TS(2749): 'ChainImage' refers to a value, but is being used ... Remove this comment to see the full error message
          <ChainImage chainId={chainId} className="mr-0" />
        </div>
      </div>
    </Tooltip>
  )
}
