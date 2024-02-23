import type { Address } from 'viem'
import { usePortfolioActionHandlers } from '@/slices/portfolio/hooks'
import { shortenAddress } from '@/utils/shortenAddress'
import { ClearSearchButton } from './ClearSearchButton'

export const ViewSearchAddressBanner = ({
  viewingAddress,
}: {
  viewingAddress: Address
}) => {
  const { clearSearchResults } = usePortfolioActionHandlers()
  const shortened: string = shortenAddress(viewingAddress, 4)
  return (
    <div
      id="view-search-address-banner"
      className={`
        flex justify-between p-3 mb-3
        border border-synapsePurple rounded-sm
      `}
      style={{
        background:
          'linear-gradient(310.65deg, rgba(172, 143, 255, 0.2) -17.9%, rgba(255, 0, 255, 0.2) 86.48%)',
      }}
    >
      <div className="flex space-x-1">
        <div className="text-secondary ">Viewing</div>
        <div className="font-bold text-primary">{shortened}</div>
      </div>
      <ClearSearchButton onClick={clearSearchResults} show={true} />
    </div>
  )
}
