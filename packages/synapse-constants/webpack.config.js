const path = require('path')

const ImageMinimizerPlugin = require('image-minimizer-webpack-plugin')

module.exports = {
  mode: 'production',

  entry: './dist/',

  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js',
  },

  resolve: {
    extensions: ['.ts', '.tsx', '.js'],
    modules: [path.resolve(__dirname, '../../node_modules')],
  },

  module: {
    rules: [
      {
        test: /\.ts|tsx$/,
        loader: 'ts-loader',
        exclude: /node_modules/,
      },
      {
        test: /\.svg$/,
        loader: 'svg-inline-loader',
      },
      {
        test: /\.(png|jpg|jpeg|gif)$/,
        type: 'asset',
      },
    ],
  },
  optimization: {
    minimizer: [
      '...',
      new ImageMinimizerPlugin({
        minimizer: {
          implementation: ImageMinimizerPlugin.imageminMinify,
          options: {
            // Lossless optimization with custom option
            // Feel free to experiment with options for better result for you
            plugins: [
              ['optipng', { optimizationLevel: 5 }],
              // Svgo configuration here https://github.com/svg/svgo#configuration
              [
                'svgo',
                {
                  plugins: [
                    {
                      name: 'preset-default',
                      params: {
                        overrides: {
                          removeViewBox: false,
                          inlineStyles: {
                            onlyMatchedOnce: false,
                          },
                        },
                      },
                    },
                  ],
                  multipass: true,
                },
              ],
            ],
          },
        },
      }),
    ],
  },
}
