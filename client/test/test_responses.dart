import 'package:dice_game_ui/responses.dart';
import 'package:flutter_test/flutter_test.dart';

main(){
  test("new game response", (){
    var responseText = """{"Error":"","GameState":{"ActiveGame":null,"ActionLinks":{"0":{"Token":6663395697926291085,"Url":"localhost:35015/NewGame?token=6663395697926291085\u0026winningScore={WinningScore}","Method":"POST"}}}}""";
    var response = textToResponse(responseText);
    print(response);
  });

  test("full response", (){
    var responseText = """{"Error":"","GameState":{"ActiveGame":{"WinningScore":10000,"CurrentScore":0,"Turns":[],"ActiveTurn":{"Score":0,"LastRoll":[],"LastScoringDice":[]}},"ActionLinks":{"0":{"Token":8268648633486974450,"Url":"localhost:35015/NewGame?token=8268648633486974450\u0026winningScore={WinningScore}","Method":"POST"},"1":{"Token":1565024250055940264,"Url":"localhost:35015/Roll?token=1565024250055940264","Method":"POST"}}}}""";
    var response = textToResponse(responseText);
    print(response);
  });

  test("error only", (){
    var responseText = """{"Error":"Something"}""";
    var response = textToResponse(responseText);
    print(response);
  });
}