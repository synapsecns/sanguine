import { useMemo } from 'react'

import { generateTheme } from '@/utils/generateTheme'

export const useThemeVariables = (theme, customTheme) => {
  return useMemo(() => {
    if (theme === 'dark') return generateTheme({ bgColor: 'dark' })
    if (customTheme) return generateTheme(customTheme)
    return generateTheme()
  }, [theme, customTheme])
}
