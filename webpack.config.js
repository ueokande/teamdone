var path = require('path');

const src = path.resolve(__dirname, 'assets');
const dist = path.resolve(__dirname, 'build');

module.exports = {
  entry: {
    index: path.join(src, 'index.js'),
  },

  output: {
    path: dist,
    filename: '[name].js'
  },

  module: {
    loaders: [
      {
        test: /\.jsx$/,
        exclude: [
          /node_modules/,
          /_test.jsx$/
        ],
        loader: 'babel-loader',
        query: {
          presets: ['es2015', 'react']
        }
      },
      {
        test: /\.js$/,
        exclude: [
          /node_modules/,
          /_test.js$/
        ],
        loader: 'babel-loader',
        query: {
          presets: ['es2015']
        }
      },
      {
        test: /\.css$/,
        loader: 'style-loader!css-loader',
      },
      {
        test: /\.scss$/,
        loader: 'style-loader!css-loader!sass-loader',
      },
    ]
  },

  resolve: {
    extensions: ['.js', '.jsx']
  },
};
