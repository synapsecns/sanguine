


import { UnsupportedChainIdError, useWeb3React } from '@web3-react/core'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { setupNetwork } from '@utils/wallet'


export function useWeb3Connect() {
  const { chainId } = useActiveWeb3React()
  const { activate, deactivate } = useWeb3React()


  function activateWallet(wallet) {
    activate(wallet.connector, async (error) => {
      if (error instanceof UnsupportedChainIdError) {
        const hasSetup = await setupNetwork(chainId)
        if (hasSetup) {
          activate(wallet.connector)
        }
        //
      } else {
        console.log(error)
        // TODO: handle error
      }
    })
  }


  return {
    activateWallet,
    deactivate
  }

}