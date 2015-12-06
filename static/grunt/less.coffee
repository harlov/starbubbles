module.exports =
      options:
        paths: ['src/less/']
      dist:
        options:
          sourceMap: true
          sourceMapURL: 'styles.css.map'
        files:
          'dist/styles.css': ['src/**/*.less', '!src/less/import/**/*.less']
      build:
        options:
          sourceMap: false
          compress: true
        files:
          'build/styles.min.css': ['src/**/*.less', '!src/less/import/**/*.less']