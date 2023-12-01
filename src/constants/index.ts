import { CustomTheme, CustomThemeVariables } from 'types'

export const nightTheme: CustomTheme = {
  primary: 'rgb(255,255,255)',
  secondary: 'rgb(146,150,167)',
  small: 'rgb(57,63,94)',
  separator: 'rgb(57,63,94)',
  background: 'rgb(8,9,27)',
  surface: 'rgb(22,24,46)',
  accent: 'rgb(18,106,143)',
}

export const lightThemeVariables: CustomThemeVariables = {
  '--h': '30',
  '--s': '0%',
  '--primary':    'hsl(var(--h), var(--s), 7%)',
  '--secondary':  'hsl(var(--h), var(--s), 41%)',
  '--small':      'hsl(var(--h), var(--s), 66%)',
  '--accent':     'hsl(var(--h), var(--s), 96%)',
  '--separator':  'hsl(var(--h), var(--s), 86%)',
  '--surface':    'hsl(var(--h), var(--s), 100%)',
  '--background': 'hsl(var(--h), var(--s), 96%)',
}

export const darkThemeVariables: CustomThemeVariables = {
  '--h': '30',
  '--s': '0%',
  '--primary':    'hsl(var(--h), var(--s), 96%)',
  '--secondary':  'hsl(var(--h), var(--s), 86%)',
  '--small':      'hsl(var(--h), var(--s), 66%)',
  '--accent':     'hsl(var(--h), var(--s), 29%)',
  '--separator':  'hsl(var(--h), var(--s), 13%)',
  '--surface':    'hsl(var(--h), var(--s), 13%)',
  '--background': 'hsl(var(--h), var(--s), 7%)',
}