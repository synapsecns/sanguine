import { CustomThemeVariables } from 'types'

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