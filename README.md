# Morpion Codeforces
A discord bot built with Go that allows two players to play tic-tac-toe by solving codeforces problems. Inspired by the bot NeoTLE, I wanted to create my own version of a codeforces duel bot in the form of tic-tac-toe games.

# How the game works
By challenging the other player alongside giving your handles, the bot will randomly choose 9 problems from the competitive programming platform Codeforces based on the the criteria specified (i.e. 2000 +constructive-algorithms +dp ~flows).

![image](https://github.com/user-attachments/assets/a20e002f-d231-4aa4-a278-6476d0d59a0b)

Then, by solving one of the 9 problems, the corresponding spot on the tictactoe board will be marked with your symbol (X or O). The game will end when three Xs or Os appear in the same row, same column or the same diagonal.

![image](https://github.com/user-attachments/assets/91d36a7a-7a11-45e2-8e6f-49cc910fbb80)

# Commands

- `help`: Display the help message
- `challenge` `@opponent` `your handle` `opponent's handle` `rating` `+tags` `~tags`: Challenge the `@opponent` to a tictactoe duel, with the given rating (leave empty for any rating), and criteria for tags included (+) and tags excluded (~)
- `accept`: Accept a challenge if you are being challenged
- `end`: End a challenge or an ongoing duel
- `update`: Update the current duel, which will update the board if the duelists solve more problems. Should be manually called from time to time.

# Usage
Clone the repository. Make sure that you put your discord bot token in a `.env` file with the field name `DISCORD_TOKEN`.

Make sure that you have `go1.23.4` or more. Then, go in the directory and just run

```
go run .
```
