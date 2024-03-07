import './index.css'
;('use client')
import { BridgeProps } from 'types'
import { Provider } from 'react-redux'
import { Web3Provider } from 'providers/Web3Provider'
import { SynapseProvider } from 'providers/SynapseProvider'

import { Widget } from '@/components/Widget'
import { store } from '@/state/store'
import * as BRIDGEABLE from '@/constants/bridgeable'
import { CHAINS_ARRAY } from '@/constants/chains'
import { BackgroundListenerProvider } from '@/providers/BackgroundListenerProvider'
import { useBridgeSelectionData } from '@/hooks/useBridgeSelectionData'

export const Bridge = ({
  web3Provider,
  customRpcs,
  customTheme,
  container,
  targetChainIds,
  targetTokens,
  protocolName,
}: BridgeProps) => {
  if (!web3Provider) {
    return null
  }

  return (
    <Web3Provider config={web3Provider}>
      <SynapseProvider chains={CHAINS_ARRAY} customRpcs={customRpcs}>
        <Provider store={store}>
          <BackgroundListenerProvider>
            <Widget
              customTheme={customTheme}
              container={container}
              targetChainIds={targetChainIds}
              targetTokens={targetTokens}
              protocolName={protocolName}
            />
          </BackgroundListenerProvider>
        </Provider>
      </SynapseProvider>
    </Web3Provider>
  )
}

export const useBridgeSelections = () => useBridgeSelectionData()

export const AGEUR = BRIDGEABLE.AGEUR
export const AVAX = BRIDGEABLE.AVAX
export const BTCB = BRIDGEABLE.BTCB
export const BUSD = BRIDGEABLE.BUSD
export const CRVUSD = BRIDGEABLE.CRVUSD
export const DAI = BRIDGEABLE.DAI
export const DAIe = BRIDGEABLE.DAIe
export const DOG = BRIDGEABLE.DOG
export const ETH = BRIDGEABLE.ETH
export const FRAX = BRIDGEABLE.FRAX
export const FTM = BRIDGEABLE.FTM
export const GMX = BRIDGEABLE.GMX
export const GOHM = BRIDGEABLE.GOHM
export const H2O = BRIDGEABLE.H2O
export const HIGH = BRIDGEABLE.HIGH
export const JEWEL = BRIDGEABLE.JEWEL
export const JUMP = BRIDGEABLE.JUMP
export const KLAY = BRIDGEABLE.KLAY
export const L2DAO = BRIDGEABLE.L2DAO
export const LINK = BRIDGEABLE.LINK
export const LUSD = BRIDGEABLE.LUSD
export const MATIC = BRIDGEABLE.MATIC
export const METISUSDC = BRIDGEABLE.METISUSDC
export const MOVR = BRIDGEABLE.MOVR
export const NETH = BRIDGEABLE.NETH
export const NEWO = BRIDGEABLE.NEWO
export const NFD = BRIDGEABLE.NFD
export const NOTE = BRIDGEABLE.NOTE
export const NUSD = BRIDGEABLE.NUSD
export const ONEDAI = BRIDGEABLE.ONEDAI
export const ONEETH = BRIDGEABLE.ONEETH
export const ONEUSDC = BRIDGEABLE.ONEUSDC
export const ONEUSDT = BRIDGEABLE.ONEUSDT
export const PEPE = BRIDGEABLE.PEPE
export const PLS = BRIDGEABLE.PLS
export const SDT = BRIDGEABLE.SDT
export const SFI = BRIDGEABLE.SFI
export const SOLAR = BRIDGEABLE.SOLAR
export const SUSD = BRIDGEABLE.SUSD
export const SYN = BRIDGEABLE.SYN
export const SYNFRAX = BRIDGEABLE.SYNFRAX
export const SYNJEWEL = BRIDGEABLE.SYNJEWEL
export const UNIDX = BRIDGEABLE.UNIDX
export const USDBC = BRIDGEABLE.USDBC
export const USDC = BRIDGEABLE.USDC
export const USDCe = BRIDGEABLE.USDCe
export const USDT = BRIDGEABLE.USDT
export const USDTe = BRIDGEABLE.USDTe
export const VSTA = BRIDGEABLE.VSTA
export const WAVAX = BRIDGEABLE.WAVAX
export const WBTC = BRIDGEABLE.WBTC
export const WETH = BRIDGEABLE.WETH
export const WETHE = BRIDGEABLE.WETHE
export const WFTM = BRIDGEABLE.WFTM
export const WJEWEL = BRIDGEABLE.WJEWEL
export const WKLAY = BRIDGEABLE.WKLAY
export const WMATIC = BRIDGEABLE.WMATIC
export const WMOVR = BRIDGEABLE.WMOVR
export const WSOHM = BRIDGEABLE.WSOHM
export const XJEWEL = BRIDGEABLE.XJEWEL
