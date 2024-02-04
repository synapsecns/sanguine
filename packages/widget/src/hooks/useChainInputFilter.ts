import { useState, useEffect } from 'react'
import _ from 'lodash'
import { Chain } from 'types'

export const useChainInputFilter = (
  options: Chain[],
  remaining: Chain[],
  targets: Chain[],
  isActive: boolean
) => {
  const [filterValue, setFilterValue] = useState('')

  useEffect(() => {
    if (!isActive) {
      setFilterValue('')
    }
  }, [isActive])

  const filterChains = (chains: Chain[], filter: string) => {
    const lowerFilter = filter.toLowerCase()
    return _.filter(chains, (option) => {
      const name = option.name.toLowerCase()
      return name.includes(lowerFilter) || name === lowerFilter
    })
  }

  const filteredOptions = filterChains(options, filterValue)
  const filteredRemaining = filterChains(remaining, filterValue)
  const filteredTargets = filterChains(targets, filterValue)

  const hasFilteredOptions = !_.isEmpty(filteredOptions)
  const hasFilteredRemaining = !_.isEmpty(filteredRemaining)
  const hasFilteredResults = hasFilteredOptions || hasFilteredRemaining
  const hasFilteredTargets = !_.isEmpty(filteredTargets)

  return {
    filterValue,
    setFilterValue,
    filteredOptions,
    filteredRemaining,
    filteredTargets,
    hasFilteredOptions,
    hasFilteredRemaining,
    hasFilteredResults,
    hasFilteredTargets,
  }
}
