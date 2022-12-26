const fs = require('fs');

let input = fs
  .readFileSync('./input.txt', 'utf-8')
  .split(/\r\n/)
  .map((line) => {
    let split = line.split(/ \(contains /);
    return [split[0].split(/ /), split[1].replace(/\)/, '').split(/, /)];
  });

let ingredients = [];
let allergens = new Map();

input.forEach(([currentIngredients, currentAllergens]) => {
  ingredients.push(...currentIngredients);
  currentAllergens.forEach((allergen) => {
    let currentIngredientSet = new Set(currentIngredients);
    if (allergens.has(allergen)) {
      let globalIngredientSet = allergens.get(allergen);
      let intersection = new Set(
        [...currentIngredientSet].filter((x) => globalIngredientSet.has(x))
      );
      allergens.set(allergen, intersection);
    } else {
      allergens.set(allergen, currentIngredientSet);
    }
  });
});

allergens = new Set(
  Array.from(allergens.values())
    .map((set) => {
      return Array.from(set);
    })
    .flat()
);

let count = ingredients.filter((ingredient) => !allergens.has(ingredient))
  .length;

console.log(`Part One: ${count}`);
