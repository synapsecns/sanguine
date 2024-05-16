const PAUSED_CHAINS_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/test/pauses/packages/synapse-interface/public/pauses/v1/paused-chains.json'
const PAUSED_MODULES_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/527d855ce9b1d908874760628d35fb2c4a6ec5dc/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json'

enum LocalStorageKey {
  CHAIN_PAUSE = 'synapse-paused-chains',
  MODULE_PAUSE = 'synapse-paused-modules',
  TIMESTAMP = 'synapse-paused-data-timestamp',
}

export const getSynapsePauseData = () => {
  const fetchData = async () => {
    try {
      console.log('[Synapse Widget] Fetching pause data')
      const chainsData = await fetchJSONData(PAUSED_CHAINS_URL)
      const modulesData = await fetchJSONData(PAUSED_MODULES_URL)

      localStorage.setItem(
        LocalStorageKey.CHAIN_PAUSE,
        JSON.stringify(chainsData)
      )
      localStorage.setItem(
        LocalStorageKey.MODULE_PAUSE,
        JSON.stringify(modulesData)
      )
      localStorage.setItem(LocalStorageKey.TIMESTAMP, Date.now().toString())
    } catch (error) {
      console.error(
        '[Synapse Widget] Failed to fetch paused chains/modules: ',
        error
      )
    }
  }

  const readData = () => {
    const chainPause = JSON.parse(
      localStorage.getItem(LocalStorageKey.CHAIN_PAUSE)
    )
    const modulePause = JSON.parse(
      localStorage.getItem(LocalStorageKey.MODULE_PAUSE)
    )

    return { chainPause, modulePause }
  }

  const checkIsDataValid = () => {
    const lastFetchTime = localStorage.getItem(LocalStorageKey.TIMESTAMP)

    if (lastFetchTime) {
      const previousTime = Number(lastFetchTime)
      const currentTime = Date.now()

      const millisecondsPerHour = 1000 * 60 * 60 // milliseconds in an hour
      const timePastInHours = (currentTime - previousTime) / millisecondsPerHour

      return false
    } else {
      return false
    }
  }

  const isValid = checkIsDataValid()

  if (!isValid) {
    fetchData()
  }

  return readData()
}

const fetchJSONData = async (url: string) => {
  const response = await fetch(url)
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`)
  }
  return response.json()
}
