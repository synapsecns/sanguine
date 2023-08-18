import type { NamespacesNeeded, Ni18nOptions } from 'ni18n'
import { loadTranslations as ni18nLoadTranslations } from 'ni18n'
import path from 'path'

export const ni18nConfig: Ni18nOptions = {
  supportedLngs: ['en', 'fr'],
  ns: ['default', 'other-namespace'],
}

export const loadTranslations = async (
  initialLocale?: string | undefined,
  namespacesNeeded?: NamespacesNeeded | undefined
) => {
  const locales = path.resolve('./', './public/locales')

  return await ni18nLoadTranslations(
    ni18nConfig,
    initialLocale,
    namespacesNeeded
  )
}
