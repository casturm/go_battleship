'use strict';

angular.module('myApp.view1', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/view1', {
    templateUrl: 'view1/view1.html',
    controller: 'View1Ctrl'
  });
}])

.controller('View1Ctrl', ['$scope', '$location', 'Game', function($scope, $location, Game) {
  console.log("View1Ctrl " + Game);
  $scope.Player = {}
  $scope.start = function(valid) {
    console.log("Player.name: " + $scope.Player.name + " " + valid);
    $scope.submitted = true;
    if (valid) {
      Game.newPlayer($scope.Player.name).then(function() {
        console.log("player registered");
        Game.getPlayers().then(function() {
          console.log("got players");
          $scope.players = Game.players()
          $scope.currentPlayer = Game.getThisPlayer();
          Game.getGames().then(function() {
            console.log("got games");
            $scope.games = Game.games();
            $location.path( "/players" );
          });
        });
      });
    }
  }
}]);
