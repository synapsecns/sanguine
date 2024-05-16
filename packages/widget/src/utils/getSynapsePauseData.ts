// const PAUSED_CHAINS_URL =
//   'https://raw.githubusercontent.com/synapsecns/sanguine/master/packages/synapse-interface/public/pauses/v1/paused-chains.json'
// const PAUSED_MODULES_URL =
//   'https://raw.githubusercontent.com/synapsecns/sanguine/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json'

const PAUSED_CHAINS_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/test/pauses/packages/synapse-interface/public/pauses/v1/paused-chains.json'
const PAUSED_MODULES_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/test/pauses/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json'

enum SessionStorageKey {
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

      sessionStorage.setItem(
        SessionStorageKey.CHAIN_PAUSE,
        JSON.stringify(chainsData)
      )
      sessionStorage.setItem(
        SessionStorageKey.MODULE_PAUSE,
        JSON.stringify(modulesData)
      )
      sessionStorage.setItem(SessionStorageKey.TIMESTAMP, Date.now().toString())
    } catch (error) {
      console.error(
        '[Synapse Widget] Failed to fetch paused chains/modules: ',
        error
      )
    }
  }

  const readData = () => {
    if (sessionStorage) {
      const chainPause = JSON.parse(
        sessionStorage.getItem(SessionStorageKey.CHAIN_PAUSE)
      )
      const modulePause = JSON.parse(
        sessionStorage.getItem(SessionStorageKey.MODULE_PAUSE)
      )
      return { chainPause, modulePause }
    }
    return { chainPause: null, modulePause: null }
  }

  const checkIsDataValid = (): boolean => {
    const lastFetchTime = sessionStorage.getItem(SessionStorageKey.TIMESTAMP)

    if (lastFetchTime) {
      const previousTime = Number(lastFetchTime)
      const currentTime = Date.now()

      const msPerHr = 1000 * 60 * 60
      const timeElapsedInHrs = (currentTime - previousTime) / msPerHr

      return timeElapsedInHrs < 1
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

const fetchJSONData = async (url: string): Promise<any> => {
  const response = await fetch(url)
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`)
  }
  return response.json()
}
