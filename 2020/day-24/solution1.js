const fs = require('fs');

let input = fs.readFileSync('./input.txt', 'utf-8').split(/\r\n/);

const tokenize = (string) =>
  [...string.matchAll(/e|se|sw|w|nw|ne/g)].map((x) => x[0]);

const containsAll = (elements, array) =>
  elements.every((element) => array.includes(element));

const removeOnce = (elements, array) => {
  elements.forEach((element) => array.splice(array.indexOf(element), 1));
};

const removeTriplets = (array) => {
  const triplets = [
    ['ne', 'se', 'w'],
    ['e', 'nw', 'sw'],
  ];

  triplets.forEach((triplet) => {
    while (containsAll(triplet, array)) {
      removeOnce(triplet, array);
    }
  });
};

const removeDoubles = (array) => {
  const doubles = [
    ['e', 'w'],
    ['ne', 'sw'],
    ['nw', 'se'],
  ];

  doubles.forEach((double) => {
    while (containsAll(double, array)) {
      removeOnce(double, array);
    }
  });
};

const replaceDoubles = (array) => {
  let conversions = [
    [['ne', 'w'], 'nw'],
    [['nw', 'e'], 'ne'],
    [['nw', 'sw'], 'w'],
    [['ne', 'se'], 'e'],
    [['se', 'w'], 'sw'],
    [['sw', 'e'], 'se'],
  ];

  conversions.forEach((conversion) => {
    while (containsAll(conversion[0], array)) {
      removeOnce(conversion[0], array);
      array.push(conversion[1]);
    }
  });
};

const removeRedundant = (array) => {
  removeTriplets(array);
  removeDoubles(array);
  replaceDoubles(array);
  array.sort();
};

const lines = input.map((line) => tokenize(line));

lines.forEach((line) => {
  removeRedundant(line);
});

tileMap = {};

lines
  .sort()
  .map((line) => JSON.stringify(line))
  .forEach((line) => {
    if (tileMap[line]) {
      tileMap[line]++;
    } else {
      tileMap[line] = 1;
    }
  });

const finalTiles = Object.values(tileMap);
const numBlack = finalTiles.filter((num) => num % 2 === 1).length;

console.log(`Part One: ${numBlack}`);
