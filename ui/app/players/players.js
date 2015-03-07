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
  $scope.games = Game.games()
  $scope.currentPlayer = Game.getThisPlayer();

  $scope.play = function(gameId) {
    console.log("loadGame " + gameId);
    Game.loadGame(gameId).then(function() {
      $location.path("/view2");
    });
  }
  $scope.createGame = function(player) {
    Game.createGame(player).then(function() {
      $scope.player1 = Game.getPlayer(0);
      $scope.player2 = Game.getPlayer(1);
      console.log(angular.toJson($scope.player1));
      console.log(angular.toJson($scope.player2));
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
  $scope.reloadGames = function() {
    console.log("reloadGames");
    Game.getGames().then(function() {
      console.log("reset games list");
      $scope.games = Game.games();
    });
  };
}]);
