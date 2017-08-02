const assert = require('assert');

// Have the function FirstReverse(str) take the str parameter being passed and return the string in reversed order. For example: if the input string is "Hello World and Coders" then your program should return the string sredoC dna dlroW olleH. 

function FirstReverse(str) { 
  const chars = str.split("");
  // console.log("chars ", chars);
  const arr = [];
  const ln = str.length;
  let idx = ln;
  for(let i = 0; i < ln; i++) {
      // console.log('ln ',ln);
    // console.log('idx char', chars[idx - 1]);
      arr.push(chars[idx - 1]);
      idx = idx - 1;
    // console.log('arr ', arr);
  }
  
  const revStr = arr.join("");
  // console.log('revStr ', revStr)
  return revStr; 
}
   
a1 = "coderbyte";
r1 = "etybredoc";
t1 = FirstReverse(a1);                            
// console.log(t1);
assert.strictEqual(t1, r1, `should be ${r1}`);

a2 = "I Love Code";
r2 = "edoC evoL I";
t2 = FirstReverse(a2);                            
// console.log(t2);
assert.strictEqual(t2, r2, `should be ${r2}`);
