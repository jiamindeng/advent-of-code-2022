const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8").split(/\r\n/);

let steps = input.map((step) => {
  step = step.split(" ");
  step[1] = parseInt(step[1]);
  return { move: step[0], distance: step[1] };
});

let current = { accumulator: 0, location: 0 };

const move = {
  jmp: (current, distance) => {
    current.location += distance;
  },
  acc: (current, distance) => {
    current.accumulator += distance;
    current.location++;
  },
  nop: (current, distance) => {
    current.location++;
  },
};

// ...Brute force
const traverseUntilRetread = (current, steps) => {
  let traversed = Array.from({ length: steps.length }, () => {
    false;
  });
  while (true) {
    if (traversed[current.location]) {
      return current;
    }
    let currentStep = steps[current.location];
    traversed[current.location] = true;
    move[currentStep.move](current, currentStep.distance);
  }
};

console.log(`Part One: ${traverseUntilRetread(current, steps).accumulator}.`);

const findWrongOp = (steps) => {
  let found = false;
  for (let [index, step] of steps.entries()) {
    switch (step.move) {
      case "jmp":
        found = swap(index, "nop", steps);
        break;
      case "nop":
        found = swap(index, "jmp", steps);
        break;
    }
    if (found) {
      break;
    }
  }
  return found;
};

const swap = (index, modMove, steps) => {
  let modSteps = JSON.parse(JSON.stringify(steps));
  modSteps[index].move = modMove;
  return traverseUntilEnd(modSteps);
};

const traverseUntilEnd = (steps) => {
  let current = { accumulator: 0, location: 0 };
  let traversed = Array.from({ length: steps.length }, () => {
    false;
  });

  while (true) {
    if (current.location === steps.length) return current.accumulator;
    if (current.location > steps.length) return false;
    if (traversed[current.location]) return false;
    traversed[current.location] = true;
    let currentStep = steps[current.location];
    move[currentStep.move](current, currentStep.distance);
  }
};

console.log(`Part Two: ${findWrongOp(steps)}`);
