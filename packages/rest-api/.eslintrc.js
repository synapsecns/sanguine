module.exports = {
  extends: '../../.eslintrc.js',
  overrides: [
    {
      files: ['jest.config.js'],
      rules: {
        'prettier/prettier': 'off',
      },
    },
    {
      files: ['**/*.ts'],
      rules: {
        'guard-for-in': 'off',
        'jsdoc/check-indentation': 'off',
        '@typescript-eslint/no-empty-function': 'off',
        '@typescript-eslint/no-unused-vars': 'off',
      },
    },
  ],
}
