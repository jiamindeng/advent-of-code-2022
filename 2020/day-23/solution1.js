const fs = require('fs');

let cups = fs
  .readFileSync('./input.txt', 'utf-8')
  .split('')
  .map((num) => parseInt(num));

let currentIndex = 0;
let count = 100;

while (count > 0) {
  const start = (currentIndex + 1) % cups.length;
  const stop = (currentIndex + 4) % cups.length;

  const currentCup = cups[currentIndex];

  const removedCups =
    start > stop
      ? [...cups.splice(start), ...cups.splice(0, stop)]
      : cups.splice(start, 3);

  let destinationCup = currentCup - 1;

  while (removedCups.includes(destinationCup) || destinationCup === 0) {
    destinationCup--;
    if (destinationCup < Math.min(...cups)) {
      destinationCup = Math.max(...cups);
      break;
    }
  }

  cups.splice(cups.indexOf(destinationCup) + 1, 0, ...removedCups);
  currentIndex = (cups.indexOf(currentCup) + 1) % cups.length;

  count--;
}

const result = [
  ...cups.slice(cups.indexOf(1)),
  ...cups.slice(0, cups.indexOf(1)),
];

result.shift();

console.log(`Part One: ${result.join('')}`);
