Game = require './game'
$ = window.jQuery = require ('jquery')

class window.Application
    constructor: ->
        console.log 'create application instance'
    getUserLogin: ->
        return $('#player_name')[0].value
    afterLogin: ->
        $('#login-container').addClass('hided')
        $('#game-container').removeClass('hided')

    newGame: ->
        username = @getUserLogin()
        console.log username
        @game = new Game username
        @afterLogin()


$ ->
    window.app = new Application

module.exports = window.Application