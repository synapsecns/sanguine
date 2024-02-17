import { useState } from "react"
import Image from "next/image"

import { TokenAndBalance } from "@/utils/actions/fetchPortfolioBalances"
import { HoverTokenAndBalance } from "./HoverTokenAndBalance"

export const ChainTokens = ({
  balanceTokens=[],
  hoverClassName="",
}: {
  balanceTokens?: TokenAndBalance[]
  hoverClassName?: string
}) => {
  const [isT1Hovered, setIsT1Hovered] = useState<boolean>(false)
  const [isT2Hovered, setIsT2Hovered] = useState<boolean>(false)
  const [isT3Hovered, setIsT3Hovered] = useState<boolean>(false)

  const hasNoTokens: boolean =
    !balanceTokens || (balanceTokens?.length === 0)
  const hasOneToken: boolean = balanceTokens?.length > 0
  const hasTwoTokens: boolean = balanceTokens?.length > 1
  const numOverTwoTokens: number =
    balanceTokens?.length - 2 > 0 ? balanceTokens.length - 2 : 0
  const hasOnlyOneToken: boolean = balanceTokens?.length === 1


  if (hasNoTokens) {
    return (
      <div
        data-test-id="portfolio-token-visualizer"
        className="flex flex-row items-center mr-4 cursor-pointer hover-trigger text-secondary"
      >
        -
      </div>
    )
  }
  return (
    <div
      data-test-id="portfolio-token-visualizer"
      className="flex flex-row items-center space-x-2 cursor-pointer hover-trigger"
    >
      {hasOneToken && (
        <div>
          <Image
            loading="lazy"
            className="w-6 h-6 rounded-md"
            alt={`${balanceTokens[0].token.symbol} img`}
            src={balanceTokens[0].token.icon}
            onMouseEnter={() => setIsT1Hovered(true)}
            onMouseLeave={() => setIsT1Hovered(false)}
          />
          <HoverTokenAndBalance
            isHovered={isT1Hovered}
            tokens={[balanceTokens[0]]}
            hoverClassName={hoverClassName}
          />
        </div>
      )}
      {hasOnlyOneToken && (
        <div className="text-white whitespace-nowrap">
          {balanceTokens[0].parsedBalance} {balanceTokens[0].token.symbol}
        </div>
      )}
      {hasTwoTokens && (
        <div>
          <Image
            loading="lazy"
            className="w-6 h-6 rounded-md"
            alt={`${balanceTokens[1].token.symbol} img`}
            src={balanceTokens[1].token.icon}
            onMouseEnter={() => setIsT2Hovered(true)}
            onMouseLeave={() => setIsT2Hovered(false)}
          />
          <HoverTokenAndBalance
            isHovered={isT2Hovered}
            tokens={[balanceTokens[1]]}
            hoverClassName={hoverClassName}
          />
        </div>
      )}
      {numOverTwoTokens > 0 && (
        <div
          className="text-white"
          onMouseEnter={() => setIsT3Hovered(true)}
          onMouseLeave={() => setIsT3Hovered(false)}
        >
          + {numOverTwoTokens}
        </div>
      )}
      <HoverTokenAndBalance
        isHovered={isT3Hovered}
        tokens={balanceTokens}
        startFrom={2}
        hoverClassName={`${hoverClassName} mt-3`}
      />
    </div>
  )
}