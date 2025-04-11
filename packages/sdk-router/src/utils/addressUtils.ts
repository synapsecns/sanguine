import { AddressZero } from '@ethersproject/constants'

export const ETH_NATIVE_TOKEN_ADDRESS =
  '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'

export const isSameAddress = (
  addr1?: string | null,
  addr2?: string | null
): boolean => {
  return !!addr1 && !!addr2 && addr1.toLowerCase() === addr2.toLowerCase()
}

export const handleNativeToken = (tokenAddr: string) => {
  return isNativeToken(tokenAddr) ? ETH_NATIVE_TOKEN_ADDRESS : tokenAddr
}

export const isNativeToken = (tokenAddr?: string): boolean => {
  return (
    tokenAddr === '' ||
    tokenAddr === AddressZero ||
    isSameAddress(tokenAddr, ETH_NATIVE_TOKEN_ADDRESS)
  )
}
