import { createContext, useContext, useEffect, useMemo, useState } from 'react'


import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useChainSwitcher } from '@hooks/wallet/useChainSwitcher'
import { useTerraWallet } from '@hooks/terra/useTerraWallet'

import { ChainId } from '@constants/networks'
import { useWeb3Connect } from '@hooks/wallet/useWeb3Connect'
import { WALLETS } from '@constants/wallets'
import { WalletModalContext } from './WalletModalStore'



export function NetworkStore({ children }) {
  const { chainId, account } = useActiveWeb3React()
  const preTriggerChainSwitch = useChainSwitcher()
  const { activateWallet, deactivate } = useWeb3Connect()
  const [showWalletModal, setShowWalletModal] = useContext(WalletModalContext)
  const { terraAddress: tempTerraAddress, disconnect } = useTerraWallet()
  const terraAddress = useMemo(() => tempTerraAddress, [tempTerraAddress])


  const [activeChainId, setActiveChainId] = useState(chainId)

  const nonEvmChainId = useMemo(
    () => {
      if (terraAddress) {
        return ChainId.TERRA
      } else {
        return undefined
      }
    },
    [terraAddress]
  )

  function connectToChain(cid) {
    if (account) {
      // preTriggerChainSwitch(cid)
      return preTriggerChainSwitch(cid)
    } else {
      setShowWalletModal(true)
      return new Promise() // prev had this line commented
    }
  }


  function triggerChainSwitch(newChainId) {
    // if (chainId != ChainId.TERRA) {
    //   connectToChain(newChainId)
    // }
    return (
      connectToChain(newChainId) // preTriggerChainSwitch(newChainId)
        .then(() => {
          setActiveChainId(newChainId)
        })
        .catch(e => {})
    )
  }

  function disconnectChain(newChainId) {
    if ((newChainId == ChainId.TERRA) && (activeChainId == ChainId.TERRA))  {
      disconnect()
      setActiveChainId(chainId)
    } else if (newChainId == ChainId.TERRA) {
      disconnect()
    } else {
      deactivate()
    }
  }

  useEffect(
    () => {
      let tempFromChainId
      if (terraAddress && !account) {
        tempFromChainId = ChainId.TERRA
      } else if (!terraAddress && account) {
        tempFromChainId = chainId
      } else {
        tempFromChainId = activeChainId ?? chainId
      }
      setActiveChainId(tempFromChainId)

    },
    [terraAddress, account]
  )


  return (
    <NetworkContext.Provider
      value={{
        activeChainId,
        evmChain: chainId,
        nonEvmChainId,
        account,
        terraAddress,
        triggerChainSwitch,
        setActiveChainId,
        disconnectChain,
        connectToChain//: preTriggerChainSwitch

      }}
    >
      {children}
    </NetworkContext.Provider>
  )
}

export const NetworkContext = createContext({})



