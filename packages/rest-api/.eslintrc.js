module.exports = {
  extends: '../../.eslintrc.js',
  overrides: [
    {
      files: ['jest.config.js'],
      rules: {
        'prettier/prettier': 'off',
        'guard-for-in': 'off',
      },
    },
    {
      files: ['**/*.ts', '**/*.tsx'],
      rules: {
        'guard-for-in': 'off',
      },
    },
  ],
}
