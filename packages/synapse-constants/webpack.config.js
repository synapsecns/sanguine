const path = require('path');


const config = {
  stats: {
    errorDetails: true,
  },
  resolve: {
    extensions: ['.ts', '.js'],
  },
  mode: 'production',
  devtool: 'source-map',
  entry: './dist/index.js',
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js',
    libraryTarget: 'umd',
  },
  target: 'web',
  module: {
    rules: [
      {
        test: /\.(js|ts)$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
        },
      },
      {
        test: /\.svg$/,
        use: [
          {
            loader: 'svg-inline-loader',
          },
        ],
      },
      {
        test: /\.(png|jpe?g)$/i,
        use: [
          {
            loader: 'url-loader',
          },
        ],
      },
    ],
  },
};

module.exports = config
