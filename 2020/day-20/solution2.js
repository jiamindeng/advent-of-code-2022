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

let puzzleSideLength = Math.sqrt(tiles.length);

const flipTile = (tile) => {
  tile.body = tile.body.reverse();
  tile.sides.top = tile.body[0];
  tile.sides.bottom = tile.body[tile.body.length - 1];
  tile.sides.left = tile.body.map((line) => line[0]).join('');
  tile.sides.right = tile.body.map((line) => line[line.length - 1]).join('');
};

const rotateTile = (tile) => {
  let bodyArray = tile.body.map((line) => line.split(''));
  bodyArray = bodyArray[0].map((_, colIndex) =>
    bodyArray.map((row) => row[colIndex])
  );

  tile.body = bodyArray.reverse().map((line) => line.join(''));
  tile.sides.top = tile.body[0];
  tile.sides.bottom = tile.body[tile.body.length - 1];
  tile.sides.left = tile.body.map((line) => line[0]).join('');
  tile.sides.right = tile.body.map((line) => line[line.length - 1]).join('');
};

const isCorrectPair = (currentProp, prop) => {
  return (
    (currentProp === 'top' && prop === 'bottom') ||
    (currentProp === 'bottom' && prop === 'top') ||
    (currentProp === 'left' && prop === 'right') ||
    (currentProp === 'right' && prop === 'left')
  );
};

const getConfigurations = (tile) => {
  const configurations = [];
  let notFlipped = false;
  let tileCopy = JSON.parse(JSON.stringify(tile));
  for (let i = 0; i < 2; i++) {
    for (let j = 0; j < 4; j++) {
      if (i && !notFlipped) {
        flipTile(tileCopy);
        notFlipped = true;
      }
      rotateTile(tileCopy);
      configurations.push(JSON.parse(JSON.stringify(tileCopy)));
    }
  }
  return configurations;
};

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

let puzzle = Array.from({ length: puzzleSideLength * 2 }, () =>
  Array.from({ length: puzzleSideLength * 2 }, () => 0)
);

tiles.forEach((tile) => labelAdjacentTiles(tile));

let corners = tiles.filter((tile) => tile.neighborTiles.size === 2);

let configurations = {};
tiles.forEach((tile) =>
  Object.assign(configurations, { [tile.id]: getConfigurations(tile) })
);

let start = configurations[corners[0].id][0];

puzzle[puzzleSideLength][puzzleSideLength] = start;

const putSides = () => {
  puzzle.forEach((row, rowIndex) =>
    row.forEach((tile, tileIndex) => {
      if (tile !== 0) {
        Object.entries(tile.sides).forEach(([tileProp, side]) => {
          for (id in configurations) {
            if (id !== tile.id) {
              configurations[id].forEach((configuration) => {
                for (prop in configuration.sides) {
                  if (
                    isCorrectPair(tileProp, prop) &&
                    side === configuration.sides[prop]
                  ) {
                    switch (tileProp) {
                      case 'top':
                        puzzle[rowIndex - 1][tileIndex] = configuration;
                        break;
                      case 'bottom':
                        puzzle[rowIndex + 1][tileIndex] = configuration;
                        break;
                      case 'left':
                        puzzle[rowIndex][tileIndex - 1] = configuration;
                        break;
                      case 'right':
                        puzzle[rowIndex][tileIndex + 1] = configuration;
                        break;
                    }
                  }
                }
              });
            }
          }
        });
      }
    })
  );
};

let count = tiles.length;

while (count > 0) {
  putSides();
  count--;
}

const constrain = (puzzle) => {
  return puzzle
    .map((row) => row.filter((tile) => tile))
    .filter((row) => row.length);
};

let finishedPuzzle = constrain(puzzle).map((row) =>
  row.map((tile) => tile.body)
);

let mergedPuzzle = [];

finishedPuzzle.forEach((horizontalChunk) => {
  for (let i = 1; i < horizontalChunk[0].length - 1; i++) {
    let mergedPuzzleRow = '';
    for (let j = 0; j < puzzleSideLength; j++) {
      mergedPuzzleRow += horizontalChunk[j][i].slice(
        1,
        horizontalChunk[j][i].length - 1
      );
    }
    mergedPuzzle.push(mergedPuzzleRow);
  }
});

const getPuzzleConfigurations = (puzzle) => {
  const flip = (puzzle) => {
    return puzzle.reverse();
  };

  const rotate = (puzzle, n) => {
    let puzzleCopy = puzzle.map((line) => line.split(''));
    puzzleCopy = puzzleCopy[0].map((_, colIndex) =>
      puzzleCopy.map((row) => row[colIndex])
    );
    return puzzleCopy.reverse().map((line) => line.join(''));
  };

  const configurations = [];
  for (let i = 0; i < 2; i++) {
    for (let j = 0; j < 4; j++) {
      let puzzleCopy = JSON.parse(JSON.stringify(puzzle));
      if (i) {
        puzzleCopy = flip(puzzleCopy);
      }
      for (let k = 0; k < j; k++) {
        puzzleCopy = rotate(puzzleCopy);
      }

      configurations.push(puzzleCopy);
    }
  }
  return configurations;
};

let puzzleConfigurations = getPuzzleConfigurations(mergedPuzzle);

const findNessie = (puzzle) => {
  let nessieInput = fs.readFileSync('./nessie.txt', 'utf-8');
  let nessie = nessieInput.split(/\r\n/).map((line) => line.split(''));

  let nessieCoordinates = nessie
    .map((line, lineIndex) =>
      line
        .map((char, charIndex) => (char === '#' ? [lineIndex, charIndex] : 0))
        .filter((char) => char)
    )
    .flat();

  let puzzleArray = puzzle.map((line) => line.split(''));

  for (let i = 0; i < puzzleArray.length; i++) {
    for (let j = 0; j < puzzleArray[0].length; j++) {
      let isNessie = true;
      let nessieFound = [];
      outer: for ([dy, dx] of nessieCoordinates) {
        let y = i + dy;
        let x = j + dx;
        if (y < puzzleArray.length && x < puzzleArray.length) {
          if (puzzleArray[y][x] !== '#') {
            isNessie = false;
            nessieFound = [];
            break outer;
          } else {
            nessieFound.push([y, x]);
          }
        } else {
          nessieFound = [];
          break outer;
        }
      }
      if (isNessie) {
        nessieFound.forEach(([y, x]) => (puzzleArray[y][x] = 0));
      }
    }
  }
  return puzzleArray.map((line) => line.join('')).join('\n');
};

let search = puzzleConfigurations
  .map((puzzleConfiguration) => findNessie(puzzleConfiguration))
  .filter((puzzleConfiguration) => puzzleConfiguration.includes(0));

[correctPuzzle] = search;

console.log(correctPuzzle);

console.log(`Part Two: ${correctPuzzle.match(/#/g).length}`);
