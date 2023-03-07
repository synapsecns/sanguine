import _ from 'lodash'
import Grid from '@tw/Grid'
import { LandingPageWrapper } from '@components/layouts/LandingPageWrapper'
import { useRouter } from 'next/router'
import { useRef, useState, useEffect } from 'react'

// import BridgeCard from './BridgeCard'
// import BridgeWatcher from './BridgeWatcher'
import { ActionCardFooter } from '@components/ActionCardFooter'
import { useAccount, useBalance, useNetwork } from 'wagmi'
import { HOW_TO_BRIDGE_URL } from '@constants/urls'
import { ChainId } from '@constants/networks'

const sanitizeQueryData = (data: string) => {}

export default function BridgePage() {
  // Get data from query url.
  const router = useRouter()
  const { outputChain, inputCurrency, outputCurrency } = router.query
  const toChainQuery = String(outputChain)
  const fromQuery = String(inputCurrency)
  const toQuery = String(outputCurrency)

  // Get data from wagmi.
  const { address } = useAccount()
  const { chain } = useNetwork()

  // Set data from query url.

  // Initialize state variables needed for bridging
  const [fromChainId, setFromChainId] = useState(ChainId.ETH)
  const [toChainId, setToChainId] = useState(ChainId.ARBITRUM)
  const [fromCoin, setFromCoin] = useState('USDC')
  const [toCoin, setToCoin] = useState<string>(toQuery)

  return (
    <LandingPageWrapper>
      <main className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none">
        <div className="items-center px-4 py-8 mx-auto mt-4 2xl:w-3/4 sm:mt-6 sm:px-8 md:px-12">
          <div>
            <Grid
              cols={{ xs: 1 }}
              gap={6}
              className="justify-center px-2 py-16 sm:px-6 md:px-8"
            >
              {/* <HarmonyCheck fromChainId={fromChainId} toChainId={toChainId} /> */}
              <div className="flex justify-center">
                <div className="pb-3 place-self-center">
                  {/* <BridgeCard
                  {...{
                    fromChainTokens,
                    toChainTokens,
                    fromChainId,
                    toChainId,
                    fromCoin,
                    fromValue,
                    toCoin,
                    toValue,
                    onSelectFromCoin,
                    onSelectToCoin,
                    onSelectFromChain,
                    onSelectToChain,
                    swapFromToCoins,
                    swapFromToChains,
                    onChangeFromAmount,
                    onChangeToAmount,
                    error,
                    priceImpact,
                    exchangeRate,
                    feeAmount,
                    fromRef,
                    toRef,
                    destinationAddress,
                    setDestinationAddress,
                  }}
                /> */}
                  <ActionCardFooter link={HOW_TO_BRIDGE_URL} />
                </div>
              </div>
              <div>{/* <BridgeWatcher /> */}</div>
            </Grid>
          </div>
        </div>
      </main>
    </LandingPageWrapper>
  )
}

// export function HarmonyCheck({ fromChainId, toChainId }: { fromChainId: number; toChainId: number }) {
//   return (
//     <>
//       {(toChainId === ChainId.HARMONY || fromChainId === ChainId.HARMONY) && (
//         <div
//           className={`bg-gray-800 shadow-lg pt-3 px-6 pb-6 rounded-lg text-white`}
//         >
//           The native Harmony bridge has been exploited, which lead to a hard depeg of the following Harmony-specific tokens: 1DAI, 1USDC, 1USDT, 1ETH.
//           <br /> Please see the{' '}
//           <a
//             className="underline"
//             href="https://twitter.com/harmonyprotocol/status/1540110924400324608"
//           >
//             official Harmony Twitter
//           </a>{' '}
//           for status updates and exercise caution when interacting with Harmony.
//         </div>
//       )}
//     </>
//   )
// }
