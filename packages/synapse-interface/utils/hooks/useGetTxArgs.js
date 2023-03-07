// import { parseUnits } from '@ethersproject/units'
// import { getGasPrice } from '@utils/gas'
// import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { checkCleanedValue } from "@utils/checkCleanedValue"
/**
 * gets the common tx args requiref for submitting txns including
 * things like slippage, deadline, gas price etc...
 */
export function useGetTxArgs() {
  // const { chainId } = useActiveWeb3React()

  return function getTxArgs(opts) {
    const { deadlineMinutes } = opts ?? {}

    // const { gasPriceSelected } = opts ?? {}
    // let gasPrice = await getGasPrice(gasPriceSelected, chainId)
    // gasPrice = parseUnits(String(gasPrice) || '45', 9)
    const numDeadlineMinutes = Number(deadlineMinutes)

    let dm
    if (!checkCleanedValue(numDeadlineMinutes)) {
      dm = numDeadlineMinutes
    }

    return {
      slippageCustom:            null,
      slippageSelected:          'ONE_TENTH',
      infiniteApproval:          true,
      transactionDeadline:       getTimeMinutesFromNow(dm ?? 10),                     // the time on samechain when a swap expires
      bridgeTransactionDeadline: getTimeMinutesFromNow(dm ? (dm + 60) : (60 * 24)),   // the time on crosschain when a swap expires
      // gasPrice:                  gasPrice
    }
  }
}


function getTimeMinutesFromNow(minutesFromNow) {
  const currentTimeSeconds = new Date().getTime() / 1000

  return Math.round(
    currentTimeSeconds + 60 * minutesFromNow
  )
}