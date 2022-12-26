const fs = require('fs');

let cups = fs
  .readFileSync('./input.txt', 'utf-8')
  .split('')
  .map((num) => parseInt(num));

let currentIndex = 0;
let count = 10000000;
let cupNum = Math.max(...cups);
let pad = Array.from({ length: 1000000 - cups.length }, () => ++cupNum);
cups = [...cups, ...pad];

// Source: https://medium.com/coding-at-dawn/the-fastest-way-to-find-minimum-and-maximum-values-in-an-array-in-javascript-2511115f8621
const forLoopMinMax = (array) => {
  let min = array[0],
    max = array[0];

  for (let i = 1; i < array.length; i++) {
    let value = array[i];
    min = value < min ? value : min;
    max = value > max ? value : max;
  }

  return [min, max];
};

const getMin = (array) => forLoopMinMax(array)[0];
const getMax = (array) => forLoopMinMax(array)[1];

// Reimplement later with linkedlist
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
    if (destinationCup < getMin(cups)) {
      destinationCup = getMax(cups);
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

console.log(`Part Two: ${result[0] * result[1]}`);
