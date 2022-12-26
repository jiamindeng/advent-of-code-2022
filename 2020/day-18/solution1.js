const fs = require('fs');

let input = fs
  .readFileSync('./input.txt', 'utf-8')
  .split(/\r\n/)
  .map((line) =>
    line
      .split('')
      .filter((char) => char !== ' ')
      .map((char) => {
        return Number.isNaN(parseInt(char)) ? char : parseInt(char);
      })
  );

// Anything to avoid using a binary expression tree
const evaluate = (subExpression) => {
  let subExp = subExpression.filter((char) => char !== '(' && char !== ')');
  while (subExp.length > 1) {
    let operation = subExp[1];
    let result;
    if (operation === '+') {
      result = subExp[0] + subExp[2];
    } else if (operation === '*') {
      result = subExp[0] * subExp[2];
    }
    subExp.splice(0, 3);
    subExp.unshift(result);
  }
  return subExp[0];
};

const runExpression = (line) => {
  while (line.indexOf('(') !== -1) {
    let parentheses = [];
    for (let i = 0; i < line.length; i++) {
      if (line[i] === '(') {
        parentheses[0] = i;
      } else if (line[i] === ')') {
        parentheses.push(i);
        break;
      }
    }
    let subExp = line.slice(parentheses[0] + 1, parentheses[1]);
    line.splice(
      parentheses[0],
      parentheses[1] - parentheses[0] + 1,
      evaluate(subExp)
    );
  }
  return evaluate(line);
};

let sum = input
  .map((line) => runExpression(line))
  .reduce((acc, current) => acc + current);

console.log(`Part One: ${sum}`);
