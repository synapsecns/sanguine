import { EXISTING_SWAP_ROUTES } from '@/constants/existingSwapRoutes'
import { getFromTokensFunc } from '@/utils/route/getFromTokensFunc'

export const getSwapFromTokens = getFromTokensFunc(EXISTING_SWAP_ROUTES)