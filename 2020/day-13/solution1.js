const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8").split(/\r\n/);

[currentTime, buses] = input;

let busIds = buses
  .replace(/,x/g, "")
  .split(/,/)
  .map((busId) => parseInt(busId));

let multiples = busIds.map((id, index) => {
  return Math.ceil(currentTime / id) * id;
});

console.log(
  `Part One: ${
    (Math.min(...multiples) - currentTime) *
    busIds[multiples.indexOf(Math.min(...multiples))]
  }`
);
