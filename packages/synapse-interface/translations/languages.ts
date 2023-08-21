import { MessageDescriptor } from '@lingui/core'
import { msg } from '@lingui/macro'

interface Languages {
  locale: string
  msg: MessageDescriptor
  territory?: string
  rtl: boolean
}

export type LOCALES = 'en-us' | 'fr-fr' | 'pseudo'

const languages: Languages[] = [
  {
    locale: 'en-us',
    msg: msg`English`,
    territory: 'US',
    rtl: false,
  },
  {
    locale: 'fr-fr',
    msg: msg`French`,
    territory: 'FR',
    rtl: false,
  },
]

if (process.env.NODE_ENV !== 'production') {
  languages.push({
    locale: 'pseudo',
    msg: msg`Pseudo`,
    territory: 'US',
    rtl: false,
  })
}

export default languages
