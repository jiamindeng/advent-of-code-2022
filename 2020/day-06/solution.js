const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8").split(/\r\n\r\n/);

let questionCounts = input.map((group) => {
  group = group.replace(/\r\n/g, "");
  question_set = new Set(group.split(""));
  return question_set.size;
});

console.log(
  `Part One: ${questionCounts.reduce((acc, current) => acc + current)}`
);

let questionCountsAgain = input.map((group) => {
  group = group.split(/\n/).map((line) => line.split(""));
  intersection = group.reduce((a, b) =>
    b.filter(Set.prototype.has, new Set(a))
  );
  return intersection.length;
});

console.log(
  `Part Two: ${questionCountsAgain.reduce((acc, current) => acc + current)}`
);
