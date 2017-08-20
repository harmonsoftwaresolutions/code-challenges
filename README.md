# CoderByte / Code Challenges
No code exists in the Master branch. Check out the language specific branches
for examples.

# Easy Challenges

## First Factorial
Have the function FirstFactorial(num) take the num parameter being passed
and return the factorial of it (e.g. if num = 4, return (4 * 3 * 2 * 1)).
For the test cases, the range will be between 1 and 18 and the input will
always be an integer.

### Test Cases:
8 40320

### Reverse String
Have the function FirstReverse(str) take the str parameter being passed and
return the string in reversed order. For example: if the input string is
"Hello World and Coders" then your program should return the string
"sredoC dna dlroW olleH".

### Test Cases:
Hello World and Coders
sredoC dna dlroW olleH

### Longest Word
Have the function LongestWord(sen) take the sen parameter being passed and
return the largest word in the string. If there are two or more words that
are the same length, return the first word from the string with that length.
Ignore punctuation and assume sen will not be empty.

### Test Cases:
a beautiful sentence^&!
letter after letter
hello world
123456789 98765432

## Capitalize Letter
Have the function LongestWord(sen) take the sen parameter being passed and
return the largest word in the string. If there are two or more words that
are the same length, return the first word from the string with that length.
Ignore punctuation and assume sen will not be empty.

### Test Cases:
this is all lower case letters

## Letter Changes
Have the function LetterChanges(str) take the str parameter being passed
and modify it using the following algorithm. Replace every letter in the
string with the letter following it in the alphabet
(ie. c becomes d, z becomes a). Then capitalize every vowel in this new
string (a, e, i, o, u) and finally return this modified string.

### Test Cases:
Input:"hello*3"
Output:"Ifmmp*3"

Input:"fun times!"
Output:"gvO Ujnft!"

## Simple Symbols
Have the function SimpleSymbols(str) take the str parameter being passed and
determine if it is an acceptable sequence by either returning the string
true or false. The str parameter will be composed of + and = symbols with
several letters between them (ie. ++d+===+c++==a) and for the string to be
true each letter must be surrounded by a + symbol. So the string to the left
would be false. The string will not be empty and will have at least one
letter.

### Test Cases:
Input:"+d+=3=+s+"
Output:"true"

Input:"f++d+"
Output:"false"

## Simple Adding
Have the function SimpleAdding(num) add up all the numbers from 1 to num.
For example: if the input is 4 then your program should return 10 because
1 + 2 + 3 + 4 = 10. For the test cases, the parameter num will be any number
from 1 to 1000.

### Test Cases:
Input:12
Output:78

Input:140
Output:9870

## Check Nums
have the function CheckNums(num1,num2) take both parameters being passed and return the string true if num2 is greater than num1, otherwise return the string false. If the parameter values are equal to each other then return the string -1. 

### Test Cases:
Input:3 & num2 = 122
Output:"true"

Input:67 & num2 = 67
Output:"-1"

## Time Convert
Have the function TimeConvert(num) take the num parameter being passed and
return the number of hours and minutes the parameter converts to
(ie. if num = 63 then the output should be 1:3). Separate the number of
hours and minutes with a colon. 

### Test Cases:
Input:126
Output:"2:6"

Input:45
Output:"0:45"

## Vowel Count
Have the function VowelCount(str) take the str string parameter being
passed and return the number of vowels the string contains
(ie. "All cows eat grass and moo" would return 8). Do not count y as a
vowel for this challenge. 

### Test Cases:
Input:"hello"
Output:2

Input:"coderbyte"
Output:3

## Word Count
Have the function WordCount(str) take the str string parameter being passed
and return the number of words the string contains
(e.g. "Never eat shredded wheat or cake" would return 6). Words will be
separated by single spaces. 

### Test Cases:
Input:"Hello World"
Output:2

Input:"one 22 three"
Output:3

## Alphabet Soup
Have the function AlphabetSoup(str) take the str string parameter being
passed and return the string with the letters in alphabetical order
(ie. hello becomes ehllo). Assume numbers and punctuation symbols will not
be included in the string.

### Test Cases:
Input:"coderbyte"
Output:"bcdeeorty"

Input:"hooplah"
Output:"ahhloop"

## Swap Case
Have the function SwapCase(str) take the str parameter and swap the case
of each character. For example: if str is "Hello World" the output should
be hELLO wORLD. Let numbers and symbols stay the way they are.

### Test Cases:
Input:"Hello-LOL"
Output:"hELLO-lol"

Input:"Sup DUDE!!?"
Output:"sUP dude!!?"
