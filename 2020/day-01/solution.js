const fs = require("fs");

let numbers = fs
  .readFileSync("./input.txt", "utf-8")
  .split(/\r\n/)
  .map((num) => {
    return parseInt(num);
  });

// Why didn't I just sort this?
const solution1 = (numbers, target) => {
  let set = new Set();
  for (let i = 0; i < numbers.length; i++) {
    if (set.has(target - numbers[i])) {
      return numbers[i] * (target - numbers[i]);
    }
    set.add(numbers[i]);
  }
};

// Good question
const solution2 = (numbers, target) => {
  numbers = numbers.sort((a, b) => a - b);
  for (let i = 0; i < numbers.length; i++) {
    let left = i + 1;
    let right = numbers.length - 1;
    while (left < right) {
      if (numbers[i] + numbers[left] + numbers[right] == target) {
        return numbers[i] * numbers[left] * numbers[right];
      } else if (numbers[i] + numbers[left] + numbers[right] < target) {
        left++;
      } else {
        right--;
      }
    }
  }
  return -1;
};

console.log(`Part One: ${solution1(numbers, 2020)}`);

console.log(`Part Two: ${solution2(numbers, 2020)}`);
