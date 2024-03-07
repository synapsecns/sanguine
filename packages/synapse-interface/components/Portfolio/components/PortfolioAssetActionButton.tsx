import { Token } from "@/utils/types"
import { getBorderStyleForCoin } from "@/styles/tokens"
type PortfolioAssetActionButtonProps = {
  token: Token
  selectCallback: () => void
  isDisabled: boolean
  isSelected: boolean
}

export const PortfolioAssetActionButton = ({
  token,
  selectCallback,
  isDisabled,
  isSelected,
}: PortfolioAssetActionButtonProps) => {
  return (
    <>
      <button
        id="portfolio-asset-action-button"
        className={`
          py-1 px-6 rounded-md text-sm
          border ${isSelected ? getBorderStyleForCoin(token.color): "border-white/10"}
          ${!isDisabled && 'cursor-pointer hover:bg-bgBase/20 active:border-white/30'}
        `}
        onClick={selectCallback}
        disabled={isDisabled}
      >
        Select{isSelected && 'ed'}
      </button>
    </>
  )
}
