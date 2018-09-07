# 1-4-24

1-4-24 Dice Game in Python (with tests and type checking!)

# Notes

## WebSockets / API

- Everything gets rendered in React
- Entire state of entire 1-4-24 game needs to be in a single JSON thing

(sent from back-end to front-end)

```go
sendUpdatedGameState(game_id)

// load from database
// send over websockets
```

```javascript
{
  current_pot: 1,
  bet_size: 1,
  current_player: "austin",

  gameplays: [
    {
      player: { name: "austin", balance: 420 },
      qualifiers: [1, 4]
    },
    {
      ...
    }
  ],

}
```

(sent from front-end to back-end)

```go
receiveTurn(turnData) # -> sendUpdatedGameState()

// load game from database, initialize new Game, Player, etc..
// update game state
// write to database
// send over websockets

```

```javascript
{
  game_id: 1234,
  player: "austin",
  keeps: [4, 2]
}
```
