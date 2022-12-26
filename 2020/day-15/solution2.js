const fs = require('fs');

let input = fs
  .readFileSync('./input.txt', 'utf-8')
  .split(/,/)
  .map((num) => parseInt(num));

let turns = new Map();

let iterator = input.length;
let lastNum = input[input.length - 1];

let start = new Date().getTime();

input.forEach((num, index) => {
  turns.set(num, [index, -1]);
  //Object.assign(turns, { [num]: { recent: index, lastSeen: -1 } });
});

while (iterator < 30000000) {
  let recent = turns.get(lastNum)[0];
  let lastSeen = turns.get(lastNum)[1];
  // let recent = turns[lastNum].recent;
  // let lastSeen = turns[lastNum].lastSeen;

  if (lastSeen === -1) {
    lastNum = 0;
    turns.set(0, [iterator, turns.get(0) ? turns.get(0)[0] : -1]);
    //  turns[0] = { recent: iterator, lastSeen: turns[0] ? turns[0].recent : -1 };
  } else {
    let difference = recent - lastSeen;
    lastNum = difference;
    turns.set(difference, [
      iterator,
      turns.get(difference) ? turns.get(difference)[0] : -1,
    ]);
  }
  iterator++;
}
let end = new Date().getTime();

console.log(`Time difference: ${end - start}`);
console.log(`Part Two: ${lastNum}`);
