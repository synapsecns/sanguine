import _ from 'lodash'
import { CHAIN_INFO_MAP } from '@constants/networks'
import { TOKEN_HASH_MAP } from '@constants-new/tokens/index'
import { getCoinTextColor } from '@utils/styles/coins'
import { getNetworkTextColor } from '@utils/styles/networks'
import Grid from '@components/tailwind/Grid'
import Tooltip from '@components/tailwind/Tooltip'

import { AssetImage } from './AssetImage'
import { ChainImage } from './ChainImage'

export function MostActive({ data }) {
  return (
    <div className="flex justify-center my-2 text-sm">
      {data.slice(0, 5).map((chainTokenProps, i) => (
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
    <Tooltip
      content={
        <Grid cols={{ xs: 1 }}>
          <div className="opacity-100">
            <AssetImage
              tokenAddress={tokenAddress}
              chainId={chainId}
              className="!w-4 !h-4 !opacity-100"
            />
            <span className={getCoinTextColor(token)}>{name}</span> on{' '}
            <ChainImage chainId={chainId} className="ml-1 mr-1" />
            <span className={getNetworkTextColor(chainId)}>{chainName}</span>
          </div>
        </Grid>
      }
      className="!opacity-100"
      tooltipClassName="!-mt-16 !-ml-16"
    >
      <div className="flex mr-3 hover:opacity-50">
        <AssetImage
          tokenAddress={tokenAddress}
          chainId={chainId}
          className="w-8 h-8"
        />
        <div className="z-10 mt-3 -ml-3">
          <ChainImage chainId={chainId} className="mr-0" />
        </div>
      </div>
    </Tooltip>
  )
}
