const fs = require("fs");

let input = fs.readFileSync("./input.txt", "utf-8").split(/\r\n/);

// I hope you like spaghetti.
const getProperties = (fragment) => {
  return fragment.split(" ").map((property) => {
    let propertyPair = property.split(":");
    return { [`${propertyPair[0]}`]: propertyPair[1] };
  });
};

const hasRequiredProperties = (passport) => {
  const propertiesToTest = ["iyr", "eyr", "byr", "hcl", "pid", "hgt", "ecl"];
  return propertiesToTest.every((x) => x in passport);
};

// ABANDON HOPE
const hasRequiredPropertiesAgain = (passport) => {
  const propertiesToTest = ["iyr", "eyr", "byr", "hcl", "pid", "hgt", "ecl"];
  if (!propertiesToTest.every((x) => x in passport)) {
    return false;
  }

  let validByr =
    parseInt(passport.byr) >= 1920 && parseInt(passport.byr) <= 2002;

  let validIyr =
    parseInt(passport.iyr) >= 2010 && parseInt(passport.iyr) <= 2020;

  let validEyr =
    parseInt(passport.eyr) >= 2020 && parseInt(passport.eyr) <= 2030;

  let validHgt;
  let units = passport.hgt.slice(-2);
  let height = parseInt(passport.hgt.replace(/\D/g, ""));
  if (units === "cm" && height >= 150 && height <= 193) {
    validHgt = true;
  } else if (units === "in" && height >= 59 && height <= 76) {
    validHgt = true;
  } else {
    validHgt = false;
  }

  let validHcl = false;
  if (
    passport.hcl[0] === "#" &&
    passport.hcl.slice(1).length === 6 &&
    passport.hcl.slice(1).replace(/[0-9a-f]/g, "").length === 0
  ) {
    validHcl = true;
  }

  let eyeColors = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"];

  let validEcl = eyeColors.includes(passport.ecl) ? true : false;

  let validPid = passport.pid.length === 9;
  try {
    parseInt(passport.pid);
  } catch (e) {
    validPid = false;
  }

  return (
    validByr &&
    validIyr &&
    validEyr &&
    validHgt &&
    validHcl &&
    validEcl &&
    validPid
  );
};

let passports = [];
let passport = {};

input.forEach((fragment) => {
  if (fragment) {
    getProperties(fragment).forEach((property) => {
      Object.assign(passport, property);
    });
  } else {
    passports.push(passport);
    passport = {};
  }
});

let numValidPassports = passports
  .map((passport) => hasRequiredProperties(passport))
  .filter((status) => status === true).length;

let numValidPassportsAgain = passports
  .map((passport) => hasRequiredPropertiesAgain(passport))
  .filter((status) => status === true).length;

console.log(`Part One: ${numValidPassports}`);
console.log(`Part Two: ${numValidPassportsAgain}`);
