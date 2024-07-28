jest.mock('node:crypto', () => {
  return {
    ...jest.requireActual('crypto'),
    randomBytes: jest.fn(() => Buffer.from('mockedRandomBytes')),
  }
})
