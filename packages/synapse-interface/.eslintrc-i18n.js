const { ignorePatterns } = require('./.eslintrc')

module.exports = {
  root: true,
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 2021,
    sourceType: 'module',
    ecmaFeatures: {
      jsx: true,
    },
  },
  plugins: ['i18next'],
  ignorePatterns: [
    'pages/lifi/index.tsx',
    'components/Maintenance/example/EcotoneForkUpgrade.tsx',
    'pages/_app.tsx',
  ],
  rules: {
    // Enable the i18next rule
    'i18next/no-literal-string': 'error',

    // Disable unrelated rules
    '@typescript-eslint/no-explicit-any': 'off',
    '@typescript-eslint/explicit-module-boundary-types': 'off',
    '@typescript-eslint/no-non-null-assertion': 'off',
    'prefer-arrow/prefer-arrow-functions': 'off',
    'import/order': 'off',
    'prefer-const': 'off',
    'jsdoc/newline-after-description': 'off',
    'no-duplicate-imports': 'off',
    eqeqeq: 'off',
    'no-return-await': 'off',
    'object-shorthand': 'off',
    'prettier/prettier': 'off',
    'no-redeclare': 'off',
    radix: 'off',
    'jsdoc/check-alignment': 'off',
    'no-undef-init': 'off',
  },
  // Ignore all other plugins and extends
  extends: [],
}
