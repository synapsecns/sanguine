import Image from 'next/image'
import { TokenAndBalance } from '@/utils/actions/fetchPortfolioBalances'


export function HoverTokenAndBalance({
  isHovered,
  tokens,
  hoverClassName,
  startFrom=0
}: {
  isHovered: boolean
  tokens: TokenAndBalance[]
  hoverClassName?: string
  startFrom?: number
}) {
  return (
    <div className="relative">
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
export const HoverContent = ({
  isHovered,
  children,
  className
}: {
  isHovered: boolean
  children: React.ReactNode
  className?: string
}) => {
  if (isHovered) {
    return (
     <div className="absolute !overflow-visible ">
        <div
            className={`
            absolute z-[100] hover-content py-2 px-3 text-white
            border border-white/20 bg-bgBase/10 backdrop-blur-xl
            rounded-md text-left -translate-x-[69%]
            ${className}
            `}
        >
        <div className="pr-6">
            {children}
        </div>

      </div>
     </div>

    )
  }
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