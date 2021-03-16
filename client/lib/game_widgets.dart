import 'package:dice_game_ui/bloc.dart';
import 'package:dice_game_ui/app_state.dart';
import 'package:flutter/material.dart';
import 'package:dice_game_ui/responses.dart' as r;

class GameWidgets extends StatelessWidget {
  Widget build(BuildContext context) {
    return Row(
      children: [
        Expanded(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              SizedBox(height: 10),
              RolledDice(),
              SizedBox(height: 10),
              RollScore(),
              SizedBox(height: 10),
              Row(
                children: [
                  SizedBox(
                    width: 200,
                    height: 70,
                    child: RollButton(),
                  ),
                  SizedBox(width: 10),
                  SizedBox(
                    width: 200,
                    height: 70,
                    child: NewTurnButton(),
                  ),
                ],
              ),
            ],
          ),
        ),
        SizedBox(
          width: 200,
          child: Column(
            children: [
              ScoreTitle(),
              SizedBox(height: 10),
              Expanded(
                child: Turns()
              ),
            ],
          ),
        ),
      ],
    );
  }
}

class ScoreTitle extends StatelessWidget {
  Widget build(BuildContext context) {
    var appState = BlocProvider.of<AppState>(context);
    return StreamBuilder<r.Game>(
      stream: appState.game,
      builder: (_, game){
        if (game.hasData && game.data != null) {
          var totalScore = game.data.currentScore;
          var winningScore = game.data.winningScore;
          return Text("$totalScore/$winningScore", style: TextStyle(fontSize: 20));
        }
        return Container();
      },
    );
  }
}

class Turns extends StatelessWidget {
  Widget build(BuildContext context) {
    var appState = BlocProvider.of<AppState>(context);
    return StreamBuilder<r.Game>(
      stream: appState.game,
      builder: (_, game) {
        List<Widget> children = [];
        if (game.hasData && game.data != null) {
          game.data.turns.forEach((e)=>children.insert(0, Text(e.toString(), textAlign: TextAlign.end)));
        }
        return ListView(
          children: children,
        );
      },
    );
  }
}

class RolledDice extends StatelessWidget {
  Widget build(BuildContext context) {
    var appState = BlocProvider.of<AppState>(context);
    return StreamBuilder<r.Turn>(
      stream: appState.currentTurn,
      builder: (_, turnSnapshot) {
        if (!turnSnapshot.hasData || turnSnapshot.data == null) {
          return Container();
        }
        var turn = turnSnapshot.data;
        List<Widget> dice = [];
        for (var index = 0; index < turn.lastRoll.length; index++) {
          var isScoring = turn.lastScoringDice.contains(index);
          dice.add(Dice(turn.lastRoll[index], isScoring));
        }
        return Row(
          children: dice,
        );
      },
    );
  }
}

class Dice extends StatelessWidget {
  Dice(this.value, this.isScoring);
  final int value;
  final bool isScoring;

  Widget build(BuildContext context) {
    return SizedBox(
      height: 80,
      width: 80,
      child: Card(
        child: Image.asset("images/dice_$value.png"),
        color: isScoring ? Colors.amberAccent : Colors.white,
      ),
    );
  }
}

class RollScore extends StatelessWidget {
  Widget build(BuildContext context) {
    var appState = BlocProvider.of<AppState>(context);
    return StreamBuilder<r.Turn>(
      stream: appState.currentTurn,
      builder: (_, turnSnapshot) {
        if (!turnSnapshot.hasData || turnSnapshot.data == null) {
          return Container();
        }
        var turn = turnSnapshot.data;
        return Text("Total score for turn: ${turn.score}", style: TextStyle(fontSize: 20));
      },
    );
  }
}

class RollButton extends StatelessWidget {
  Widget build(BuildContext context) {
    var appState = BlocProvider.of<AppState>(context);
    return StreamBuilder<r.Actions>(
      stream: appState.actions,
      builder: (_, snapshot) {
        if (!snapshot.hasData || snapshot.data == null) {
          return Container();
        }
        var actions = snapshot.data;
        if (actions.rollLink == null) {
          return Container();
        }
        return ElevatedButton(
          child: Text("Roll", style: TextStyle(fontSize: 20)),
          onPressed: ()=>appState.sendAction(actions.rollLink.url),
        );
      },
    );
  }
}

class NewTurnButton extends StatelessWidget {
  Widget build(BuildContext context) {
    var appState = BlocProvider.of<AppState>(context);
    return StreamBuilder<r.Actions>(
      stream: appState.actions,
      builder: (_, snapshot) {
        if (!snapshot.hasData || snapshot.data == null) {
          return Container();
        }
        var actions = snapshot.data;
        if (actions.newTurnLink == null) {
          return Container();
        }
        var text = "Bank score";
        if (actions.rollLink == null) {
          text = "New turn";
        }
        return ElevatedButton(
          child: Text(text, style: TextStyle(fontSize: 20)),
          onPressed: ()=>appState.sendAction(actions.newTurnLink.url),
        );
      },
    );
  }
}