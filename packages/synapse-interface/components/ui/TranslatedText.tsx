import { useTranslations } from 'use-intl'

export const TranslatedText = ({ namespace, id }) => {
  const t = useTranslations(namespace)

  return <>{t(id)}</>
}
