const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8").split(/\r\n/);

const move = {
  N: (current, distance) => {
    current.waypoint.dy += distance;
    return current;
  },
  S: (current, distance) => {
    current.waypoint.dy -= distance;
    return current;
  },
  E: (current, distance) => {
    current.waypoint.dx += distance;
    return current;
  },
  W: (current, distance) => {
    current.waypoint.dx -= distance;
    return current;
  },
  F: (current, times) => {
    current.coord.x += current.waypoint.dx * times;
    current.coord.y += current.waypoint.dy * times;
    return current;
  },
  R: (current, degrees) => {
    let degreesToDir = {
      0: [current.waypoint.dx, current.waypoint.dy],
      90: [current.waypoint.dy, -current.waypoint.dx],
      180: [-current.waypoint.dx, -current.waypoint.dy],
      270: [-current.waypoint.dy, current.waypoint.dx],
    };
    current.waypoint.dx = degreesToDir[degrees % 360][0];
    current.waypoint.dy = degreesToDir[degrees % 360][1];
    return current;
  },
  L: (current, degrees) => {
    let degreesToDir = {
      0: [current.waypoint.dx, current.waypoint.dy],
      90: [-current.waypoint.dy, current.waypoint.dx],
      180: [-current.waypoint.dx, -current.waypoint.dy],
      270: [current.waypoint.dy, -current.waypoint.dx],
    };
    current.waypoint.dx = degreesToDir[degrees % 360][0];
    current.waypoint.dy = degreesToDir[degrees % 360][1];
    return current;
  },
};

let current = { coord: { x: 0, y: 0 }, dir: "E", waypoint: { dx: 10, dy: 1 } };

input.forEach((nextMove) => {
  let times = parseInt(nextMove.slice(1));
  move[nextMove[0]](current, times);
});

console.log(
  `Part Two: ${Math.abs(current.coord.x) + Math.abs(current.coord.y)}`
);
