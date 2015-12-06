GamePipe = require './core/game_pipe'

class window.Application
    constructor: ->
        console.log 'create application instance'
    startGame: ->
        @game_pipe = new GamePipe()

module.exports = window.Application