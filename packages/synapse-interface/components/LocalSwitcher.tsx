import { useLingui } from '@lingui/react'
import { useRouter } from 'next/router'
import { useState } from 'react'
import ReactCountryFlag from 'react-country-flag'
import languages, { LOCALES } from '../translations/languages'

export const LocaleSwitcher = () => {
  const { i18n } = useLingui()
  const router = useRouter()

  const [locale, setLocale] = useState<LOCALES>(
    router.locale!.split('-')[0] as LOCALES
  )

  function handleChange(event: React.ChangeEvent<HTMLSelectElement>) {
    const selectedLocale = event.target.value as LOCALES
    setLocale(selectedLocale)

    const href = {
      pathname: router.pathname,
      query: router.query,
    }

    router.push(href, href, { locale: selectedLocale }).catch((e) => {
      console.log(`Error: ${e}`)
    })
  }

  return (
    <div className="space-x-2">
      <select onChange={handleChange} value={locale}>
        {languages.map((language) => {
          return (
            <option key={language.locale} value={language.locale}>
              {i18n._(language.msg)}{' '}
              {language.territory && `(${language.territory})`}
            </option>
          )
        })}
      </select>
    </div>
  )
}
