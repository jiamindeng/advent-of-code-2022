const fs = require('fs');

let input = fs.readFileSync('./input.txt', 'utf-8').split(/\r\n\r\n/);

let tiles = input.map((string) => {
  let lines = string.split(/\r\n/);
  let tile = {
    id: lines[0].replace(/Tile /, '').replace(/:/, ''),
    sides: {
      top: lines[1],
      bottom: lines[lines.length - 1],
      left: lines
        .slice(1)
        .map((line) => line[0])
        .join(''),
      right: lines
        .slice(1)
        .map((line) => line[line.length - 1])
        .join(''),
    },
    body: lines.slice(1),
    neighborTiles: new Set(),
    sharedSides: new Set(),
  };
  return tile;
});

const labelAdjacentTiles = (currentTile) => {
  for (tile of tiles) {
    if (tile.id !== currentTile.id) {
      for (currentProp in currentTile.sides) {
        for (prop in tile.sides) {
          if (
            prop !== 'neighbors' &&
            currentProp !== 'neighbors' &&
            currentTile.sides[currentProp] === tile.sides[prop]
          ) {
            currentTile.neighborTiles.add(tile.id);
            tile.neighborTiles.add(currentTile.id);
            currentTile.sharedSides.add(currentProp);
            tile.sharedSides.add(prop);
          } else if (
            prop !== 'neighbors' &&
            currentProp !== 'neighbors' &&
            !Array.isArray(tile.sides[prop]) &&
            currentTile.sides[currentProp] ===
              tile.sides[prop].split('').reverse().join('')
          ) {
            currentTile.neighborTiles.add(tile.id);
            tile.neighborTiles.add(currentTile.id);
            currentTile.sharedSides.add(currentProp);
            tile.sharedSides.add(prop);
          }
        }
      }
    }
  }
};

tiles.forEach((tile) => labelAdjacentTiles(tile));

let corners = tiles
  .filter((tile) => tile.neighborTiles.size === 2)
  .map((tile) => parseInt(tile.id))
  .reduce((acc, current) => acc * current);

console.log(`Part One: ${corners}`);
