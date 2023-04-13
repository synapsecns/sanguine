import { defineConfig } from 'cypress'

export default defineConfig({
  e2e: {
    baseUrl: 'http://localhost:3000/',
  },
  video: true,
  trashAssetsBeforeRuns: true,
  screenshotOnRunFailure: true,
  screenshotsFolder: 'cypress/visual-states/current-screenshots',
  videosFolder: 'cypress/visual-states/current-videos',
})
