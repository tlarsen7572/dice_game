import 'package:dice_game_ui/bloc.dart';
import 'package:dice_game_ui/responses.dart';
import 'package:rxdart/rxdart.dart' as rx;
import 'package:http/http.dart' as http;

class AppState extends BlocState {
  AppState();

  var _error = rx.BehaviorSubject<String>.seeded('');
  Stream get error => _error.stream;
  var _actions = rx.BehaviorSubject<Actions>.seeded(null);
  Stream get actions => _actions.stream;
  var _game = rx.BehaviorSubject<Game>.seeded(null);
  Stream get game => _game.stream;
  var _turns = rx.BehaviorSubject<List<int>>.seeded(null);
  Stream get turns => _turns.stream;
  var _currentTurn = rx.BehaviorSubject<Turn>.seeded(null);
  Stream get currentTurn => _currentTurn.stream;

  void notifyNewResponse(Response response) {
    _error.add(response.error);
    if (response.gameState == null) {
      _actions.add(null);
      _game.add(null);
      _turns.add(null);
      _currentTurn.add(null);
      return;
    }

    _actions.add(response.gameState.actionLinks);
    var game = response.gameState.activeGame;
    _game.add(game);
    if (game != null) {
      _turns.add(game.turns);
      _currentTurn.add(game.activeTurn);
    } else {
      _turns.add([]);
      _currentTurn.add(null);
    }
  }

  Future sendAction(String url) async {
    var request = await http.post(Uri.parse(url));
    var response = textToResponse(request.body);
    notifyNewResponse(response);
  }

  void dispose() {
    _error.close();
    _actions.close();
    _game.close();
    _turns.close();
    _currentTurn.close();
  }

  Future initialize() async {}
}