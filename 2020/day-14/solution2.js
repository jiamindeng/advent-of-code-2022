const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8").split(/\r\n/);

let groups = [];
let maskGroup = [];

input.forEach((line, index) => {
  if (line.startsWith("mask")) {
    groups.push(maskGroup);
    maskGroup = [line.replace(/mask = /, "")];
  } else {
    maskGroup = [
      ...maskGroup,
      line
        .replace(/mem\[/, "")
        .split(/\] = /)
        .map((item) => parseInt(item)),
    ];
  }
});

groups.push(maskGroup);

const binaryToDecimal = (binary) => {
  return parseInt(binary, 2);
};
const indexOfAll = (arr, val) =>
  arr.reduce((acc, el, i) => (el === val ? [...acc, i] : acc), []);

const decimalToBinary = (decimal) => decimal.toString(2);

const getAddress = (mask, update) => {
  let value = Array.from({ length: mask.length }, () => "0");
  let binaryValue = Array.from(decimalToBinary(update[0]));
  let maskIndexOnes = indexOfAll(Array.from(mask), "1");
  let maskIndexExes = indexOfAll(Array.from(mask), "X");
  let addresses = [];

  for (
    let i = value.length - binaryValue.length, j = 0;
    i < value.length && j < binaryValue.length;
    i++, j++
  ) {
    value[i] = binaryValue[j];
  }

  maskIndexOnes.forEach((index) => {
    value[index] = "1";
  });

  for (let i = 0; i < Math.pow(2, maskIndexExes.length); i++) {
    let bitstring = Array.from(
      decimalToBinary(i).padStart(maskIndexExes.length, 0)
    );
    maskIndexExes.forEach((index, i) => {
      value[index] = bitstring[i];
    });
    addresses.push({ [binaryToDecimal(value.join(""))]: update[1] });
  }
  return addresses;
};

values = {};

groups.forEach((group) => {
  let mask = group[0];
  let updates = group.slice(1);
  updates.forEach((update) => {
    getAddress(mask, update).forEach((operation) => {
      Object.assign(values, operation);
    });
  });
});

let onlyValues = Object.entries(values).map((value) => value[1]);

console.log(`Part Two: ${onlyValues.reduce((acc, current) => acc + current)}`);
