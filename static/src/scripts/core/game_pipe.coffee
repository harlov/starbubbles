class GamePipe
    constructor: ->
        console.log 'game pipe created'
        @create_socket()

    create_socket: ->
        @socket = new WebSocket("ws://"+location.host+"/game_ws")

module.exports = GamePipe