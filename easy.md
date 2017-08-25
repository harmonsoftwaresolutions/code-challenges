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

## AB Check
Have the function ABCheck(str) take the str parameter being passed and
return the string true if the characters a and b are separated by exactly
3 places anywhere in the string at least once (ie. "lane borrowed" would
result in true because there is exactly three characters between a and b).
Otherwise return the string false. 

### Test Cases:
Input:"after badly"
Output:"false"

Input:"Laura sobs"
Output:"true"

## Ex Oh
Have the function ExOh(str) take the str parameter being passed and
return the string true if there is an equal number of x's and o's,
otherwise return the string false. Only these two letters will be entered
in the string, no punctuation or numbers. For example: if str is
"xooxxxxooxo" then the output should return false because there are 6 x's
and 5 o's. 

### Test Cases:
Input:"xooxxo"
Output:"true"

Input:"x"
Output:"false"

## Palindrome
Have the function Palindrome(str) take the str parameter being passed
and return the string true if the parameter is a palindrome, (the string
is the same forward as it is backward) otherwise return the string false.
For example: "racecar" is also "racecar" backwards. Punctuation and numbers
will not be part of the string. 

### Test Cases:
Input:"never odd or even"
Output:"true"

Input:"eye"
Output:"true"

## Arith Geo
Have the function ArithGeo(arr) take the array of numbers stored in arr
and return the string "Arithmetic" if the sequence follows an arithmetic
pattern or return "Geometric" if it follows a geometric pattern. If the
sequence doesn't follow either pattern return -1. An arithmetic sequence
is one where the difference between each of the numbers is consistent,
where as in a geometric sequence, each term after the first is multiplied
by some constant or common ratio. Arithmetic example: [2, 4, 6, 8] and
Geometric example: [2, 6, 18, 54]. Negative numbers may be entered as
parameters, 0 will not be entered, and no array will contain all the same
elements. 

### Test Cases:
Input:5,10,15
Output:"Arithmetic"

Input:2,4,16,24
Output:-1

## Array Addition 1
Have the function ArrayAdditionI(arr) take the array of numbers stored in
arr and return the string true if any combination of numbers in the array
can be added up to equal the largest number in the array, otherwise return
the string false. For example: if arr contains [4, 6, 23, 10, 1, 3] the
output should return true because 4 + 6 + 10 + 3 = 23. The array will not
be empty, will not contain all the same elements, and may contain negative
numbers. 

### Test Cases:
Input:5,7,16,1,2
Output:"false"

Input:3,5,-1,8,12
Output:"true"

## Letter Count 1
Have the function LetterCountI(str) take the str parameter being passed
and return the first word with the greatest number of repeated letters.
For example: "Today, is the greatest day ever!" should return greatest
because it has 2 e's (and 2 t's) and it comes before ever which also has
2 e's. If there are no words with repeating letters return -1. Words will
be separated by spaces. 

### Test Cases:
Input:"Hello apple pie"
Output:"Hello"

Input:"No words"
Output:-1

## Second Great Low
Have the function SecondGreatLow(arr) take the array of numbers stored in
arr and return the second lowest and second greatest numbers, respectively,
separated by a space. For example: if arr contains [7, 7, 12, 98, 106] the
output should be 12 98. The array will not be empty and will contain at
least 2 numbers. It can get tricky if there's just two numbers! 

### Test Cases:
Input:1, 42, 42, 180
Output:"42 42"

Input:4, 90
Output:"90 4"

## Division Stringified
Have the function DivisionStringified(num1,num2) take both parameters being
passed, divide num1 by num2, and return the result as a string with properly
formatted commas. If an answer is only 3 digits long, return the number with
no commas (ie. 2 / 3 should output "1"). For example: if num1 is 123456789
and num2 is 10000 the output should be "12,346". 

### Test Cases:
Input:5 & num2 = 2
Output:"3"

Input:6874 & num2 = 67
Output:"103"

## Counting Minutes 1
Have the function CountingMinutesI(str) take the str parameter being passed
which will be two times (each properly formatted with a colon and am or pm)
separated by a hyphen and return the total number of minutes between the
two times. The time will be in a 12 hour clock format. For example: if str
is 9:00am-10:00am then the output should be 60. If str is 1:00pm-11:00am
the output should be 1320. 

### Test Cases:
Input:"12:30pm-12:00am"
Output:690

Input:"1:23am-1:08am"
Output:1425

## Mean Mode
Have the function MeanMode(arr) take the array of numbers stored in arr
and return 1 if the mode equals the mean, 0 if they don't equal each other
(ie. [5, 3, 3, 3, 1] should return 1 because the mode (3) equals the mean
(3)). The array will not be empty, will only contain positive integers,
and will not contain more than one mode. 

### Test Cases:
Input:1, 2, 3
Output:0

Input:4, 4, 4, 6, 2
Output:1

## Dash Insert
Have the function DashInsert(str) insert dashes ('-') between each two
odd numbers in str. For example: if str is 454793 the output should be
4547-9-3. Don't count zero as an odd number.

### Test Cases:
Input:99946
Output:9-9-946

Input:56730
Output:567-30
