import { type Chain, type Token, type ActionTypes } from '@/utils/types'

export interface BridgeCardTypes {
  bridgeRef: React.RefObject<HTMLDivElement>
  children: React.ReactNode
}

export interface SelectorTypes {
  dataTestId?: string
  isOrigin: boolean
  label?: string
  placeholder?: string
  selectedItem: Token | Chain
  itemListFunction?: Function
  setFunction?: Function
}

export interface TokenSelectorTypes extends SelectorTypes {
  selectedItem: Token
  action: ActionTypes
  disabled: boolean
}

export interface ChainSelectorTypes extends SelectorTypes {
  selectedItem: Chain
  action: ActionTypes
  disabled: boolean
}
