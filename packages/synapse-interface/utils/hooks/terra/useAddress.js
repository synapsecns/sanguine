import { useConnectedWallet } from "@terra-money/wallet-provider"

export function useAddress() {
  const connectedWallet = useConnectedWallet()
  return connectedWallet?.terraAddress ?? ""
}
