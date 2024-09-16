module.exports = {
  extends: '../../.eslintrc.js',
  overrides: [
    {
      files: ['jest.config.js'],
      rules: {
        'prettier/prettier': 'off',
      },
    },
  ],
}
