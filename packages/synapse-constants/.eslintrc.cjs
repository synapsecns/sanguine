module.exports = {
  extends: '../../.eslintrc.js',
  overrides: [
    {
      files: ['**/*.ts'],
      rules: {
        'guard-for-in': 'off',
        'prefer-arrow/prefer-arrow-functions': 'off',
      },
    },
  ],
}
