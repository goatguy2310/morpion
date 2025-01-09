# Tictactoe Codeforces
A discord bot built for hosting tictactoe games between two players making moves by solving codeforces problems. Inspired by the bot NeoTLE, I wanted to write my own version of codeforces duel in the form of tictactoe games.

# How the game works
By challenging the other player alongside giving your handles, the bot will randomly choose 9 problems from the competitive programming website codeforces that matches with the criteria that you give it. (i.e. 2000 +constructive-algorithms +dp ~flows) 

![image](https://github.com/user-attachments/assets/a20e002f-d231-4aa4-a278-6476d0d59a0b)

Then, by solving one of the 9 problems, the corresponding spot on the tictactoe board will be marked with your symbol (X or O). The game will end when three Xs or Os appear in the same row, same column or the same diagonal.

![image](https://github.com/user-attachments/assets/91d36a7a-7a11-45e2-8e6f-49cc910fbb80)

# Usage
Clone the repository. Make sure that you put your discord bot token in a `.env` file with the field name `DISCORD_TOKEN`.

Make sure that you have `go1.23.4` or more. Then, go in the directory and just run

```
go run .
```
