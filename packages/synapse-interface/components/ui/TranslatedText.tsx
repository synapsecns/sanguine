import { useTranslations } from 'use-intl'

export const TranslatedText = ({ key, text }) => {
  const t = useTranslations(key)

  return <>{t(text)}</>
}
