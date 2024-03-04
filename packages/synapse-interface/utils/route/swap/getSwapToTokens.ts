import { EXISTING_SWAP_ROUTES } from '@/constants/existingSwapRoutes'
import { getToTokensFunc } from '@/utils/route/getToTokensFunc'

export const getSwapToTokens = getToTokensFunc(EXISTING_SWAP_ROUTES)