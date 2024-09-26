import { useTranslations } from 'next-intl'

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
  const t = useTranslations('Activity')

  return (
    <>
      <button
        id="portfolio-asset-action-button"
        className={`
          py-1 px-6 rounded-sm
          border border-fuchsia-400
          ${!isDisabled && 'cursor-pointer hover:bg-surface active:opacity-70'}
        `}
        onClick={selectCallback}
        disabled={isDisabled}
      >
        {isSelected ? t('Selected') : t('Select')}
      </button>
    </>
  )
}
