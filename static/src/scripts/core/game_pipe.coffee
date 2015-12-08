class GamePipe
    constructor: (username) ->
        console.log 'game pipe created'
        @createSocket username
        @dataCbs = []

    createSocket: (username)->
        ctx = @
        @socket = new WebSocket("ws://"+location.host+"/game_ws?name="+username)
        @socket.onopen = @onSocketOpen
        @socket.onclose = @onSocketClose
        @socket.onmessage = (evt) -> ctx.onSocketMessage(evt)
        @socket.onerror = @onSocketError

    subscribePipe: (cb) ->
        @dataCbs.push(cb)

    onSocketOpen: ->

    onSocketClose: ->

    onSocketMessage: (evt) ->
        mess = JSON.parse(evt.data)
        for cb in @dataCbs
            cb(mess)

    onSocketError: (err) ->
        console.error err

module.exports = GamePipe