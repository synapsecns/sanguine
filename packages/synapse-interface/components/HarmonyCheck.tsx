import * as CHAINS from '@/constants/chains/master'

const HarmonyCheck = ({
  fromChainId,
  toChainId,
}: {
  fromChainId: number
  toChainId: number
}) => {
  return (
    <>
      {(toChainId === CHAINS.HARMONY.id ||
        fromChainId === CHAINS.HARMONY.id) && (
        <div
          className={`bg-gray-800 shadow-lg pt-3 px-6 pb-6 rounded-md text-white`}
        >
          The native Harmony bridge has been exploited, which lead to a hard
          depeg of the following Harmony-specific tokens: 1DAI, 1USDC, 1USDT,
          1ETH.
          <br /> Please see the{' '}
          <a
            className="underline"
            href="https://twitter.com/harmonyprotocol/status/1540110924400324608"
          >
            official Harmony Twitter
          </a>{' '}
          for status updates and exercise caution when interacting with Harmony.
        </div>
      )}
    </>
  )
}
export default HarmonyCheck
