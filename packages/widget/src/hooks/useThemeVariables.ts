import { useMemo } from 'react'

import { generateTheme } from '@/utils/generateTheme'

export const useThemeVariables = (customTheme) => {
  return useMemo(() => {
    if (customTheme) return generateTheme(customTheme)
    return generateTheme()
  }, [customTheme])
}
