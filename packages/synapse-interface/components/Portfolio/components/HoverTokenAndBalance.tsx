import Image from 'next/image'
import type { TokenAndBalance } from '@/utils/actions/fetchPortfolioBalances'
import { HoverContent } from './HoverContent'

export function HoverTokenAndBalance({
  isHovered,
  tokens,
  className,
  hoverClassName,
  startFrom=0
}: {
  isHovered: boolean
  tokens: TokenAndBalance[]
  className?: string
  hoverClassName?: string
  startFrom?: number
}) {
  return (
    <div className={`relative ${className}`}>
      <HoverContent isHovered={isHovered} className={hoverClassName}>
        {tokens?.map((token: TokenAndBalance, key: number) => {
          if (key >= startFrom) {
            return (
              <TokenImageAndBalance
                tokenAndBalance={token}
                key={key}
              />
            )
          }
        })}
      </HoverContent>
    </div>
  )
}

function TokenImageAndBalance({
  tokenAndBalance
}: {
  tokenAndBalance: TokenAndBalance
}) {
  return (
    <div className="whitespace-nowrap">
      <ImageForToken
        token={tokenAndBalance.token}
        className="!size-5 inline -mt-1 mr-2 -ml-0.5"
      />
      {tokenAndBalance?.parsedBalance}{' '}
      {tokenAndBalance?.token.symbol}
    </div>
  )
}

function ImageForToken({ token, className, ...props }) {
  return (
    <Image
      loading="lazy"
      className={`size-6 rounded-md ${className}`}
      alt={`${token.symbol} img`}
      src={token.icon}
      {...props}
    />
  )
}