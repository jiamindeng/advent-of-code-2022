const fs = require('fs');

let input = fs
  .readFileSync('./input.txt', 'utf-8')
  .split(/\r\n\r\n/)
  .map((deck) =>
    deck
      .split(/\r\n/)
      .filter((string) => !string.match(/Player/))
      .map((num) => parseInt(num))
  );

let deckOne = input[0],
  deckTwo = input[1];

const transferCards = (winner, loser) => {
  let winnerCard = winner.shift();
  let loserCard = loser.shift();
  winner.push(winnerCard, loserCard);
};

const getWinner = (playerOne, playerTwo) => {
  let deckOne = [...playerOne];
  let deckTwo = [...playerTwo];

  let previousDecks = new Set();

  while (deckOne.length > 0 && deckTwo.length > 0) {
    let tempOne = [...deckOne],
      tempTwo = [...deckTwo];

    if (previousDecks.has(JSON.stringify({ deckOne, deckTwo })))
      return [deckOne, deckTwo, 'repeat'];

    previousDecks.add(JSON.stringify({ deckOne: tempOne, deckTwo: tempTwo }));
    if (deckOne.length - 1 >= deckOne[0] && deckTwo.length - 1 >= deckTwo[0]) {
      let [one, two, repeat] = getWinner(
        deckOne.slice(1, deckOne[0] + 1),
        deckTwo.slice(1, deckTwo[0] + 1)
      );

      if (repeat === 'repeat') {
        transferCards(deckOne, deckTwo);
      } else {
        two.length === 0
          ? transferCards(deckOne, deckTwo)
          : transferCards(deckTwo, deckOne);
      }
    } else {
      deckOne[0] > deckTwo[0]
        ? transferCards(deckOne, deckTwo)
        : transferCards(deckTwo, deckOne);
    }
  }

  return [deckOne, deckTwo];
};

const calculateScore = (deck) => {
  let multiplier = 1;
  let score = 0;
  deck.reverse().forEach((cardNum) => {
    score += cardNum * multiplier;
    multiplier++;
  });
  return score;
};

[playerOne, playerTwo] = getWinner(deckOne, deckTwo);

console.log(`Part One: ${calculateScore(playerOne)}`);
