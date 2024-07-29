export const fetchJsonData = async (url: string): Promise<any> => {
  // Configurable parameters
  const maxRetries = 3 // maximum number of retries
  const initialDelay = 1000 // initial delay in milliseconds

  // Helper function to delay for a given amount of time
  const delay = (duration) =>
    new Promise((resolve) => setTimeout(resolve, duration))

  for (let attempt = 0; attempt < maxRetries; attempt++) {
    try {
      const response = await fetch(url, { method: 'GET', cache: 'no-store' })

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()
      return data
    } catch (error) {
      console.error(`Attempt ${attempt + 1}: Error fetching ${url}`, error)

      // If it's the last attempt, reject the promise
      if (attempt === maxRetries - 1) {
        throw new Error(`Retries failed fetching ${url}`)
      }

      // Exponential backoff
      const delayDuration = initialDelay * 2 ** attempt
      console.log(`Retrying in ${delayDuration}ms...`)
      await delay(delayDuration)
    }
  }
}
