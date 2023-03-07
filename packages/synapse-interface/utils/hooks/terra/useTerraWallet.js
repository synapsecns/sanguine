import { ConnectType, useWallet } from '@terra-money/wallet-provider'
import { useMemo } from 'react'



export function useTerraWallet() {

  const { connect, wallets, ...rest } = useWallet()

  const terraAddress = useMemo(() => wallets?.[0]?.terraAddress, [wallets?.[0]?.terraAddress])
  // const terraAddress = wallets?.[0]?.terraAddress
  return ({
    ...rest,
    wallets,
    terraAddress,
    connect,
    connectTerraStation: async () => {


      return connect(ConnectType.EXTENSION)
    },
  })
}