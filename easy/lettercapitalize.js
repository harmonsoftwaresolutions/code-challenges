const assert = require('assert');

// Have the function LetterCapitalize(str) take the str parameter being passed
// and capitalize the first letter of each word. Words will be separated by
// only one space. 

// Input:"hello world"
// Output:"Hello World"
// 
// Input:"i ran there"
// Output:"I Ran There"

function letterCapitalize(str) {
  // split and set flag for spaces
  const arr = str.split("");
  let fl = true;
  for (let i = 0; i < arr.length; i++) {
    if (fl) {
      const item = arr[i].toString();
      arr[i] = item.toUpperCase();
      fl = false;
    };

    if (arr[i] === " ") {
      fl = true;
    }
  }
  str = arr.join("");
  return str;
}

const in1 = "hello world";
const expect1 = "Hello World";
const test1 = letterCapitalize(in1);
assert.strictEqual(test1, expect1, `should be ${expect1}`);


const in2 = "i ran there";
const expect2 = "I Ran There";
const test2 = letterCapitalize(in2);
assert.strictEqual(test2, expect2, `should be ${expect2}`);
