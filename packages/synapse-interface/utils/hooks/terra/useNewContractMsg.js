import { MsgExecuteContract, Coins, Coin } from "@terra-money/terra.js"

import { useAddress } from "@hooks/terra/useAddress"

export function useNewContractMessage() {
  const sender = useAddress()

  /**
   * @param {string} contract
   * @param {object} msg
   * @param {Coin?} coin { denom: string; amount: string }
   */
  function newContractMsgFunc(
    contract,
    msg,
    coin
  ) {
    return new MsgExecuteContract(
      sender,
      contract,
      msg,
      new Coins(coin ? [Coin.fromData(coin)] : [])
    )
  }

  return newContractMsgFunc
}