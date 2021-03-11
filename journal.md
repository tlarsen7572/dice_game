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
