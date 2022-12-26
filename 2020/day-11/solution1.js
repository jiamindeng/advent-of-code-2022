const fs = require("fs");

let input = fs
  .readFileSync("./input.txt", "utf-8")
  .split(/\r\n/)
  .map((line) => line.split(""));

const numAdjacent = (y, x, input, target) => {
  let count = 0;
  const deltaY = [0, 0, 1, 1, 1, -1, -1, -1];
  const deltaX = [1, -1, 0, 1, -1, 0, 1, -1];

  for (let i = 0; i < deltaY.length; i++) {
    let newY = y + deltaY[i];
    let newX = x + deltaX[i];

    if (
      newY < input.length &&
      newY >= 0 &&
      newX >= 0 &&
      newX < input[0].length
    ) {
      if (input[newY][newX] === target) {
        count++;
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
        (input[y][x] === "L" && numAdjacent(y, x, input, "#") === 0) ||
        (input[y][x] === "#" && numAdjacent(y, x, input, "#") < 4)
      ) {
        newRow.push("#");
      } else if (input[y][x] === ".") {
        newRow.push(".");
      } else if (input[y][x] === "#" && numAdjacent(y, x, input, "#") >= 4) {
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

console.log(
  `Part One: ${current.flat().filter((seat) => seat === "#").length}`
);
