const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8").split(/\r\n/);

const toSeat = (seatId) => {
  let minRow = 0;
  let maxRow = 128;
  let minSeat = 0;
  let maxSeat = 8;

  seatId.split("").forEach((letter) => {
    switch (letter) {
      case "B":
        minRow = (maxRow - minRow) / 2 + minRow;
        break;
      case "F":
        maxRow = (maxRow - minRow) / 2 + minRow;
        break;
      case "L":
        maxSeat = (maxSeat - minSeat) / 2 + minSeat;
        break;
      case "R":
        minSeat = (maxSeat - minSeat) / 2 + minSeat;
        break;
      default:
        throw new Error("Letter not found.");
    }
  });
  return minRow * 8 + minSeat;
};

let seatIds = input.map((seatId) => toSeat(seatId));

let allSeats = [];
for (let index = 100; index < 924; index++) {
  allSeats.push(index);
}

let all = new Set(allSeats);
let actual = new Set(seatIds);
let missing = new Set([...all].filter((x) => !actual.has(x)));

console.log(`Part One: ${Math.max(...seatIds)}`);
console.log(`Part Two: ${[...missing].join("")}`);
