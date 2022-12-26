const fs = require("fs");

let input = fs
  .readFileSync("./input.txt", "utf-8")
  .split(/\r\n/)
  .map((number) => {
    return parseInt(number);
  });

const findOutlier = (input, preamble) => {
  let traversed = [];
  let outlier = -1;
  for (let [index, number] of input.entries()) {
    traversed.push(number);
    if (
      index >= preamble &&
      !canAdd(number, traversed.slice(index - preamble, index))
    ) {
      outlier = number;
      break;
    }
  }
  return outlier;
};

const canAdd = (target, traversed) => {
  let numbers = Array.from(traversed);
  let set = new Set();
  for (let i = 0; i < numbers.length; i++) {
    if (set.has(target - numbers[i])) {
      return true;
    }
    set.add(numbers[i]);
  }
  return false;
};

console.log(`Part One: ${findOutlier(input, 25)}`);
const sum = (array) => array.reduce((acc, current) => acc + current);

const findContiguousAdd = (target, input) => {
  for (let i = 0; i < input.length; i++) {
    for (let j = i + 1; j < input.length; j++) {
      let currentSum = sum(input.slice(i, j));
      if (currentSum === target) {
        return Math.min(...input.slice(i, j)) + Math.max(...input.slice(i, j));
      } else if (sum >= target) {
        i++;
      }
    }
  }
};

console.log(`Part Two: ${findContiguousAdd(findOutlier(input, 25), input)}`);
