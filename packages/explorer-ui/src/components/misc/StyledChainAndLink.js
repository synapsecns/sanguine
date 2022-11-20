import { Link } from 'react-router-dom'

import { ChainInfo } from '@components/misc/ChainInfo'
import { getChainUrl } from '@urls'
import { getNetworkTextHoverColor } from '@utils/styles/networks'

export function StyledChainAndLink({ chainId }) {
  return (
    <Link
      className={`${getNetworkTextHoverColor(chainId)} hover:underline`}
      to={getChainUrl({ chainId })}
    >
      <ChainInfo chainId={chainId} />
    </Link>
  )
}
