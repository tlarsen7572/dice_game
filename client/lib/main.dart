import 'package:dice_game_ui/bloc.dart';
import 'package:dice_game_ui/app_state.dart';
import 'package:dice_game_ui/game_widgets.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:dice_game_ui/responses.dart' as r;
import 'package:http/http.dart' as http;

//String mock = """{"Error":"","GameState":{"ActiveGame":{"WinningScore":10000,"CurrentScore":150,"Turns":[50,0,100],"ActiveTurn":{"Score":100,"LastRoll":[1,2,2,3,3,4],"LastScoringDice":[0]}},"ActionLinks":{"0":{"Token":8268648633486974450,"Url":"localhost:35015/NewGame?token=8268648633486974450\u0026winningScore={WinningScore}","Method":"POST"},"1":{"Token":1565024250055940264,"Url":"localhost:35015/Roll?token=1565024250055940264","Method":"POST"},"2":{"Token":1565024250055940264,"Url":"localhost:35015/NewTurn?token=3465024250055940264","Method":"POST"}}}}""";
//String mock = """{"Error":"","GameState":{"ActiveGame":null,"ActionLinks":{"0":{"Token":6663395697926291085,"Url":"localhost:35015/NewGame?token=6663395697926291085\u0026winningScore={WinningScore}","Method":"POST"}}}}""";
String mock = "";

void main() async {
  var state = AppState();

  if (mock != '') {
    var response = r.textToResponse(mock);
    state.notifyNewResponse(response);
  } else {
    var baseUrl = Uri.base.toString();
    if (baseUrl[-1] != '/') {
      baseUrl += "/";
    }
    var request = await http.get(Uri.parse('${baseUrl}GameStatus'));
    var response = r.textToResponse(request.body);
    state.notifyNewResponse(response);
  }

  runApp(
    BlocProvider(
      bloc: state,
      child: MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Dice Game',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: MyHomePage(title: 'Dice Game'),
    );
  }
}

class MyHomePage extends StatelessWidget {
  MyHomePage({Key key, this.title}) : super(key: key);
  final String title;

  Widget build(BuildContext context) {
    var appState = BlocProvider.of<AppState>(context);
    return Scaffold(
      appBar: AppBar(
        title: Container(
          width: 800,
          child: Row(
            children: [
              Expanded(child: Text(title)),
              StreamBuilder<r.Actions>(
                stream: appState.actions,
                builder: (_, actions) {
                    if (actions.hasData && actions.data.newGameLink != null) {
                      return TextButton(
                        child: Text("New Game", style: TextStyle(color: Colors.white)),
                        onPressed: () async {
                          var url = actions.data.newGameLink.url;
                          url = url.replaceAll("{WinningScore}", '10000');
                          await appState.sendAction(url);
                        },
                      );
                    }
                    return Container();
                  }
              ),
            ],
          ),
        ),
      ),
      body: Center(child:
        Container(
          width: 800,
          child: StreamBuilder<r.Actions>(
            stream: appState.actions,
            builder: (_, actions) {
              if (!actions.hasData || actions.data == null) {
                return Center(child: Text("Loading..."));
              }
              return GameWidgets();
            },
          ),
        ),
      ),
    );
  }
}
