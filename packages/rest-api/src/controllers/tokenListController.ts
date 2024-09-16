import * as tokenList from '../constants/bridgeable'

export const tokenListController = async (_req, res) => {
  res.json(tokenList)
}
