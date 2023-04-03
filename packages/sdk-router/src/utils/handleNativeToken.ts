import { AddressZero } from '@ethersproject/constants'

export const ETH_NATIVE_TOKEN_ADDRESS =
  '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'

export const handleNativeToken = (tokenAddr: string) => {
  return tokenAddr === '' || tokenAddr === AddressZero
    ? ETH_NATIVE_TOKEN_ADDRESS
    : tokenAddr
}
