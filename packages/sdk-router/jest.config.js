module.exports = {
  roots: ['<rootDir>/src'],
  maxConcurrency: 1,
  testTimeout: 30000,
  testEnvironment: 'node',
  transform: {
    '^.+\\.(ts|tsx)$': ['ts-jest', { tsconfig: '<rootDir>/tsconfig.test.json' }],
  },
  moduleFileExtensions: ['ts', 'tsx', 'js', 'jsx', 'json', 'node'],
  collectCoverageFrom: ['src/**/*.{ts,tsx,js,jsx}'],
  testMatch: ['<rootDir>/src/**/*.(spec|test).{ts,tsx,js,jsx}'],
  moduleNameMapper: {
    /**
     * Force module uuidv7 to resolve with the CJS entry point,
     * because Jest does not support package.json.exports.
     * See https://github.com/uuidjs/uuid/issues/451
     */
    uuidv7: require.resolve('uuidv7'),
  },
}
