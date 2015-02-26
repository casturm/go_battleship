'use strict';

angular.module('myApp.players', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/players', {
    templateUrl: 'players/players.html',
    controller: 'PlayersCtrl'
  });
}])

.controller('PlayersCtrl', ['$scope', '$location', 'Game', function($scope, $location, Game) {
  console.log("PlayersCtrl" + Game);
  $scope.players = Game.players()
  $scope.createGame = function(player) {
    Game.createGame(player).then(function() {
      $location.path("/view2");
    });
  };
  $scope.reload = function() {
    console.log("reload");
    Game.getPlayers().then(function() {
      console.log("reset players list");
      $scope.players = Game.players();
    });
  };
}]);
