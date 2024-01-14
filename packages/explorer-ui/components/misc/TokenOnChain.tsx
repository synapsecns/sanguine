import _ from 'lodash'
import { TOKEN_HASH_MAP, CHAINS } from '@synapse-constants'
import { getCoinTextColor } from '@utils/styles/coins'
import { getNetworkTextColor } from '@utils/styles/networks'
import { getChainUrl } from '@urls'

import { AssetImage } from './AssetImage'
import { ChainImage } from './ChainImage'

const CHAINS_BY_ID = CHAINS.CHAINS_BY_ID

export function TokenOnChain({ tokenAddress, chainId }) {
  const chainName = CHAINS_BY_ID[chainId].name
  const token = TOKEN_HASH_MAP[chainId][tokenAddress]

  return (
    <>
      <AssetImage
        tokenAddress={tokenAddress}
        chainId={chainId}
        className="w-8 h-8"
      />
      <div className="z-10 mt-3 -ml-3">
        <ChainImage chainId={chainId} className="mr-0" />
      </div>
      <span className={`${getCoinTextColor(token)}} ml-2 mr-2`}>{token?.symbol}</span>
      <span className="mr-2 text-white font-extralight">on</span>
      <span className={getNetworkTextColor(chainId)}>
        <a className="hover:underline" href={getChainUrl({ chainId })}>
          {chainName}
        </a>
      </span>
    </>
  )
}
