class Controls
    constructor: (game) ->
        ctx = @
        @game = game
        @canvas = game.canvas
        @dc = @canvas.getContext '2d'

        @centerPos = 
            x: @canvas.width/2
            y: @canvas.height/2

        document.onmousemove = (e) ->
            ctx.mousePos = 
                x : e.pageX;
                y : e.pageY;



    drawDirectionLine: ->
        if not @mousePos
            return
        @dc.beginPath()
        @dc.moveTo @centerPos.x, @centerPos.y
        @dc.lineTo @mousePos.x,@mousePos.y
        @dc.stroke()

    calcDirectionAngel: ->
        if not @mousePos
            return
        trinX = @mousePos.x - @centerPos.x
        trinY = @mousePos.y - @centerPos.y
        trinAtan = Math.atan2(trinY, trinX)
        trianAngel = trinAtan

        command =
            Method: 'updateDirection'
            Params:
                'directionAngelRad' : trianAngel

        @game.pipe.sendCommand(command)

    update: ->
        @drawDirectionLine()
        @calcDirectionAngel()



module.exports = Controls