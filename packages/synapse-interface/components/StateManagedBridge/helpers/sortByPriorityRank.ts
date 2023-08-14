import _ from 'lodash'

import { Token } from '@/utils/types'

export const sortByPriorityRank = (tokens: Token[]) => {
  return _.orderBy(tokens, ['priorityRank', 'symbol'], ['asc', 'asc'])
}
