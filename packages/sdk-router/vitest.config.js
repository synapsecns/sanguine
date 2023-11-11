import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    coverage: {
      reporter: ['text', 'html'],
      exclude: ['node_modules/'],
    },
    // We are using larger timeout values than usual because our tests are currently
    // being run against forked blockchains.
    hookTimeout: 30000,
    testTimeout: 30000,
  },
})
