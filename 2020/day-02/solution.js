const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8").split(/[- :\n]+/);

let scenarios = [];
for (let i = 0; i < input.length; i = i + 4) {
  scenarios.push({
    min: input[i],
    max: input[i + 1],
    char: input[i + 2],
    str: input[i + 3],
  });
}

const letterCount = (scenario) => {
  let matches = scenario.str.match(new RegExp(scenario.char, "g"));
  return matches ? matches.length : 0;
};

const solution1 = (scenarios) =>
  scenarios
    .map((scenario) => {
      let occurrences = letterCount(scenario);
      return occurrences >= scenario.min && occurrences <= scenario.max;
    })
    .filter((match) => match === true).length;

console.log(`Part One: ${solution1(scenarios)}`);

const solution2 = (scenarios) =>
  scenarios
    .map((scenario) => {
      let matchesCharOne = scenario.str[scenario.min - 1] === scenario.char;
      let matchesCharTwo = scenario.str[scenario.max - 1] === scenario.char;
      if (matchesCharOne === matchesCharTwo) {
        return false;
      } else {
        return true;
      }
    })
    .filter((match) => match === true).length;

console.log(`Part Two: ${solution2(scenarios)}`);
