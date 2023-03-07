import { usePoller } from "@hooks/usePoller"
import { useCallback, useEffect, useState, useMemo } from "react"
import { useGetTerraBridgeEvent } from "./useGetTerraBridgeEvent"
import { useTerraWallet } from "./useTerraWallet"

export function useTerraKappaCheck({ kekTxSig, isTerra }) {
  const { terraAddress } = useTerraWallet()
  const [kappaExists, setKappaExists] = useState(false)
  const getTerraBridgeEvent = useGetTerraBridgeEvent()


  // const terraAddress = useMemo(() => rawTerraAddr, [rawTerraAddr])
  // console.log({ isTerra })

  const getTerraTransactions = useCallback(_.throttle(() => {
    if (isTerra && !kappaExists) {
      getTerraBridgeEvent({ kekTxSig })
        .then( result => {
          if (result?.kappa) {
            // console.log({kappa})
            setKappaExists(true)
          }
        })
    }
  }, 10000, { leading: true })
  )


  useEffect(
    getTerraTransactions,
    [
      terraAddress
    ]
  )

  usePoller(getTerraTransactions, 16900)

  return kappaExists

}