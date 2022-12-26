const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8").split("\r\n");

const getNumTrees = (input) => {
  let i = 0;
  let count = 0;
  input.forEach((line) => {
    if (line[i % line.length] === "#") {
      count++;
    }
    i += 3;
  });
  return count;
};

// I might've given up on optimizing stuff
const getNumTreesButAnnoying = (input, right, down) => {
  let i = 0;
  let count = 0;

  for (let index = 0; index < input.length; index += down) {
    let line = input[index];
    if (line[i % line.length] === "#") {
      count++;
    }
    i += right;
  }

  return count;
};

let patterns = [
  { right: 1, down: 1 },
  { right: 3, down: 1 },
  { right: 5, down: 1 },
  { right: 7, down: 1 },
  { right: 1, down: 2 },
];

let numTrees = patterns
  .map((pattern) => getNumTreesButAnnoying(input, pattern.right, pattern.down))
  .reduce((accumulator, currentNumTrees) => accumulator * currentNumTrees);

console.log(`Part One: ${getNumTrees(input, 3, 1)}`);
console.log(`Part Two: ${numTrees}`);
