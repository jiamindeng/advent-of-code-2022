const fs = require("fs");

let input = fs
  .readFileSync("./input.txt", "utf-8")
  .split(/\n/)
  .map((line) => line.split(""));

const initArray = (arr) => {
  return [[arr]];
};

const padArray = (fourDArr) => {
  fourDArr.forEach((cube) => {
    cube.forEach((slice) => {
      slice.forEach((line) => {
        line.unshift(".");
        line.push(".");
      });
      slice.unshift(Array.from({ length: slice[0].length }, () => "."));
      slice.push(Array.from({ length: slice[0].length }, () => "."));
    });
    let newSlice = Array.from({ length: cube[0].length }, () =>
      Array.from({ length: cube[0][0].length }, () => ".")
    );
    cube.unshift(newSlice);
    cube.push(newSlice);
  });

  let newCube = Array.from({ length: fourDArr[0].length }, () =>
    Array.from({ length: fourDArr[0][0].length }, () =>
      Array.from({ length: fourDArr[0][0][0].length }, () => ".")
    )
  );

  fourDArr.unshift(newCube);
  fourDArr.push(newCube);
  return fourDArr;
};

const countNeighbors = (i, j, k, l, arr) => {
  let count = 0;

  const coord = [-1, 0, 1];
  const delta = [];

  for (let x = 0; x < coord.length; x++) {
    for (let y = 0; y < coord.length; y++) {
      for (let z = 0; z < coord.length; z++) {
        for (let w = 0; w < coord.length; w++) {
          if (
            !(
              coord[x] === 0 &&
              coord[y] === 0 &&
              coord[z] === 0 &&
              coord[w] === 0
            )
          ) {
            delta.push([coord[x], coord[y], coord[z], coord[w]]);
          }
        }
      }
    }
  }

  delta.forEach((move) => {
    let x = i + move[0];
    let y = j + move[1];
    let z = k + move[2];
    let w = l + move[3];

    if (
      x < arr.length &&
      y < arr[0].length &&
      z < arr[0][0].length &&
      w < arr[0][0][0].length &&
      x >= 0 &&
      y >= 0 &&
      z >= 0 &&
      w >= 0
    ) {
      if (arr[x][y][z][w] === "#") {
        count++;
      }
    }
  });

  return count;
};

const cycleArray = (input) => {
  const arr = padArray(input);
  const newArr = JSON.parse(JSON.stringify(arr));

  for (let x = 0; x < arr.length; x++) {
    for (let y = 0; y < arr[0].length; y++) {
      for (let z = 0; z < arr[0][0].length; z++) {
        for (let w = 0; w < arr[0][0][0].length; w++) {
          let neighbors = countNeighbors(x, y, z, w, arr);
          if (arr[x][y][z][w] === "#") {
            if (!(neighbors === 2 || neighbors === 3)) {
              newArr[x][y][z][w] = ".";
            }
          } else {
            if (neighbors === 3) {
              newArr[x][y][z][w] = "#";
            }
          }
        }
      }
    }
  }
  return newArr;
};

let cycles = 6;
let currentArr = initArray(input);

while (cycles > 0) {
  currentArr = cycleArray(currentArr);
  cycles--;
}

console.log(
  `Part Two: ${
    currentArr
      .flat()
      .flat()
      .flat()
      .filter((char) => char === "#").length
  }`
);
