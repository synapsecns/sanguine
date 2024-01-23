const path = require('path')

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
    minimize: true,
  },
}
