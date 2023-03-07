
import { Zero } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'

import { ChainId } from '@constants/networks'

import { useLCDClient } from '@terra-money/wallet-provider'
import { useEffect, useState } from 'react'
import { useNetworkController } from '@hooks/wallet/useNetworkController'
import { usePoller } from '@hooks/usePoller'


const TERRA_MIN_BALANCE_NUM = BigNumber.from("500000")

export function useTerraUstBalance() {
  const { activeChainId, terraAddress } = useNetworkController()
  // const { terraAddress, ...stuff } = useTerraWallet()
  const lcd = useLCDClient()

  const [balance, setBalance] = useState(Zero)

  function getTerraUstBalance() {
    if (activeChainId == ChainId.TERRA) {
      lcd.bank.balance(terraAddress).then(r => {
        const { amount } = r[0]._coins.uusd.toData()
        if (amount && (amount !== "")) {
          let newAmt = BigNumber.from(amount)
          if (newAmt.gte(TERRA_MIN_BALANCE_NUM)) {
            newAmt = newAmt.sub(TERRA_MIN_BALANCE_NUM)
          }
          if (newAmt != balance ) {
            setBalance(newAmt)
          }

        } else {
          setBalance(Zero)
        }
      })
    }
  }

  usePoller(
    getTerraUstBalance,
    10000
  )

  useEffect(
    getTerraUstBalance
    ,
    [terraAddress]
  )


  return balance

}


