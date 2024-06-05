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
          py-1 px-6 rounded-sm
          border border-synapsePurple
          ${!isDisabled && 'cursor-pointer hover:bg-surface active:opacity-70'}
        `}
        onClick={selectCallback}
        disabled={isDisabled}
      >
        Select{isSelected && 'ed'}
      </button>
    </>
  )
}
