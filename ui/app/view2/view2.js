'use strict';

angular.module('myApp.view2', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/view2', {
    templateUrl: 'view2/view2.html',
    controller: 'View2Ctrl'
  });
}])

.controller('View2Ctrl', ['$scope', '$location', 'Game', function($scope, $location, Game) {

  $scope.player1Rows = Game.rows(0,1);
  $scope.player2Rows = Game.rows(1,0);

  $scope.turn = function(x,y) {
    console.log("take turn 0");
    if (Game.current.gameOn) {
      Game.turn(0,x,y).then(function() {
        console.log("turn complete 0")
        //console.log(Game.current)
        //$location.path("/view3");
      });
    }
    else {
      alert("place ships before you can start the game");
    }
  }

  $scope.Ship = {}
  $scope.addShip = function(valid) {
    $scope.submitted = true;
    if (valid) {
      console.log("add ship " + angular.toJson($scope.Ship));
      Game.addShip($scope.Ship).then(function() {
        console.log("ship added");
        //Game.get().then(function() {
          //console.log("got updated game");
        //});
      });
    }
  }
}]);
