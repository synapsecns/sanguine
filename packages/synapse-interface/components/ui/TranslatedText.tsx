import React from 'react'
import { useTranslations } from 'next-intl'

interface TranslatedTextProps {
  namespace: string
  id: string
}

export const TranslatedText: React.FC<TranslatedTextProps> = ({
  namespace,
  id,
}) => {
  const t = useTranslations(namespace)

  return <>{t(id)}</>
}
