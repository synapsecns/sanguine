import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'

import { ChevronDownIcon } from '@heroicons/react/solid'

import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'

export const TokenSelector = ({ token, label, onClick, ...props }) => {

  let buttonContent

  if (token) {
    const src = token?.icon?.src
    const symbol = token?.symbol

    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="flex-none hidden md:inline-block">
          <img src={src} alt={symbol} className="w-6 h-6" />
        </div>
        <div className="text-left">
          <div className="text-lg text-primaryTextColor">{symbol}</div>
        </div>
        <DropDownArrowSvg className="flex-none" />
      </div>
    )
  } else {
    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="text-left">
          <div className="text-lg text-primaryTextColor">{label}</div>
        </div>
        <DropDownArrowSvg className="flex-none" />
      </div>
    )
  }

  return (
    <button
      {...props}
      className={`
        p-md rounded-sm min-w-[80px]
        bg-slate-400/10 hover:bg-slate-400/20
        border border-white/10
        ${getMenuItemHoverBgForCoin(token?.color)}
        ${getBorderStyleForCoinHover(token?.color)}
      `}
      onClick={onClick}
    >
      {buttonContent}
    </button>
  )
}
