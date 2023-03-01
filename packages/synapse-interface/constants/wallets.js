import metamaskIcon from '@assets/icons/metamask.svg'
import binanceIcon from '@assets/icons/binance.svg'
import walletconnectIcon from '@assets/icons/walletconnect.svg'
import coinbaseIcon from '@assets/icons/coinbase.svg'
import terraIcon from '@assets/networks/terra.png'

import { injected, bsc, walletconnect, walletlink } from '@connectors'

export const WALLETS = [
  {
    id: 'metamask',
    name: 'MetaMask',
    connector: injected,
    icon: metamaskIcon,
  },
  {
    id: 'walletconnect',
    name: 'Wallet Connect',
    connector: walletconnect,
    icon: walletconnectIcon,
  },
  {
    id: 'binancewallet',
    name: 'Binance Wallet',
    connector: bsc,
    icon: binanceIcon,
  },
  {
    id: 'coinbasewallet',
    name: 'Coinbase Wallet',
    connector: walletlink,
    icon: coinbaseIcon,
  },
  {
    id: 'terrastation',
    name: 'Terra Station',
    icon: terraIcon,
  },
]
