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
  $scope.start = function() {
    Game.create().then(function() {
      $location.path( "/view2" );
    });
  }
}]);
