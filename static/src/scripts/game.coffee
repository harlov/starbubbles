GamePipe = require './core/game_pipe'
$ = window.jQuery = require ('jquery')

class Game
    constructor: (username) ->
        console.log username
        @game_pipe = new GamePipe username
        @initCanvas()
        @drawPlayerBubbles()

    initCanvas: ->
        @canvas = $('#game-container')[0]
        @dc = @canvas.getContext '2d'

    drawPlayerBubbles: ->
        @dc.beginPath()
        @dc.fillStyle = 'green'
        @dc.arc(50, 50, 15, 0, Math.PI*2, false)
        @dc.fill();
        @dc.closePath();
        

module.exports = Game