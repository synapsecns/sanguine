import resolve from '@rollup/plugin-node-resolve'
import commonjs from '@rollup/plugin-commonjs'
import typescript from '@rollup/plugin-typescript'
import dts from 'rollup-plugin-dts'
import postcss from 'rollup-plugin-postcss'
import json from '@rollup/plugin-json'
import terser from '@rollup/plugin-terser'
import peerDepsExternal from 'rollup-plugin-peer-deps-external'
import sourcemaps from 'rollup-plugin-sourcemaps'

import packageJson from './package.json' assert { type: 'json' }

export default [
  {
    input: 'src/index.tsx',
    output: [
      {
        file: packageJson.main,
        sourcemap: true,
        format: 'cjs',
      },
      {
        file: packageJson.module,
        sourcemap: true,
        format: 'esm',
      },
    ],
    external: [
      'react',
      'react/jsx-runtime',
      'ethers',
      'react-redux',
      '@reduxjs/toolkit',
      '@synapsecns/sdk-router',
      '@ethersproject/providers',
      '@ethersproject/units',
      'lodash',
    ],
    plugins: [
      peerDepsExternal(),
      resolve(),
      commonjs(),
      typescript({
        tsconfig: './tsconfig.json',
        declaration: true,
        declarationDir: 'dist',
      }),
      postcss({
        plugins: [],
      }),
      json(),
      terser(),
      sourcemaps(),
    ],
    watch: {
      buildDelay: 200,
    },
  },
  {
    input: 'src/types/index.d.ts',
    output: [{ file: 'dist/index.d.ts', format: 'esm' }],
    plugins: [dts.default()],
    external: [/\.(css|less|scss)$/],
  },
]
