import { useEffect } from 'react'
import { formatRGB } from '@/utils/formatRgb'
import { CustomTheme } from 'types'

export const useCustomTheme = (customTheme: CustomTheme) => {
  useEffect(() => {
    customTheme?.primary &&
      document.documentElement.style.setProperty(
        '--synapse-widget-primary-color',
        formatRGB(customTheme.primary)
      )
    customTheme?.secondary &&
      document.documentElement.style.setProperty(
        '--synapse-widget-secondary-color',
        formatRGB(customTheme.secondary)
      )
    customTheme?.accent &&
      document.documentElement.style.setProperty(
        '--synapse-widget-accent-color',
        formatRGB(customTheme.accent)
      )
    customTheme?.separator &&
      document.documentElement.style.setProperty(
        '--synapse-widget-separator-color',
        formatRGB(customTheme.separator)
      )
    customTheme?.background &&
      document.documentElement.style.setProperty(
        '--synapse-widget-background-color',
        formatRGB(customTheme.background)
      )
    customTheme?.surface &&
      document.documentElement.style.setProperty(
        '--synapse-widget-on-surface-color',
        formatRGB(customTheme.surface)
      )
  }, [customTheme])
}
