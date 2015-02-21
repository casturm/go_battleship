console.log("define Game module");
angular.module('myApp.game', [])

.factory('Game', ['$http', function ($http) {
  console.log("create Game factory");
  var current;
  var players = []
  var factory = {};
  factory.current = function() {
    return current;
  };
  factory.create = function() {
    console.log('create game');
    return $http.post('api/game').then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      current = resp.data;
      size = current["Size"];
      players[0] = []
      players[1] = []
      players[0][0] = factory._rows(1,size,false)
      players[0][1] = factory._rows(0,size,true)
      players[1][0] = factory._rows(0,size,false)
      players[1][1] = factory._rows(1,size,true)
    });
  };
  factory.turn = function(player,x,y) {
    var turn = {Player: player, X: x, Y: y}
    console.log("turn: " + angular.toJson(turn));
    return $http.post('api/turn', turn).then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      players[(player + 1) % 2][0][y].cells[x].style = resp.data
      players[(player + 1) % 2][1][y].cells[x].style = resp.data
    });
  };
  factory.rows = function(player,index) {
    console.log("get rows " + player + " " + index);
    return players[player][index];
  };
  factory.isBoat = function(player,x,y) {
    var ships = current["Players"][player]["Ships"]
    for (i=0; i<ships.length; i++) {
      var ship = ships[i];
      var location = ship["Location"]
      for (p=0; p<location.length; p++) {
        if (location[p]["X"] == x && location[p]["Y"] == y) {
          //console.log("isBoat " + player + " " + x + " " + y + " " + angular.toJson(ships) + " " + ships.length)
          //console.log("p " + p + " " + angular.toJson(location[p]));
          //console.log("true");
          return true;
        }
      }
    }
    return false;
  };
  factory._rows = function(player,size,showBoats) {
    console.log('create rows ' + player + ' ' + showBoats);
    var rows = new Array(size);
    for (y=0; y<size; y++) {
      row = {}
      row.cells = []
      rows[y] = row
      for (x=0; x<size; x++) {
        newCell = {}
        row.cells[x] = newCell
        newCell.id = player + "_" + x + "_" + y
        newCell.style = "empty"
        if (showBoats && this.isBoat(player,x,y)) {
          newCell.boat = "B"
        }
        else {
          newCell.boat = ""
        }
        //console.log("new cell " + angular.toJson(newCell.id))
      }
    };
    return rows;
  };
  return factory;
}]);
