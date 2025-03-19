module.exports = {
  extends: '../../.eslintrc.js',
  // Override the import/order rule to enforce alphabetical ordering
  rules: {
    'import/order': [
      'error',
      {
        groups: ['builtin', 'external', 'internal'],
        'newlines-between': 'always',
        alphabetize: {
          order: 'asc', // Sort in ascending order (a-z)
          caseInsensitive: true, // Ignore case
        },
      },
    ],
  },
}
