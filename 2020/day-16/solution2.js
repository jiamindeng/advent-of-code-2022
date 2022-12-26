const fs = require('fs');

let input = fs.readFileSync('./input.txt', 'utf-8');

const YOUR_TICKET = 'your ticket:';
const NEARBY_TICKETS = 'nearby tickets:';

// Probably my worst one to date
let requirements = input
  .substring(0, input.indexOf(YOUR_TICKET))
  .trim()
  .replace(/[a-zA-Z:]/g, '')
  .split(/[ \n]/g)
  .filter((item) => item.length > 0)
  .map((interval) => {
    let range = interval.split(/-/);
    return { min: parseInt(range[0]), max: parseInt(range[1]) };
  });

let requirementKeys = input
  .substring(0, input.indexOf(YOUR_TICKET))
  .trim()
  .replace(/(?<=:).*/g, '')
  .split(/[\r\n:]/)
  .filter((item) => item.length > 0);

let reqs = [];
let requirementObj = {};

for (let i = 0; i < requirements.length; i += 2) {
  reqs.push([requirements[i], requirements[i + 1]]);
  requirementObj[requirementKeys[i / 2]] = {
    ranges: [requirements[i], requirements[i + 1]],
    order: new Set([...Array(requirementKeys.length).keys()]),
  };
}

let yourTicket = input
  .substring(
    input.indexOf(YOUR_TICKET) + YOUR_TICKET.length,
    input.indexOf(NEARBY_TICKETS)
  )
  .trim()
  .split(/,/)
  .map((num) => parseInt(num));

let nearbyTickets = input
  .substring(input.indexOf(NEARBY_TICKETS) + NEARBY_TICKETS.length)
  .trim()
  .split(/\n/)
  .map((ticket) => ticket.split(/,/))
  .map((ticket) => ticket.map((num) => parseInt(num)));

const isValid = (ticket, reqs) => {
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
    if (!validNum) return false;
  }
  return true;
};

let validTickets = nearbyTickets.filter((ticket) => isValid(ticket, reqs));

const determineOrder = (validTicket, requirementObj) => {
  let keys = Object.keys(requirementObj);
  for (let i = 0; i < keys.length; i++) {
    for (let j = 0; j < validTicket.length; j++) {
      let num = validTicket[j];
      if (
        !(
          (num >= requirementObj[keys[i]].ranges[0].min &&
            num <= requirementObj[keys[i]].ranges[0].max) ||
          (num >= requirementObj[keys[i]].ranges[1].min &&
            num <= requirementObj[keys[i]].ranges[1].max)
        )
      ) {
        requirementObj[keys[i]].order.delete(j);
      }
    }
  }
};

validTickets.forEach((ticket) => {
  determineOrder(ticket, requirementObj);
});

let orderedReq = {};
let keys = Object.keys(requirementObj);
keys.forEach((key) => {
  Object.assign(orderedReq, { [key]: requirementObj[key].order });
});

orderedReq = Object.entries(requirementObj).sort((a, b) => {
  return a[1].order.size - b[1].order.size;
});

for (let i = orderedReq.length - 1; i > 0; i--) {
  orderedReq[i][1].order = new Set(
    [...orderedReq[i][1].order].filter(
      (x) => !orderedReq[i - 1][1].order.has(x)
    )
  );
}

let departures = orderedReq.filter((req) => !req[0].indexOf('departure'));

let product = 1;
departures.forEach((property) => {
  let index = property[1].order.values().next().value;
  product *= yourTicket[index];
});

console.log(`Part Two: ${product}`);
