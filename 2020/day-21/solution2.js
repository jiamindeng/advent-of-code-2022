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

allergens = [...allergens.entries()].sort((a, b) => a[1].size - b[1].size);

let processedAllergens = [];

while (allergens[0]) {
  let [allergen, ingredients] = allergens[0];
  for (let i = 1; i < allergens.length; i++) {
    allergens[i][1].delete(ingredients.values().next().value);
  }
  processedAllergens.push(allergens.shift());
  allergens = allergens.sort((a, b) => a[1].size - b[1].size);
}

processedAllergens = processedAllergens.sort((a, b) => {
  if (a[0] < b[0]) return -1;
  if (a[0] > b[0]) return 1;
  return 0;
});

console.log(
  `Part Two: ${processedAllergens
    .map((allergen) => allergen[1].values().next().value)
    .join()}`
);
