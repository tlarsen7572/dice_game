### March 9

Started looking into rules of the game to begin understanding the complexity required. Biggest challenge will probably be calculating multiple scoring dice on any given roll. Also did a quick refresher on REST APIs as most of my personal work is loosely REST-ish and I could probably benefit from a better understanding. I found an interesting concept about the server returning the URLs for client applications to subsequently call.  This makes the API self-discoverable and, I imagine, a bit more fluid.  I like this idea and want to try it out.

### March 10

I had a few minutes in the morning. I'm still not decided on the architecture, but wanted to get some coding done. I know there are behaviors that will need to be coded no matter the server architecture, so I spent my morning session starting with rolling dice and making sure it's pseudo-random.

Later while I was driving my mind went to the REST API. My idea for the REST API is to start by loading the web page. The first step after loading the page is to GET the current status of the game. Each response might contain the following items:

- Current status
- Actions with URLs and methods
- Current game state containing score history and current turn history

Each action called from the client, using the provided URL, returns the same response structure. In this way, the server borrows from functional programming ideas and all of the state exists on the server rather than the UI.  The UI just has to render the current state and send actions to the server.

The server is responsible for parsing client requests, game logic, and keeping track of valid actions.

Now that this structure is decided I can start coding the different pieces of the server. I plan on coding the rules of the game next. After that will be building the web server and designing the ability to process and generate actions. To wrap everything up I will build a simple UI that can be served via HTML from the web server. I will use Flutter for the UI.

### March 11

Found some time tonight to work on the scoring algorithm. Started with handling ones and am beginning to implement sets of three. Still not sure on the right structure for sets of three. There is common logic which suggests an opportunity for abstraction, but it has not come to me yet. I will keep implementing the algorithm and refactoring and see if the abstraction pops into focus as I move forward.

### March 12

Spent about an hour wrapping up the scoring logic. All of my tests pass and I cannot think of any further edge cases at this time. Next up will be coding the turn logic, which should mean building up a game state object. The game object should only concern itself with turns. I want to keep the REST API logic completely separate, which means generating the self-discovery URLs will be separated out from the game.

### March 14

Struggled to get the game-level logic started. I ended up putting too much into the Game object and finally realized there was a Turn object wanting to be separately defined. Once I pulled out the turn logic into its own struct, Game and Turn started making more sense. I also made the mistake of trying to incorporate some of the API design ideas into these classes, which really muddied things up. The API should depend on the game and turn, not the other way around.

### March 15

I finally got some time to focus on this project, rather than a few minutes here and there as in prior days. The web server is wrapped up and an initial working version is completed. I started building a basic UI in Flutter, because why build a REST API and not have something fun to test it with? The REST API is designed to always return the full game state in a JSON payload, so the UI in Flutter should essentially boil down to a bunch of StreamBuilders that update when a core state object is updated with new JSON. There should be very little logic in the UI. I did not get as far as I wanted with the UI, but am hoping to wrap it up early tomorrow.

### March 16

Success! The UI is finished and in a working state. This was a useful exercise because the UI quickly identified 2 errors with my server code that I was able to put tests around and fix. Sometimes there is no substitute for seeing the output of your server logic visually. The dice game is PoC finished at this point.
