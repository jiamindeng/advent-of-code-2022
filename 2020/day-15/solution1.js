const fs = require('fs');

let input = fs
  .readFileSync('./input.txt', 'utf-8')
  .split(/,/)
  .map((num) => parseInt(num));

let iterator = input.length;
let turns = [...input];
let lastNum;

while (iterator <= 2020) {
  lastNum = turns[turns.length - 1];
  let last = turns.lastIndexOf(lastNum);
  let numOccurrences = turns.filter((num) => num === lastNum).length;
  if (numOccurrences === 1) {
    turns.push(0);
  } else if (last > -1) {
    let difference = last - turns.slice(0, last).lastIndexOf(lastNum);
    turns.push(difference);
  }
  iterator++;
}

console.log(`Part One: ${lastNum}`);
