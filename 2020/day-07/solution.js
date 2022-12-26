const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8").split(/\n/);

// Part One
let bags = {};

const lineToBag = (line) => {
  let innerBagTypes = line
    .replace(/.*?bags/, "") // Remove the top level bag
    .split(",")
    .map((x) =>
      x
        .replace(/.*\d /, "") // Remove everything before and including the number of bags
        .replace(/[^a-zA-Z ]/g, "") // Remove punctuation
        .replace(/(bags|bag)/, "") // Remove bags/bag
        .trim()
    );
  return { [line.replace(/bags.*/, "").trim()]: innerBagTypes };
};

const containsGold = (bagName, bags) => {
  if (!bags[bagName]) {
    return false;
  }
  if (bags[bagName].includes("shiny gold")) {
    return true;
  }
  let innerContainsGold = false;
  for (const innerBagName of bags[bagName]) {
    if (containsGold(innerBagName, bags)) {
      innerContainsGold = true;
    }
  }
  return innerContainsGold;
};

input.forEach((line) => {
  newBag = lineToBag(line);
  Object.assign(bags, newBag);
});

let count = 0;
for (bag in bags) {
  if (containsGold(bag, bags)) {
    count++;
  }
}

console.log(`Part One: ${count}.`);

// Part Two
count = 0;
bags = {};

const lineToBagAgain = (line) => {
  bag = {};
  let innerBagTypes = line
    .replace(/.*?bags/, "")
    .split(",")
    .map((x) => [
      Number(x.replace(/[^\d]+/g, "")),
      x
        .replace(/.*\d /, "")
        .replace(/[^a-zA-Z ]/g, "")
        .replace(/(bags|bag)/, "")
        .trim(),
    ]);
  return { [line.replace(/bags.*/, "").trim()]: innerBagTypes };
};

input.forEach((line) => {
  Object.assign(bags, lineToBagAgain(line));
});

const countBags = (bagName, bags) => {
  if (!bags[bagName]) {
    return 0;
  }
  let count = 0;
  bags[bagName].forEach((innerBag) => {
    count += innerBag[0] + innerBag[0] * countBags(innerBag[1], bags);
  });
  return count;
};

console.log(`Part Two: ${countBags("shiny gold", bags)}.`);
