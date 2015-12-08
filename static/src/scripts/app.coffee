Game = require './game'
$ = window.jQuery = require ('jquery')

class window.Application
    constructor: ->
        console.log 'create application instance'
    getUserLogin: ->
        return $('#player_name')[0].value
    afterLogin: ->
        $('#content-wrap').addClass('hided')
        $('#game-wrap').removeClass('hided')

    newGame: ->
        username = @getUserLogin()
        console.log username
        @game = new Game username
        @afterLogin()


$ ->
    window.app = new Application

module.exports = window.Application