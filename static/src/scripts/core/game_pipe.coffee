class GamePipe
    constructor: (username) ->
        console.log 'game pipe created'
        @create_socket username

    create_socket: (username)->
        @socket = new WebSocket("ws://"+location.host+"/game_ws?name="+username)

module.exports = GamePipe