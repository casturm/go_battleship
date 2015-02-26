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
      Game.player($scope.Player.name).then(function() {
        console.log("player registered");
        Game.getPlayers().then(function() {
          console.log("got players");
          $scope.players = Game.players()
          $location.path( "/players" );
        });
        //Game.get().then(function() {
          //if (Game.current()) {
            //console.log("game started " + Game.current());
            //$location.path( "/view2" );
          //}
          //else {
            //console.log("waiting for another player")
          //}
        //});
      });
    }
  }
}]);
