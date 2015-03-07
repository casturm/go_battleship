console.log("define Game module");
angular.module('myApp.game', [])

.factory('Game', ['$http', function ($http) {
  console.log("create Game factory");
  var current;
  var players = []
  var playersList = []
  var games = [];
  var thisPlayer;
  var factory = {};
  factory.current = function() {
    return current;
  };
  factory.getPlayer = function(player) {
    return players[player]
  };
  factory.getThisPlayer = function() {
    return thisPlayer;
  };
  factory.loadGame = function(gameId) {
    return $http.get('api/game/' + gameId).then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      setupGame(resp.data);
      refresh();
    });
  };
  factory.getGames = function() {
    return $http.get('api/games').then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      games = resp.data
    });
  };
  factory.games = function() {
    return games;
  };
  factory.addShip = function(ship) {
    ship.Player = thisPlayer.Id
    ship.Size = parseInt(ship.Size)
    ship.LocationX = parseInt(ship.LocationX)
    ship.LocationY = parseInt(ship.LocationY)
    return $http.post('api/ship', ship).then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      if (players[0].Ships == null) {
        players[0].Ships = []
      }
      players[0].Ships.push(resp.data)
      refresh();
    });
  }
  factory.newPlayer = function(name) {
    console.log('player ' + name);
    player= {Name: name}
    return $http.post('api/player', player).then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      thisPlayer = resp.data;
    });
  };
  factory.players = function() {
    return playersList;
  };
  factory.getPlayers = function(name) {
    console.log('getPlayers');
    return $http.get('api/players').then(function (resp) {
      console.log('  server response: ' + angular.toJson(resp.data));
      playersList = resp.data;
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
  setupGame = function(data) {
    current = data;
    if (thisPlayer.Id == current.Players[0].Id) {
      players = current.Players
    }
    else {
      players[0] = current.Players[1]
      players[1] = current.Players[0]
    }
    size = current.Size
    current.gameOn = false

    // primary player view
    players[0][0] = setupRows(1)
    players[0][1] = setupRows(0)

    // for the other player's view
    players[1][0] = setupRows(0)
    players[1][1] = setupRows(1)
  };
  isBoat = function(player,x,y) {
    var ships = players[player]["Ships"]
    if (ships != null) {
      for (i=0; i<ships.length; i++) {
        var ship = ships[i];
        var points = ship["Location"]
        for (var p=0; p<points.length; p++) {
          if (points[p]["X"] == x && points[p]["Y"] == y) {
            return true;
          }
        }
      }
    }
    return false;
  };
  isHit = function(player,x,y) {
    var ships = players[player]["Ships"]
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
    var points = players[player]["Misses"]
    if (points != null) {
      for (p=0; p<points.length; p++) {
        if (points[p]["X"] == x && points[p]["Y"] == y) {
          return true;
        }
      }
    }
    return false;
  };
  refresh = function() {
    size = current.Size;
    for (p=0; p<2; p++) {
      for (b=0; b<2; b++) {
        for (y=0; y<size; y++) {
          for (x=0; x<size; x++) {
            cell = factory.rows(p,b)[y].cells[x]
            if (current.gameOn) {
              if (this.isHit(p,x,y)) {
                cell.style = "hit"
              }
              else if (this.isMiss(p,x,y)) {
                cell.style = "miss"
              }
            }
            if (b == 1 && this.isBoat(p,x,y)) {
              cell.boat = "B"
            }
            else {
              cell.boat = ""
            }
          }
        }
      }
    };
  };
  setupRows = function(player) {
    size = current.Size;
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
      }
    };
    return rows;
  };
  return factory;
}]);
