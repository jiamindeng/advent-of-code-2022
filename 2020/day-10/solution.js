const fs = require("fs");

let input = fs
  .readFileSync("./input.txt", "utf-8")
  .trim()
  .split(/\r\n/)
  .map((number) => {
    return parseInt(number);
  });
input.unshift(0, Math.max(...input) + 3);
input.sort((a, b) => a - b);

let differences = [];

for (let i = 1; i < input.length; i++) {
  differences.push(input[i] - input[i - 1]);
}

console.log(
  `Part One: ${
    differences.filter((number) => number === 1).length *
    differences.filter((number) => number === 3).length
  }`
);

const findCombos = (input) => {
  let combos = Array.from({ length: input.length }, () => 0);
  // Index of combos represents index of sorted input element
  combos[0] = 1;
  for (let i = 0; i < input.length; i++) {
    for (let j = i + 1; j <= i + 3 && j < input.length; j++) {
      if (input[j] - input[i] <= 3) {
        combos[j] += combos[i];
        if (input[j] - input[i] === 3) {
          break;
        }
      }
    }
  }
  return combos;
};

console.log(`Part Two: ${findCombos(input)[input.length - 1]}`);
