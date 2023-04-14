import { useAccount } from 'wagmi'
import Grid from '@tw/Grid'
import SwapCard from './SwapCard'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import { PageHeader } from '@components/PageHeader'
const SwapPage = () => {
  const { address } = useAccount()

  // const tokenIndexFrom = poolTokens.findIndex((i) =>
  //   matchSymbolWithinPool(i, fromCoin)
  // )
  // const tokenIndexTo = poolTokens.findIndex((i) =>
  //   matchSymbolWithinPool(i, toCoin)
  // )

  // function triggerRateAndImpact({ amountToGive, amountToReceive }) {
  //   setExchangeRate(
  //     calculateExchangeRate(
  //       amountToGive,
  //       fromCoin.decimals[chainId],
  //       amountToReceive,
  //       toCoin.decimals[chainId]
  //     )
  //   )
  // }

  // function checkIfBalanceSufficient({ amountToGive }) {
  //   if (amountToGive.gt(fromBalance)) {
  //     setError('Insufficent Balance')
  //   } else {
  //     setError(null)
  //   }
  // }

  // async function calculateSwapAmount() {
  //   if (swapContract == null) return

  //   let cleanedFromValue = sanitizeValue(fromValue)
  //   if (checkCleanedValue(cleanedFromValue)) {
  //     setToValue('')
  //     return
  //   }

  //   const amountToGive = parseUnits(
  //     cleanedFromValue,
  //     fromCoin.decimals[chainId]
  //   )

  //   checkIfBalanceSufficient({ amountToGive })

  //   const amountToReceive = await calcAmountToRecieve({
  //     swapContract,
  //     tokenIndexFrom,
  //     tokenIndexTo,
  //     amountToGive,
  //   })

  //   if (sanitizeValue(fromRef.current.value) == sanitizeValue(fromValue)) {
  //     setToValue(formatUnits(amountToReceive, toCoin.decimals[chainId]))
  //     triggerRateAndImpact({ amountToGive, amountToReceive })
  //   }
  // }

  // async function calculateInverseSwapAmount() {
  //   if (swapContract == null) return
  //   const cleanedToValue = sanitizeValue(toValue)
  //   if (checkCleanedValue(cleanedToValue)) {
  //     setFromValue('')
  //     return
  //   }

  //   const amountToReceive =
  //     parseUnits(cleanedToValue, toCoin.decimals[chainId]) ?? One

  //   const amountToGive = await estimateAmountToGive({
  //     targetAmountToRecieve: amountToReceive,
  //     swapContract,
  //     tokenIndexFrom,
  //     tokenIndexTo,
  //     fromCoin,
  //     toCoin,
  //     chainId,
  //   })

  //   checkIfBalanceSufficient({ amountToGive })
  //   if (sanitizeValue(toRef.current.value) == sanitizeValue(toValue)) {
  //     setFromValue(formatUnits(amountToGive, fromCoin.decimals[chainId]))
  //     triggerRateAndImpact({ amountToGive, amountToReceive })
  //   }
  // }

  return (
    <LandingPageWrapper>
      <div>
        <Grid
          cols={{ xs: 1 }}
          gap={6}
          className="justify-center px-2 py-16 sm:px-6 md:px-8"
        >
          <div className="pb-3 place-self-center">
            <div className="flex justify-between mb-5 ml-5 mr-5">
              <PageHeader
                title="Swap"
                subtitle="Exchange stablecoins on-chain."
              />
            </div>
            <SwapCard address={address} />
          </div>
        </Grid>
      </div>
    </LandingPageWrapper>
  )
}

export default SwapPage
