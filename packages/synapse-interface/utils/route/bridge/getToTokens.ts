import { EXISTING_BRIDGE_ROUTES } from '@/constants/existingBridgeRoutes'
import { getToTokensFunc } from '@/utils/route/getToTokensFunc'

export const getToTokens = getToTokensFunc(EXISTING_BRIDGE_ROUTES)