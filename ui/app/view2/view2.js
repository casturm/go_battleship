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
    Game.turn(0,x,y).then(function() {
      console.log("turn complete 0")
      //console.log(Game.current)
      //$location.path("/view3");
    });
  }
}]);
