import { useTranslations } from 'next-intl'

export const SearchResults = ({ searchStr }: { searchStr: string }) => {
  const t = useTranslations('Search')

  return (
    <div>
      {searchStr && (
        <div className="p-2 text-sm">
          {t('No other results found for')} <q>{searchStr}</q>.
          <div className="pt-2 align-bottom text-primaryTextColor text-md">
            {t('Want to see it supported on Synapse? Let us know!')}
          </div>
        </div>
      )}
    </div>
  )
}
