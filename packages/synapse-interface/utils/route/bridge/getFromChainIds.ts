import _ from 'lodash'

import { EXISTING_BRIDGE_ROUTES } from '@/constants/existingBridgeRoutes'
import { getTokenAndChainId } from '@/utils/route/getTokenAndChainId'
import { getFromChainIdsFunc } from '@/utils/route/getFromChainIdsFunc'


export const getAllFromChainIds = () => {
  return _(EXISTING_BRIDGE_ROUTES)
    .keys()
    .map((token) => getTokenAndChainId(token).chainId)
    .uniq()
    .value()
}

export const getFromChainIds = getFromChainIdsFunc(EXISTING_BRIDGE_ROUTES)