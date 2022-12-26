const fs = require('fs');
const crt = require('./nodejs-chinese-remainder');

let input = fs.readFileSync('./input.txt', 'utf-8').split(/\r\n/);

[currentTime, buses] = input;

let busIds = buses
  .replace(/x/g, '0')
  .split(/,/)
  .map((busId) => parseInt(busId));

let i = 0;
let timeIncrement = Array.from({ length: busIds.length }, () => i++).filter(
  (time, index) => busIds[index] !== 0
);

moduli = busIds.filter((busId) => busId !== 0);

let residues = moduli.map((modulus, i) => modulus - timeIncrement[i]);

// ): This only works for the examples. Numbers are too big for JS, I used Python for my actual input.

console.log(`Part Two: ${crt(moduli, residues)}`);
