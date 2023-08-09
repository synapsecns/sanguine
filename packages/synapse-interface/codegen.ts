import { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  overwrite: true,
  schema: 'https://explorer.omnirpc.io/graphql',
  generates: {
    './slices/api/generated.ts': {
      plugins: [
        'typescript',
        'typescript-operations',
        {
          'typescript-rtk-query': {
            importBaseApiFrom: './slice',
            exportHooks: true,
          },
        },
      ],
    },
  },
}

export default config
