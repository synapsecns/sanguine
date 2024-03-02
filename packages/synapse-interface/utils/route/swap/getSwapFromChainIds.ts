import _ from 'lodash'

import {
  EXISTING_SWAP_ROUTES,
  SWAP_CHAIN_IDS,
} from '@/constants/existingSwapRoutes'

import { getFromChainIdsFunc } from '@/utils/route/getFromChainIdsFunc'


export const getAllFromChainIds = () => SWAP_CHAIN_IDS

export const getSwapFromChainIds = getFromChainIdsFunc(EXISTING_SWAP_ROUTES)