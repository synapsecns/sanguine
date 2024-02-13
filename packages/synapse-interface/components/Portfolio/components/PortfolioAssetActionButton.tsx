type PortfolioAssetActionButtonProps = {
  selectCallback: () => void
  isDisabled: boolean
  isSelected: boolean
}

export const PortfolioAssetActionButton = ({
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
          border border-white/10 ${isSelected && "!border-synapsePurple"}
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
