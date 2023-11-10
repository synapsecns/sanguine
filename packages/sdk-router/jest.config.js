module.exports = {
  setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
  testTimeout: 30000,
  transform: {
    '\\.[jt]sx?$': 'babel-jest',
  },
}
