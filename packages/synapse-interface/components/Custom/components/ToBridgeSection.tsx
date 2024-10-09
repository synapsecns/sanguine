import { useAccount } from 'wagmi'

import { CHAINS_BY_ID } from '@/constants/chains'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { ImageOverlayComponent } from './ImageOverlay'

export const ToBridgeSection = ({ fromChainId, fromToken, toChainId }) => {
  const { isConnected } = useAccount()

  const { balances } = usePortfolioState()

  const fromChainBalances = balances[fromChainId]
  const fromTokenBalance = fromChainBalances?.find(
    (token) => token.tokenAddress === fromToken.addresses[fromChainId]
  ).parsedBalance

  return (
    <div className="rounded-md bg-zinc-100 dark:bg-bgBase">
      <div className="flex items-center p-3 space-x-2 text-lg">
        <img
          src={fromToken?.icon.src}
          alt={fromToken?.symbol}
          className="w-6 h-6"
        />
        <div>To Bridge</div>
      </div>
      <div className="border-b border-zinc-200 dark:border-zinc-700"></div>
      {isConnected ? (
        <div className="flex items-center p-3 space-x-3 text-lg">
          <ImageOverlayComponent
            bigImageSrc={fromToken?.icon.src}
            smallImageSrc={CHAINS_BY_ID[fromChainId]?.chainImg.src}
            altTextBig={fromToken?.symbol}
            altTextSmall={CHAINS_BY_ID[fromChainId]?.name}
          />
          <div className="text-sm opacity-75">
            {fromTokenBalance}/{fromToken.symbol} from{' '}
            {CHAINS_BY_ID[fromChainId].name} to {CHAINS_BY_ID[toChainId].name}
          </div>
        </div>
      ) : (
        <div className="p-3 opacity-75">
          Connect your wallet to see bridgeable tokens
        </div>
      )}
    </div>
  )
}
