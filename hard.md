# Hard Challenges

## Pattern Chaser
Have the function PatternChaser(str) take str which will be a string and
return the longest pattern within the string. A pattern for this challenge
will be defined as: if at least 2 or more adjacent characters within the
string repeat at least twice. So for example "aabecaa" contains the pattern
aa, on the other hand "abbbaac" doesn't contain any pattern. Your program
should return yes/no pattern/null. So if str were "aabejiabkfabed" the
output should be yes abe. If str were "123224" the output should return no
null. The string may either contain all characters (a through z only),
integers, or both. But the parameter will always be a string type. The
maximum length for the string being passed in will be 20 characters. If a
string for example is "aa2bbbaacbbb" the pattern is "bbb" and not "aa".
You must always return the longest pattern possible. 

### Test Cases:
Input:"da2kr32a2"
Output:"yes a2"

Input:"sskfssbbb9bbb"
Output:"yes bbb"

## Gas Station
Have the function GasStation(strArr) take strArr which will be an an array
consisting of the following elements: N which will be the number of gas
stations in a circular route and each subsequent element will be the string
g:c where g is the amount of gas in gallons at that gas station and c will
be the amount of gallons of gas needed to get to the following gas station.
For example strArr may be: ["4","3:1","2:2","1:2","0:1"]. Your goal is to
return the index of the starting gas station that will allow you to travel
around the whole route once, otherwise return the string impossible. For
the example above, there are 4 gas stations, and your program should
return the string 1 because starting at station 1 you receive 3 gallons of
gas and spend 1 getting to the next station. Then you have 2 gallons + 2
more at the next station and you spend 2 so you have 2 gallons when you get
to the 3rd station. You then have 3 but you spend 2 getting to the final
station, and at the final station you receive 0 gallons and you spend your
final gallon getting to your starting point. Starting at any other gas
station would make getting around the route impossible, so the answer is 1.
If there are multiple gas stations that are possible to start at, return
the smallest index (of the gas station). N will be >= 2. 

### Test Cases:
Input:"4","1:1","2:2","1:2","0:1"
Output:"impossible"

Input:"4","0:1","2:2","1:2","3:1"
Output:"4"

## Parallel Sums
Have the function ParallelSums(arr) take the array of integers stored in
arr which will always contain an even amount of integers, and determine how
they can be split into two even sets of integers each so that both sets add
up to the same number. If this is impossible return -1. If it's possible to
split the array into two sets, then return a string representation of the
first set followed by the second set with each integer separated by a comma
and both sets sorted in ascending order. The set that goes first is the set
with the smallest first integer. 

For example: if arr is [16, 22, 35, 8, 20, 1, 21, 11], then your program
should output 1,11,20,35,8,16,21,22 

### Test Cases:
Input:1,2,3,4
Output:1,4,2,3

Input:1,2,1,5
Output:-1

## Bracket Combinations
Have the function BracketCombinations(num) read num which will be an
integer greater than or equal to zero, and return the number of valid
combinations that can be formed with num pairs of parentheses. For example,
if the input is 3, then the possible combinations of 3 pairs of parenthesis,
namely: ()()(), are ()()(), ()(()), (())(), ((())), and (()()). There are 5
total combinations when the input is 3, so your program should return 5. 

### Test Cases:
Input:3
Output:5

Input:2
Output:2
