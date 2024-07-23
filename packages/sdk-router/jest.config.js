module.exports = {
  maxConcurrency: 1,
  testTimeout: 30000,
  moduleNameMapper: {
    /**
     * Force module uuidv7 to resolve with the CJS entry point,
     * because Jest does not support package.json.exports.
     * See https://github.com/uuidjs/uuid/issues/451
     */
    "uuidv7": require.resolve('uuidv7'),
  }
}
