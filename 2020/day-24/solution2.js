const fs = require('fs');

let input = fs.readFileSync('./input.txt', 'utf-8').split(/\r\n/);

const tokenize = (string) =>
  [...string.matchAll(/e|se|sw|w|nw|ne/g)].map((x) => x[0]);

// Redid Part One with axial coordinates
// https://www.redblobgames.com/grids/hexagons/#coordinates
const move = {
  nw: { dx: 0, dy: -1 },
  ne: { dx: 1, dy: -1 },
  w: { dx: -1, dy: 0 },
  e: { dx: 1, dy: 0 },
  sw: { dx: -1, dy: 1 },
  se: { dx: 0, dy: 1 },
};

let blackTiles = new Set();

const lines = input.map((line) => tokenize(line));

lines.forEach((line) => {
  let x = 0;
  let y = 0;
  line.forEach((direction) => {
    x += move[direction].dx;
    y += move[direction].dy;
  });

  if (blackTiles.has(`${x},${y}`)) {
    blackTiles.delete(`${x},${y}`);
  } else {
    blackTiles.add(`${x},${y}`);
  }
});

console.log(`Part One: ${blackTiles.size}`);

const getNeighbors = (x, y) => {
  const neighbors = [];
  for (direction in move) {
    neighbors.push({ x: x + move[direction].dx, y: y + move[direction].dy });
  }
  return neighbors;
};

let count = 100;
while (count > 0) {
  const nextBlackTiles = new Set();

  const currentBlackTiles = [...blackTiles.keys()].map((coordinates) =>
    coordinates.split(',').map((char) => parseInt(char))
  );

  currentBlackTiles.forEach((tile) => {
    const [x, y] = tile;
    const cells = getNeighbors(x, y);
    cells.push({ x, y });

    cells.forEach((cell) => {
      const cellId = `${cell.x},${cell.y}`;
      const neighbors = getNeighbors(cell.x, cell.y);
      const totalBlackTiles = neighbors.filter((neighbor) =>
        blackTiles.has(`${neighbor.x},${neighbor.y}`)
      ).length;

      if (blackTiles.has(cellId)) {
        if (totalBlackTiles === 0 || totalBlackTiles > 2) {
          nextBlackTiles.delete(cellId);
        } else {
          nextBlackTiles.add(cellId);
        }
      } else {
        if (totalBlackTiles === 2) {
          nextBlackTiles.add(cellId);
        }
      }
    });
  });

  blackTiles = nextBlackTiles;
  count--;
}

console.log(`Part Two: ${blackTiles.size}`);
