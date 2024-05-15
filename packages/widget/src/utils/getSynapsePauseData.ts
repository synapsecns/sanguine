export const getSynapsePauseData = () => {
  const fetchJSONData = async (url) => {
    const response = await fetch(url)
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    return response.json()
  }

  const fetchPauseData = async () => {
    const pausedChainsUrl =
      'https://raw.githubusercontent.com/synapsecns/sanguine/master/packages/synapse-interface/public/pauses/v1/paused-chains.json'
    const pausedModulesUrl =
      'https://raw.githubusercontent.com/synapsecns/sanguine/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json'

    try {
      console.log('fetching and storing pause data in client browser')
      const chainsData = await fetchJSONData(pausedChainsUrl)
      const modulesData = await fetchJSONData(pausedModulesUrl)

      localStorage.setItem('synapse-chain-pause', JSON.stringify(chainsData))
      localStorage.setItem('synapse-module-pause', JSON.stringify(modulesData))
      localStorage.setItem('synapse-pause-timestamp', Date.now().toString())
    } catch (error) {
      console.error('Failed to fetch paused chains/modules:', error)
    }
  }

  const readPauseData = () => {
    const chainPause = JSON.parse(localStorage.getItem('synapse-chain-pause'))
    const modulePause = JSON.parse(localStorage.getItem('synapse-module-pause'))

    return { chainPause, modulePause }
  }

  const checkIsDataValid = () => {
    const lastFetchTime = localStorage.getItem('synapse-pause-timestamp')

    if (lastFetchTime) {
      const previousTime = Number(lastFetchTime)
      const currentTime = Date.now()

      const millisecondsPerHour = 1000 * 60 * 60 // milliseconds in an hour
      const timePastInHours = (currentTime - previousTime) / millisecondsPerHour

      // return timePastInHours < 24
      return false
    } else {
      return false
    }
  }

  const isValid = checkIsDataValid()

  if (!isValid) {
    console.log('refetching synapse pause')
    fetchPauseData()
  }

  return readPauseData()
}
