import { ChevronDownIcon } from '@heroicons/react/solid'

import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'

export const TokenSelector = ({ token, label, onClick, ...props }) => {
  return (
    <button
      {...props}
      className={`
        group
        p-md rounded-sm min-w-[80px]
        bg-slate-400/10 hover:bg-slate-400/20
        border border-white/10
        ${getMenuItemHoverBgForCoin(token?.color)}
        ${getBorderStyleForCoinHover(token?.color)}
      `}
      onClick={onClick}
    >
      <div className="flex items-center space-x-3">
        { token &&
          <div className="flex-none hidden md:inline-block">
            <img src={token?.icon?.src} alt={token?.symbol} className="w-6 h-6" />
          </div>
        }
        <div className="text-left">
          <div className="text-lg text-primaryTextColor">
            {token?.symbol ?? label}
          </div>
        </div>
        <ChevronDownIcon
          className={`
              flex-none
              rotate-0
              text-white/30 group-hover:text-white/80 w-4 h-4
              active:text-white/80
              group-active:rotate-180 transition-all
          `}
        />
      </div>
    </button>
  )
}
