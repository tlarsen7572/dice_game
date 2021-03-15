import 'dart:convert';

Response textToResponse(String responseText) {
  var responseObj = json.decode(responseText);
  return Response.fromJson(responseObj);
}

class Response {
  Response(this.error, this.gameState);
  final String error;
  final GameState gameState;

  Response.fromJson(Map<String, dynamic> json) : error=json['Error'], gameState=json['GameState'] == null ? null : GameState.fromJson(json['GameState']);
}

class GameState {
  GameState(this.activeGame, this.actionLinks);
  final Game activeGame;
  final Actions actionLinks;

  GameState.fromJson(Map<String, dynamic> json) : activeGame=json['ActiveGame'] == null ? null : Game.fromJson(json['ActiveGame']), actionLinks=Actions.fromJson(json['ActionLinks']);
}

class Game {
  Game(this.winningScore, this.currentScore, this.turns, this.activeTurn);
  final int winningScore;
  final int currentScore;
  final List<int> turns;
  final Turn activeTurn;
  Game.fromJson(Map<String, dynamic> json) : winningScore=json['WinningScore'], currentScore=json['CurrentScore'], turns=(json['Turns'] as List<dynamic>).cast<int>(), activeTurn=Turn.fromJson(json['ActiveTurn']);
}

class Turn {
  Turn(this.score, this.lastRoll, this.lastScoringDice);
  final int score;
  final List<int> lastRoll;
  final List<int> lastScoringDice;

  Turn.fromJson(Map<String, dynamic> json) : score=json['Score'], lastRoll=(json['LastRoll'] as List<dynamic>).cast<int>(), lastScoringDice=(json['LastScoringDice'] as List<dynamic>).cast<int>();
}

class Actions {
  Actions(this.newGameLink, this.rollLink, this.newTurnLink);
  final Action newGameLink;
  final Action rollLink;
  final Action newTurnLink;

  Actions.fromJson(Map<String, dynamic> json) : newGameLink=json['0'] == null ? null : Action.fromJson(json['0']), rollLink=json['1'] == null ? null : Action.fromJson(json['1']), newTurnLink=json['2'] == null ? null : Action.fromJson(json['2']);
}

class Action {
  Action(this.token, this.url, this.method);
  final int token;
  final String url;
  final String method;

  Action.fromJson(Map<String, dynamic> json) : token=json['Token'], url=json['Url'], method=json['Method'];
}
