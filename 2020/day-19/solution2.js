const fs = require('fs');

let input = fs.readFileSync('./input.txt', 'utf-8').split(/\r\n/);

let ruleset = new Map();

input.splice(0, input.indexOf('')).forEach((rule) => {
  ruleset.set(rule.split(/:/)[0], rule.split(/:/)[1].replace(/"/g, '').trim());
});

// Adapted heavily from https://github.com/tpatel/advent-of-code-2020/blob/main/day19.js

// Rule 0 always calls these two
// Rule 42 will always be executed more times than Rule 31
ruleset.set('8', '42 | 42 8');
ruleset.set('11', '42 31 | 42 11 31');

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

const rule31 = new RegExp(getRegex(ruleset.get('31')), 'g');
const rule42 = new RegExp(getRegex(ruleset.get('42')), 'g');

// Captures text matched in <groupName>
const rule = new RegExp(
  `^(?<group42>(${getRegex(ruleset.get('42'))})+)(?<group31>(${getRegex(
    ruleset.get('31')
  )})+)$`
);

let count = 0;

passwords.forEach((password) => {
  const matches = rule.exec(password);
  if (matches) {
    let numMatchesRule31 = matches.groups.group31.match(rule31).length;
    let numMatchesRule42 = matches.groups.group42.match(rule42).length;
    count = numMatchesRule31 < numMatchesRule42 ? ++count : count;
  }
});

console.log(`Part Two: ${count}`);
