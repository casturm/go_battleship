console.log("define Game module");
angular.module('myApp.game', [])

.factory('Game', ['$http', function ($http) {
  console.log("create Game factory");
  var current;
  var players = []
  var thisPlayer;
  var factory = {};
  factory.current = function() {
    return current;
  };
  factory.addShip = function(ship) {
    ship.Player = thisPlayer.Id
    ship.Size = parseInt(ship.Size)
    ship.LocationX = parseInt(ship.LocationX)
    ship.LocationY = parseInt(ship.LocationY)
    return $http.post('api/ship', ship).then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      if (thisPlayer.Ships == null) {
        thisPlayer.Ships = []
      }
      thisPlayer.Ships.push(resp.data)
      refresh(1,current.size,true);
    });
  }
  factory.players = function() {
    return players;
  };
  factory.player = function(name) {
    console.log('player ' + name);
    player= {Name: name}
    return $http.post('api/player', player).then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      thisPlayer = resp.data;
    });
  };
  factory.getPlayers = function(name) {
    console.log('getPlayers');
    return $http.get('api/player').then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      players = resp.data;
    });
  };
  factory.createGame = function(opponentId) {
    console.log('create game with thisPlayer ' + thisPlayer.Id + " and " + opponentId);
    var gamePlayers = {Player1: thisPlayer.Id, Player2: opponentId}
    return $http.post('api/game', gamePlayers).then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      setupGame(resp.data);
    });
  };
  setupGame = function(data) {
    current = data;
    size = current["Size"];
    current.gameOn = false

    // primary player view
    players[0][0] = setupRows(1,size,false)
    players[0][1] = setupRows(0,size,true)

    // for the other player's view
    players[1][0] = setupRows(0,size,false)
    players[1][1] = setupRows(1,size,true)
  };
  factory.get = function() {
    console.log('get game');
    return $http.get('api/game').then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      setupGame(resp.data);
    });
  };
  factory.turn = function(player,x,y) {
    var turn = {gameId: current.Id, Player: player, X: x, Y: y}
    console.log("turn: " + angular.toJson(turn));
    return $http.post('api/turn', turn).then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      if (resp.data == "Not Your Turn") {
        alert("it's not your turn");
      } else {
        players[(player + 1) % 2][0][y].cells[x].style = resp.data
        players[(player + 1) % 2][1][y].cells[x].style = resp.data
      }
    });
  };
  factory.rows = function(player,index) {
    console.log("get rows " + player + " " + index);
    return players[player][index];
  };
  isBoat = function(player,x,y) {
    var ships = current["Players"][player]["Ships"]
    if (ships != null) {
      for (i=0; i<ships.length; i++) {
        var ship = ships[i];
        var points = ship["Location"]
        for (p=0; p<points.length; p++) {
          if (points[p]["X"] == x && points[p]["Y"] == y) {
            return true;
          }
        }
      }
    }
    return false;
  };
  isHit = function(player,x,y) {
    var ships = current["Players"][player]["Ships"]
    if (ships != null) {
      for (i=0; i<ships.length; i++) {
        var ship = ships[i];
        var points = ship["Hits"]
        for (p=0; p<points.length; p++) {
          if (points[p]["X"] == x && points[p]["Y"] == y) {
            return true;
          }
        }
      }
    }
    return false;
  };
  isMiss = function(player,x,y) {
    var points = current["Players"][player]["Misses"]
    if (points != null) {
      for (p=0; p<points.length; p++) {
        if (points[p]["X"] == x && points[p]["Y"] == y) {
          return true;
        }
      }
    }
    return false;
  };
  refresh = function(player,size,showBoats) {
    for (y=0; y<size; y++) {
      for (x=0; x<size; x++) {
        if (current.gameOn) {
          cell = rows(player,1)[y].cells[x]
          if (this.isHit(player,x,y)) {
            cell.style = "hit"
          }
          else if (this.isMiss(player,x,y)) {
            cell.style = "miss"
          }

          if (showBoats && this.isBoat(player,x,y)) {
            cell.boat = "B"
          }
          else {
            cell.boat = ""
          }
        }
      }
    };
  };
  setupRows = function(player,size,showBoats) {
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
        //if (current.gameOn) {
          if (this.isHit(player,x,y)) {
            newCell.style = "hit"
          }
          else if (this.isMiss(player,x,y)) {
            newCell.style = "miss"
          }

          if (showBoats && this.isBoat(player,x,y)) {
            newCell.boat = "B"
          }
          else {
            newCell.boat = ""
          }
          //console.log("new cell " + angular.toJson(newCell.id))
        //}
      }
    };
    return rows;
  };
  return factory;
}]);
