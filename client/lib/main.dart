import 'package:dice_game_ui/bloc.dart';
import 'package:dice_game_ui/app_state.dart';
import 'package:dice_game_ui/responses.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:dice_game_ui/responses.dart' as r;
import 'package:http/http.dart' as http;

void main() async {
  var baseUrl = Uri.base.toString();
  if (baseUrl[-1] != '/') {
    baseUrl += "/";
  }
  var state = AppState();
  var request = await http.get(Uri.parse('${baseUrl}GameStatus'));
  var response = textToResponse(request.body);
  state.notifyNewResponse(response);

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
                      return IconButton(icon: Icon(Icons.add), onPressed: () async {
                        var url = actions.data.newGameLink.url;
                        url = url.replaceAll("{WinningScore}", '10000');
                        await appState.sendAction(url);
                      });
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
              return Container();
            },
          ),
        ),
      ),
    );
  }
}
