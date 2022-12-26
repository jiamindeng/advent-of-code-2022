const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8").split(/\r\n/);

let dirToDegrees = {
  N: 0,
  E: 90,
  S: 180,
  W: 270,
};
let degreesToDir = {
  0: "N",
  90: "E",
  180: "S",
  270: "W",
};

const move = {
  N: (current, distance) => {
    current.coord.y += distance;
    return current;
  },
  S: (current, distance) => {
    current.coord.y -= distance;
    return current;
  },
  E: (current, distance) => {
    current.coord.x += distance;
    return current;
  },
  W: (current, distance) => {
    current.coord.x -= distance;
    return current;
  },
  F: (current, distance) => {
    switch (current.dir) {
      case "E":
        return move["E"](current, distance);
      case "W":
        return move["W"](current, distance);
      case "N":
        return move["N"](current, distance);
      case "S":
        return move["S"](current, distance);
    }
  },
  R: (current, degrees) => {
    current.dir = degreesToDir[(dirToDegrees[current.dir] + degrees) % 360];
    return current;
  },
  L: (current, degrees) => {
    current.dir =
      degreesToDir[(360 + (dirToDegrees[current.dir] - degrees)) % 360];
    return current;
  },
};

let current = { coord: { x: 0, y: 0 }, dir: "E" };

input.forEach((nextMove) => {
  let distance = parseInt(nextMove.slice(1));
  move[nextMove[0]](current, distance);
});

console.log(
  `Part One: ${Math.abs(current.coord.x) + Math.abs(current.coord.y)}`
);
