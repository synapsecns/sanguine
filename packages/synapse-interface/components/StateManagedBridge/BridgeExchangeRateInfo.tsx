import { useState, useMemo, useEffect } from 'react'
import { formatBigIntToPercentString } from '@/utils/bigint/format'
import { CHAINS_BY_ID } from '@constants/chains'
import * as CHAINS from '@constants/chains/master'
import { useCoingeckoPrice } from '@hooks/useCoingeckoPrice'
import { useGasDropAmount } from '@/utils/hooks/useGasDropAmount'
import Image from 'next/image'
import { formatBigIntToString } from '@/utils/bigint/format'
import { Token } from '@/utils/types'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '../../store/store'

const BridgeExchangeRateInfo = ({ showGasDrop }: { showGasDrop: boolean }) => {
  const [gasDropChainId, setGasDropChainId] = useState<number>(null)

  const fromAmount = useSelector((state: RootState) => state.bridge.fromValue)
  const toToken = useSelector((state: RootState) => state.bridge.toToken)
  const exchangeRate = useSelector(
    (state: RootState) => state.bridge.bridgeQuote.exchangeRate
  )
  const toChainId = useSelector((state: RootState) => state.bridge.toChainId)
  // TODO: this is ugly, refactor
  const bridgeModuleName = useSelector(
    (state: RootState) => state.bridge.bridgeQuote.bridgeModuleName
  )
  let { gasDrop: gasDropAmount, loading } = useGasDropAmount(toChainId)
  if (bridgeModuleName === 'SynapseRFQ') {
    gasDropAmount = 0n
    loading = false
  }

  const safeExchangeRate = typeof exchangeRate === 'bigint' ? exchangeRate : 0n
  const safeFromAmount = fromAmount ?? '0'

  const formattedExchangeRate = formatBigIntToString(safeExchangeRate, 18, 4)
  const numExchangeRate = Number(formattedExchangeRate)
  const slippage = safeExchangeRate - 1000000000000000000n
  const formattedPercentSlippage = formatBigIntToPercentString(slippage, 18)
  const underFee = safeExchangeRate === 0n && safeFromAmount != '0'

  console.log(numExchangeRate)

  const textColor: string = useMemo(() => {
    if (numExchangeRate >= 1) {
      return 'text-green-500'
    } else if (numExchangeRate > 0.975) {
      return 'text-amber-500'
    } else {
      return 'text-red-500'
    }
  }, [numExchangeRate])

  const isGasDropped = useMemo(() => {
    if (gasDropAmount) {
      return gasDropAmount.gt(0)
    }
  }, [gasDropAmount])

  useEffect(() => {
    setGasDropChainId(toChainId)
  }, [toChainId, isGasDropped])

  const memoizedGasDropLabel = useMemo(() => {
    if (toChainId === CHAINS.ETH.id) return null
    if (!isGasDropped || !(toChainId == gasDropChainId)) return null
    if (loading) return null
    return <GasDropLabel gasDropAmount={gasDropAmount} toChainId={toChainId} />
  }, [toChainId, gasDropChainId, isGasDropped, loading])

  const expectedToChain = useMemo(() => {
    return toChainId && <ChainInfoLabel chainId={toChainId} />
  }, [toChainId])

  return (
    <div className="flex flex-col gap-1 px-1 py-2">
      <div>
        {showGasDrop && memoizedGasDropLabel}
      </div>
      <div className="flex justify-between">
        <div className="flex gap-2 items-center">
          Expected price on {expectedToChain}
        </div>
        <div>
          {safeFromAmount != '0'
            ? `${formattedExchangeRate} ${toToken?.symbol}`
            : '−'
          }
        </div>
      </div>
      <div className="flex justify-between">
        <div>
          Slippage
        </div>
        <div className={safeFromAmount == '0' && textColor}>
          {safeFromAmount != '0' && !underFee ? formattedPercentSlippage : '−'}
        </div>
      </div>
    </div>
  )
}

const GasDropLabel = ({
  gasDropAmount,
  toChainId,
}: {
  gasDropAmount: bigint
  toChainId: number
}) => {
  let decimalsToDisplay
  const symbol = CHAINS_BY_ID[toChainId]?.nativeCurrency.symbol

  if ([CHAINS.FANTOM.id].includes(toChainId)) {
    decimalsToDisplay = 2
  } else if (
    [CHAINS.BNB.id, CHAINS.AVALANCHE.id, CHAINS.BOBA.id].includes(toChainId)
  ) {
    decimalsToDisplay = 3
  } else {
    decimalsToDisplay = 4
  }

  const formattedGasDropAmount = formatBigIntToString(
    gasDropAmount,
    18,
    decimalsToDisplay
  )

  const airdropInDollars = getAirdropInDollars(symbol, formattedGasDropAmount)

  return (
    <>
      Will also receive {formattedGasDropAmount} {symbol} {airdropInDollars && `($${airdropInDollars})`}
    </>
  )
}

const ChainInfoLabel = ({ chainId }: { chainId: number }) => {
  const chain = CHAINS_BY_ID[chainId]
  return chain ? (
    <span className="flex items-center gap-1">
      <Image
        alt="chain image"
        src={chain?.chainImg}
        className="w-4 h-4 rounded-full inline"
      />
      {chain?.name?.length > 10 ? chain?.chainSymbol : chain?.name}
    </span>
  ) : null
}

const getAirdropInDollars = (
  symbol: string,
  formattedGasDropAmount: string
) => {
  const price = useCoingeckoPrice(symbol)

  if (price) {
    const airdropInDollars = parseFloat(formattedGasDropAmount) * price

    return airdropInDollars.toFixed(2)
  } else {
    return undefined
  }
}
export default BridgeExchangeRateInfo
