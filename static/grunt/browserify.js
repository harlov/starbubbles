module.exports = {
   dist: {
    files: {
      'dist/app.js': ['src/scripts/**/*.coffee']
    },
    options: {
      transform: ['coffeeify']
    }
  }
};