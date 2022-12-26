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

const evaluate = (subExpression) => {
  let subExp = subExpression.filter((char) => char !== '(' && char !== ')');

  let addition = subExp.indexOf('+');
  while (addition !== -1) {
    let sum = subExp[addition - 1] + subExp[addition + 1];
    subExp.splice(addition - 1, 3, sum);
    addition = subExp.indexOf('+');
  }

  let multiplication = subExp.indexOf('*');
  while (multiplication !== -1) {
    let product = subExp[multiplication - 1] * subExp[multiplication + 1];
    subExp.splice(multiplication - 1, 3, product);
    multiplication = subExp.indexOf('*');
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

console.log(`Part Two: ${sum}`);
