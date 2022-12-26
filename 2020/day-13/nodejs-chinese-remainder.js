/**
 * From https://github.com/pnicorelli/nodejs-chinese-remainder/blob/master/chinese_remainder.js
 * Based on http://rosettacode.org/wiki/Chinese_remainder_theorem (python implementation)
 * solve a system of linear congruences by applying the Chinese Remainder Theorem
 *
 * 	X = a1  (mod n1)
 *  X = a2  (mod n2)
 *
 * This function will be called as:
 *
 * chineseRemainder( [a1, a2], [n1, n2])
 * @return {integer}
 */

function mul_inv(a, b) {
  let b0 = b;
  let x0 = 0;
  let x1 = 1;
  let q, tmp;
  if (b == 1) {
    return 1;
  }
  while (a > 1) {
    q = Math.floor(a / b);
    tmp = a;
    a = b;
    b = tmp % b;
    tmp = x0;
    x0 = x1 - q * x0;
    x1 = tmp;
  }
  if (x1 < 0) {
    x1 = x1 + b0;
  }
  return x1;
}

function chineseRemainder(moduli, residues) {
  let sum = 0;
  let prod = moduli.reduce((acc, current) => acc * current);
  let zip = moduli.map((e, i) => {
    return [e, residues[i]];
  });
  zip.forEach(([modulus, residue]) => {
    let p = Math.floor(prod / modulus);
    sum += residue * mul_inv(p, modulus) * p;
  });
  return sum % prod;
}

module.exports = chineseRemainder;
