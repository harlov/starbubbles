GamePipe = require './core/game_pipe'
$ = window.jQuery = require ('jquery')

class Game
    constructor: (username) ->
        ctx = @
        console.log username
        @game_pipe = new GamePipe username
        @game_pipe.subscribePipe((data)-> ctx.pipeRecive data )
        @initCanvas()        

    pipeRecive: (data) ->
        ball = data.Params.Balls[0]
        @drawPlayerBubble(ball.Position.X, ball.Position.Y, ball.Mass)

    initCanvas: ->
        @canvas = $('#game-container')[0]
        @dc = @canvas.getContext '2d'

    clearCanvas: ->
        @dc.clearRect(0,0,1000,1000);

    drawPlayerBubble: (x,y, mass) ->
        @clearCanvas()
        @dc.beginPath()
        @dc.fillStyle = 'green'
        @dc.arc(x, y, mass, 0, Math.PI*2, false)
        @dc.fill();
        @dc.closePath();
        

module.exports = Game