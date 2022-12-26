const fs = require("fs");

let input = fs
  .readFileSync("./input.txt", "utf-8")
  .split(/\r\n/)
  .map((line) => line.split(""));

// AVERT THINE EYES
const numLineOfSight = (y, x, input, target) => {
  let count = 0;
  const deltaY = [0, 0, 1, 1, 1, -1, -1, -1];
  const deltaX = [1, -1, 0, 1, -1, 0, 1, -1];

  for (let i = 0; i < deltaY.length; i++) {
    loop: for (let j = 1; j < input[0].length; j++) {
      let newY = y + deltaY[i] * j;
      let newX = x + deltaX[i] * j;
      if (
        newY < input.length &&
        newY >= 0 &&
        newX >= 0 &&
        newX < input[0].length
      ) {
        if (input[newY][newX] === "L") {
          break loop;
        } else if (input[newY][newX] === target) {
          count++;
          break loop;
        }
      }
    }
  }
  return count;
};

const flipSeats = (input) => {
  let nextRound = [];
  for (let y = 0; y < input.length; y++) {
    let newRow = [];
    for (let x = 0; x < input[0].length; x++) {
      if (
        (input[y][x] === "L" && numLineOfSight(y, x, input, "#") === 0) ||
        (input[y][x] === "#" && numLineOfSight(y, x, input, "#") < 5)
      ) {
        newRow.push("#");
      } else if (input[y][x] === ".") {
        newRow.push(".");
      } else if (input[y][x] === "#" && numLineOfSight(y, x, input, "#") >= 5) {
        newRow.push("L");
      } else {
        newRow.push(input[y][x]);
      }
    }
    nextRound.push(newRow);
  }
  return nextRound;
};

let current = flipSeats(input);
let next = flipSeats(current);

while (JSON.stringify(current) !== JSON.stringify(next)) {
  current = next;
  next = flipSeats(current);
}

console.log(`Part Two ${current.flat().filter((seat) => seat === "#").length}`);
