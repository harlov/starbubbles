GamePipe = require './core/game_pipe'
Controls = require './core/controls'
$ = window.jQuery = require ('jquery')

class Game
    constructor: (username) ->
        ctx = @
        console.log username
        @loadResources()
        @pipe = new GamePipe username
        @pipe.subscribePipe((data)-> ctx.pipeRecive data )
        @initCanvas()
        @controls = new Controls ctx

    pipeRecive: (data) ->
        ball = data.Params.Balls[0]
        console.log(@controls.centerPos)
        @clearCanvas()
        
        @drawBack()
        @drawPlayerBubble @controls.centerPos.x, @controls.centerPos.y, ball.Mass
        @controls.update()
        
        
    loadResources: () ->
        @resources = {}
        img_back = new Image()
        img_back.src = '/resources/img/back_1.png'

        @resources['back'] = img_back

    drawBack: () ->
        @dc.drawImage(@resources['back'], 0, 0)

    initCanvas: ->        
        $('#game-wrap').append('<canvas id="game-container"></canvas>')
        @canvas = $('#game-container')[0];
        
        $('#game-container').attr('width', $('#game-wrap').width()+'px')
        $('#game-container').attr('height', $('#game-wrap').height()+'px')

        @dc = @canvas.getContext '2d'

    clearCanvas: ->
        @dc.clearRect(0,0,1000,1000);

    drawPlayerBubble: (x,y, mass) ->
        @dc.beginPath()
        @dc.fillStyle = 'green'
        @dc.arc(x, y, mass/2, 0, Math.PI*2, false)
        @dc.fill()
        @dc.closePath()
        

module.exports = Game