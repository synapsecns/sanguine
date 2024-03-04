import { useState } from "react"
import Image from "next/image"

import type { TokenAndBalance } from "@/utils/actions/fetchPortfolioBalances"
import { HoverTokenAndBalance } from "./HoverTokenAndBalance"

export const ChainTokens = ({
  balanceTokens=[],
  hoverClassName="",
}: {
  balanceTokens?: TokenAndBalance[]
  hoverClassName?: string
}) => {
  const len = balanceTokens?.length
  const hasNoTokens: boolean = !balanceTokens || (len === 0)
  const hasOneToken: boolean = len > 0
  const hasTwoTokens: boolean = len > 1
  const numOverTwoTokens: number = (len > 2) ? (len - 2) : 0
  const hasOnlyOneToken: boolean = len === 1


  return (
    <div
      data-test-id="portfolio-token-visualizer"
      className="flex flex-row items-center space-x-2 cursor-pointer hover-trigger"
    >
      {hasNoTokens &&
        <span className="text-white/50"> - </span>
      }
      {hasOneToken &&
        <ChainIconAndHover
          tokenAndBalance={balanceTokens[0]}
          hoverClassName={hoverClassName}
        />
      }
      {hasOnlyOneToken && (
        <div className="text-white whitespace-nowrap">
          {balanceTokens[0].parsedBalance} {balanceTokens[0].token.symbol}
        </div>
      )}
      {hasTwoTokens &&
        <ChainIconAndHover
          tokenAndBalance={balanceTokens[1]}
          hoverClassName={hoverClassName}
        />
      }
      {numOverTwoTokens > 0 && (
        <NumOverAndHover
          balanceTokens={balanceTokens}
          numOverTwoTokens={numOverTwoTokens}
          hoverClassName={hoverClassName}
        />
      )}
    </div>
  )
}

function NumOverAndHover({
  balanceTokens,
  numOverTwoTokens,
  hoverClassName
}: {
  balanceTokens: TokenAndBalance[]
  numOverTwoTokens: number
  hoverClassName: string
}) {
  const [isHovered, setHovered] = useState<boolean>(false)

  return (
    <div
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
    >
      <div className="text-white">
        + {numOverTwoTokens}
      </div>
      <HoverTokenAndBalance
        isHovered={isHovered}
        tokens={balanceTokens}
        startFrom={2}
        hoverClassName={hoverClassName}
      />
    </div>
  )
}

function ChainIconAndHover({
  tokenAndBalance,
  hoverClassName
} : {
  tokenAndBalance: TokenAndBalance
  hoverClassName: string
}) {
  const [isHovered, setHovered] = useState<boolean>(false)

  return (
    <div
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
    >
      <Image
        loading="lazy"
        className="w-6 h-6 rounded-md"
        alt={`${tokenAndBalance.token.symbol} img`}
        src={tokenAndBalance.token.icon}
      />
      <HoverTokenAndBalance
        isHovered={isHovered}
        tokens={[tokenAndBalance]}
        hoverClassName={hoverClassName}
      />
    </div>
  )
}
