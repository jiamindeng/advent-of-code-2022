const fs = require('fs');

let input = fs.readFileSync('./input.txt', 'utf-8').split(/\r\n/);

let ruleset = new Map();

input.splice(0, input.indexOf('')).forEach((rule) => {
  ruleset.set(rule.split(/:/)[0], rule.split(/:/)[1].replace(/"/g, '').trim());
});

let passwords = input.splice(input.indexOf('') + 1);

const getRegex = (rules) => {
  let regex = '';

  if (rules === 'a' || rules === 'b') {
    regex = rules;
  } else if (!rules.includes('|')) {
    regex = rules
      .split(' ')
      .map((ruleId) => getRegex(ruleset.get(ruleId)))
      .join('');
  } else {
    let options = rules.split(' | ');
    regex = `(${getRegex(options[0])}|${getRegex(options[1])})`;
  }
  return regex;
};

const regex = new RegExp(`^${getRegex(ruleset.get('0'))}$`);

const matches = passwords
  .map((password) => regex.test(password))
  .filter((match) => match).length;

console.log(`Part One: ${matches}`);
