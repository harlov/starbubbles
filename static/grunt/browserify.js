module.exports = {
   dist: {
    files: {
      'dist/app.js': ['src/scripts/**/*.coffee']
    },
    options: {
        browserifyOptions:  {
            extensions: '.coffee'
        },
        transform: ['coffeeify']
    }
  }
};