const fs = require('fs');

let input = fs
  .readFileSync('./input.txt', 'utf-8')
  .split(/\r\n/)
  .map((int) => parseInt(int));

const [targetCardPublicKey, targetDoorPublicKey] = input;

const subjectNum = 7;

let encryptionKey = 1;
let value = 1;
let loopSize = 0;

while (value !== targetDoorPublicKey) {
  value = (value * subjectNum) % 20201227;
  encryptionKey = (encryptionKey * targetCardPublicKey) % 20201227;
  loopSize++;
}

console.log(`Part One: ${encryptionKey}`);
