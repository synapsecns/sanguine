import { HoverTooltip } from '@/components/HoverTooltip'

interface GasInfoBadgeProps {
  amount: string
  symbol: string
  tooltipText: string
  prefix?: string
}

export const GasInfoBadge = ({
  amount,
  symbol,
  tooltipText,
  prefix = '+',
}: GasInfoBadgeProps) => {
  if (!amount) return null

  return (
    <HoverTooltip hoverContent={tooltipText}>
      <div className="text-xs text-zinc-500 dark:text-zinc-400 whitespace-nowrap">
        {prefix}{amount} {symbol}
      </div>
    </HoverTooltip>
  )
}
