import typescript from 'rollup-plugin-typescript2'
import { nodeResolve } from '@rollup/plugin-node-resolve'
import commonjs from '@rollup/plugin-commonjs'
import json from '@rollup/plugin-json'
import terser from '@rollup/plugin-terser'
import url from '@rollup/plugin-url'
import { codecovRollupPlugin } from '@codecov/rollup-plugin'

export default {
  input: 'src/index.ts', // Entry point for your constants/utilities
  output: [
    {
      file: 'dist/bundle.cjs.js',
      format: 'cjs', // CommonJS output
      sourcemap: true,
    },
    {
      file: 'dist/bundle.esm.js',
      format: 'esm', // ES Module output
      sourcemap: true,
    },
  ],
  plugins: [
    nodeResolve({
      preferBuiltins: true,
    }),
    commonjs(),
    json(),
    typescript({
      tsconfig: './tsconfig.json',
      declaration: true,
      declarationDir: 'dist',
    }),
    terser(),
    url({
      include: ['**/*.svg', '**/*.png', '**/*.jpg'],
      limit: 0,
      emitFiles: false,
    }),
    codecovRollupPlugin({
      enableBundleAnalysis: process.env.CODECOV_TOKEN !== undefined,
      bundleName: 'synapse-constants',
      uploadToken: process.env.CODECOV_TOKEN,
      uploadOverrides: {
        sha: process.env.GH_COMMIT_SHA,
      },
    }),
  ],
  external: ['lodash', 'ethers'],
}
