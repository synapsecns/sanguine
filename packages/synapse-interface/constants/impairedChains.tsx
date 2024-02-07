import * as CHAINS from '@constants/chains/master'

/*
Create popups here, and use them in the IMPAIRED_CHAINS constant below.
*/
const GENERIC_CHAIN_PAUSED = () => {
  return <>This Chain is Paused</>
}

const HARMONY_WARNING = () => {
  return (
    <>
      The native Harmony bridge has been exploited, which lead to a hard depeg
      of the following Harmony-specific tokens: 1DAI, 1USDC, 1USDT, 1ETH.
      <br /> Please see the{' '}
      <a
        className="underline"
        href="https://twitter.com/harmonyprotocol/status/1540110924400324608"
      >
        official Harmony Twitter
      </a>{' '}
      for status updates and exercise caution when interacting with Harmony.
    </>
  )
}

/*
IMPAIRED_CHAINS is a global constant that is used to determine whether or not to show a popup on the bridge page.
Setting the disabled property to true will disable the swap/bridge button for that chain.
*/

export const IMPAIRED_CHAINS = {
  [CHAINS.HARMONY.id]: { content: HARMONY_WARNING, disabled: false },
}
