import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    coverage: {
      reporter: ['text', 'html'],
      exclude: ['node_modules/'],
    },
    testTimeout: 30000,
  },
})
