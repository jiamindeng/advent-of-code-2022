const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8");

const YOUR_TICKET = "your ticket:";
const NEARBY_TICKETS = "nearby tickets:";

let requirements = input
  .substring(0, input.indexOf(YOUR_TICKET))
  .trim()
  .replace(/[a-zA-Z:]/g, "")
  .split(/[ \n]/g)
  .filter((item) => item.length > 0)
  .map((interval) => {
    let range = interval.split(/-/);
    return { min: parseInt(range[0]), max: parseInt(range[1]) };
  });

let reqs = [];

for (let i = 0; i < requirements.length; i += 2) {
  reqs.push([requirements[i], requirements[i + 1]]);
}

let yourTicket = input
  .substring(
    input.indexOf(YOUR_TICKET) + YOUR_TICKET.length,
    input.indexOf(NEARBY_TICKETS)
  )
  .trim();

let nearbyTickets = input
  .substring(input.indexOf(NEARBY_TICKETS) + NEARBY_TICKETS.length)
  .trim()
  .split(/\n/)
  .map((ticket) => ticket.split(/,/))
  .map((ticket) => ticket.map((num) => parseInt(num)));

const getInvalidNumbers = (ticket, reqs) => {
  let invalidNums = [];
  for (let i = 0; i < ticket.length; i++) {
    let validNum = false;
    for (let j = 0; j < reqs.length; j++) {
      if (
        (ticket[i] >= reqs[j][0].min && ticket[i] <= reqs[j][0].max) ||
        (ticket[i] >= reqs[j][1].min && ticket[i] <= reqs[j][1].max)
      ) {
        validNum = true;
      }
    }
    if (!validNum) invalidNums.push(ticket[i]);
  }
  return invalidNums;
};

let invalidTickets = nearbyTickets.map((ticket) =>
  getInvalidNumbers(ticket, reqs)
);

console.log(
  `Part One: ${invalidTickets.flat().reduce((acc, current) => acc + current)}`
);
