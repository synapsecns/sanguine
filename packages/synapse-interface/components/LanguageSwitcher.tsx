import { useTranslation } from 'react-i18next'

type Language = {
  code: string
  translateKey: string
}

const languages: Language[] = [
  { code: 'en', translateKey: 'english' },
  { code: 'fr', translateKey: 'french' },
]

const changeLanguage = (i18n, language) => {
  window.localStorage.setItem('MY_LANGUAGE', language)
  i18n.changeLanguage(language)
}

export default function LanguageSwitcher() {
  const { t, i18n } = useTranslation()

  return (
    <div className="">
      <select
        value={i18n.language}
        onChange={(e) => changeLanguage(i18n, e.target.value)}
      >
        {languages.map((language) => (
          <option value={language.code} key={language.code}>
            {t(language.translateKey)}
          </option>
        ))}
      </select>
    </div>
  )
}
