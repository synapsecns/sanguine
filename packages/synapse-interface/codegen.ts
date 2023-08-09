import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  overwrite: true,
  schema: 'https://explorer.omnirpc.io/graphiql',
  documents: "'src/**/**/!*.d.{ts,tsx}'",
  generates: {
    './src/slices/api/': {
      preset: 'client',
      plugins: [],
    },
    './graphql.schema.json': {
      plugins: ['introspection'],
    },
  },
}

export default config
