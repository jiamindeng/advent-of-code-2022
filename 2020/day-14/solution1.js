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

// (ಥ﹏ಥ)
// TODO: Redo later
const getValue = (mask, update) => {
  let value = Array.from({ length: mask.length }, () => "0");
  let binaryValue = Array.from(decimalToBinary(update[1]));
  let maskIndexOnes = indexOfAll(Array.from(mask), "1");
  let maskIndexZeros = indexOfAll(Array.from(mask), "0");
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

  maskIndexZeros.forEach((index) => {
    value[index] = "0";
  });
  return binaryToDecimal(value.join(""));
};

const updateValue = (mask, update) => {
  return { [update[0]]: getValue(mask, update) };
};

values = {};

groups.forEach((group) => {
  let mask = group[0];
  let updates = group.slice(1);
  updates.forEach((update) => {
    Object.assign(values, updateValue(mask, update));
  });
});

let onlyValues = Object.entries(values).map((value) => value[1]);

console.log(`Part One: ${onlyValues.reduce((acc, current) => acc + current)}`);
