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

const getWinner = (deckOne, deckTwo) => {
  while (deckOne.length > 0 && deckTwo.length > 0) {
    deckOne[0] > deckTwo[0]
      ? transferCards(deckOne, deckTwo)
      : transferCards(deckTwo, deckOne);
  }
  return deckOne.length === 0 ? deckTwo : deckOne;
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

console.log(`Part One: ${calculateScore(getWinner(deckOne, deckTwo))}`);
