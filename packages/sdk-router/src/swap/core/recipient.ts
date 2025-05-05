import { AddressZero } from '@ethersproject/constants'

export enum RecipientEntity {
  Self,
  User,
  UserSimulated,
}

export type Recipient = {
  entity: RecipientEntity
  address: string
}

export const USER_SIMULATED_ADDRESS =
  '0xFAcefaCEFACefACeFaCefacEFaCeFACEFAceFAcE'

export const getForwardTo = (recipient: Recipient): string => {
  return recipient.entity === RecipientEntity.Self
    ? AddressZero
    : recipient.address
}
