import { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  overwrite: true,
  documents: 'graphql/*.graphql',
  schema: 'https://explorer.omnirpc.io/graphql',
  generates: {
    './slices/api/generated.ts': {
      plugins: [
        'typescript',
        'typescript-operations',
        {
          'typescript-rtk-query': {
            importBaseApiFrom: 'slices/api/slice',
            exportHooks: true,
          },
        },
      ],
    },
  },
}

export default config
