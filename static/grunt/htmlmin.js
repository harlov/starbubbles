module.exports = {
    prod: {                                      // Target
      options: {                                 // Target options
        removeComments: true,
        collapseWhitespace: true
      },
      files: {                                   // Dictionary of files
        'dist/index.html': 'src/index.html'   // 'destination': 'source'
      }
    },
     dev: {                                       // Another target
      files: {
        'dist/index.html': 'src/index.html'
      }
    }
};