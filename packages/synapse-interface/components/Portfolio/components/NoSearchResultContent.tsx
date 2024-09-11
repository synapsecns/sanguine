import { useTranslations } from 'next-intl'

export const NoSearchResultsContent = ({
  searchStr,
}: {
  searchStr: string
}) => {
  const t = useTranslations('Search')
  return (
    <div id="no-search-results-content" className="text-white">
      <p className="mb-3 break-words">
        {t('No results found for')} '{searchStr}'.
      </p>
    </div>
  )
}
