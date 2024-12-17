import { CustomThemeVariables } from 'types'

export const lightThemeVariables: CustomThemeVariables = {
  '--synapse-text': 'hsl(240deg 0% 7%)',
  '--synapse-secondary': 'hsl(240deg 0% 41%)',
  '--synapse-focus': 'hsl(240deg 0% 66%)',
  '--synapse-select-bg': 'hsl(240deg 0% 96%)',
  '--synapse-surface': 'hsl(240deg 0% 100%)',
  '--synapse-root': 'hsl(240deg 0% 96%)',
  '--synapse-border': 'hsl(240deg 0% 86%)',
  '--synapse-accent': 'hsl(330, 100%, 45%, 1)',
}

export const darkThemeVariables: CustomThemeVariables = {
  '--synapse-text': 'hsl(240deg 0% 96%)',
  '--synapse-secondary': 'hsl(240deg 0% 86%)',
  '--synapse-focus': 'hsl(240deg 0% 66%)',
  '--synapse-select-bg': 'hsl(240deg 0% 29%)',
  '--synapse-surface': 'hsl(240deg 0% 13%)',
  '--synapse-root': 'hsl(240deg 0% 7%)',
  '--synapse-border': 'hsl(240deg 0% 13%)',
  '--synapse-accent': 'hsl(330, 100%, 45%, 1)',
}

export const MAX_UINT256 =
  115792089237316195423570985008687907853269984665640564039457584007913129639935n

export const TRANSACTION_SUPPORT_URL =
  'https://docs.synapseprotocol.com/docs/Support/Transaction-Support'

export const PAUSED_CHAINS_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/master/packages/synapse-interface/public/pauses/v1/paused-chains.json'
export const PAUSED_MODULES_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json'
