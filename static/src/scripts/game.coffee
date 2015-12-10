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
        ctx = @
        ball = data.Params.Balls[0]
        console.log(@controls.centerPos)
        requestAnimationFrame () ->
            ctx.clearCanvas()
            ctx.drawBack(ball.Position.X, ball.Position.Y)
            ctx.drawPlayerBubble ctx.controls.centerPos.x, ctx.controls.centerPos.y, ball.Mass
            ctx.controls.update()

        
    loadResources: () ->
        @resources = {}
        img_back = new Image()
        img_back.src = '/resources/img/back_1.png'

        @resources['back'] = img_back

    drawBorder: (cX, cY) ->
        @dc.moveTo @centerPos.x, @centerPos.y
        @dc.lineTo @mousePos.x,@mousePos.y
        @dc.stroke()

    drawBack: (cX, cY) ->
        bw = 135
        bh = 135

        startPosX = - ( cX % bw )
        startPosY = - ( cY % bh )

        console.log(startPosX, startPosY)

        posX = startPosX
        posY = startPosY
        fW = @field['width']
        fH = @field['height']
        while posX < fW
            posY = startPosY
            while posY < fH
                @dc.drawImage(@resources['back'], posX, posY)
                posY += bh
            posX += bw

    initCanvas: ->        
        $('#game-wrap').append('<canvas id="game-container"></canvas>')
        @canvas = $('#game-container')[0];
        
        @field = 
            'width': $('#game-wrap').width()
            'height': $('#game-wrap').height()

        $('#game-container').attr('width', @field['width']+'px')
        $('#game-container').attr('height', @field['height']+'px')

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