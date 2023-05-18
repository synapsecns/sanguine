import {ChainInfo} from '@components/misc/ChainInfo'
import {getChainUrl} from '@urls'
import {getNetworkTextHoverColor} from '@utils/styles/networks'

export function StyledChainAndLink({ chainId }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'Link'.
    <Link
      // @ts-expect-error TS(2304): Cannot find name 'className'.
      className={`${getNetworkTextHoverColor(chainId)} hover:underline`}
      // @ts-expect-error TS(2304): Cannot find name 'to'.
      to={getChainUrl({ chainId })}
    >
      // @ts-expect-error TS(2749): 'ChainInfo' refers to a value, but is being used a... Remove this comment to see the full error message
      <ChainInfo chainId={chainId} />
    </Link>
  )
}
