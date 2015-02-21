'use strict';

angular.module('myApp.view3', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/view3', {
    templateUrl: 'view3/view3.html',
    controller: 'View3Ctrl'
  });
}])

.controller('View3Ctrl', ['$scope', '$location', 'Game', function($scope, $location, Game) {

  $scope.player1Rows = Game.rows(0,0);
  $scope.player2Rows = Game.rows(1,1);

  $scope.turn = function(x,y) {
    console.log("take turn 1");
    Game.turn(1,x,y).then(function() {
      console.log("turn complete 1")
      //console.log(Game.current)
      //$location.path("/view2");
    });
  }
}]);
