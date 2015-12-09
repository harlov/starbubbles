module.exports = resources: files: [{
    expand: true
    src: [ 'src/resources/img/*' ]
    dest: 'dist/resources/img'
    flatten: true
    }]