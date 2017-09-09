# Medium Challenges

## Binary Search Tree LCA
Have the function BinarySearchTreeLCA(strArr) take the array of strings
stored in strArr, which will contain 3 elements: the first element will
be a binary search tree with all unique values in a preorder traversal
array, the second and third elements will be two different values, and
your goal is to find the lowest common ancestor of these two values. For
example: if strArr is ["[10, 5, 1, 7, 40, 50]", "1", "7"] then this tree
looks like the following: 

 

For the input above, your program should return 5 because that is the
value of the node that is the LCA of the two nodes with values 1 and 7.
You can assume the two nodes you are searching for in the tree will
exist somewhere in the tree. 

### Test Cases:
Input:"[10, 5, 1, 7, 40, 50]", "5", "10"
Output:10

Input:"[3, 2, 1, 12, 4, 5, 13]", "5", "13"
Output:12

## Max Subarray
Have the function MaxSubarray(arr) take the array of numbers stored in
arr and determine the largest sum that can be formed by any contiguous
subarray in the array. For example, if arr is [-2, 5, -1, 7, -3] then
your program should return 11 because the sum is formed by the subarray
[5, -1, 7]. Adding any element before or after this subarray would make
the sum smaller. 

### Test Cases
Input:1, -2, 0, 3
Output:3

Input:3, -1, -1, 4, 3, -1
Output:8

## K Unique Characters
Have the function KUniqueCharacters(str) take the str parameter being
passed and find the longest substring that contains k unique characters,
where k will be the first character from the string. The substring will
start from the second position in the string because the first character
will be the integer k. For example: if str is "2aabbacbaa" there are
several substrings that all contain 2 unique characters, namely:
["aabba", "ac", "cb", "ba"], but your program should return "aabba"
because it is the longest substring. If there are multiple longest
substrings, then return the first substring encountered with the longest
length. k will range from 1 to 6. 

### Test Cases:
Input:"3aabacbebebe"
Output:"cbebebe"

Input:"2aabbcbbbadef"
Output:"bbcbbb"

## Binary Tree
Have the function BinaryTreeLCA(strArr) take the array of strings stored
in strArr, which will contain 3 elements: the first element will be a
binary tree with all unique values in a format similar to how a binary
heap is implemented with NULL nodes at any level represented with a #,
the second and third elements will be two different values, and your goal
is to find the lowest common ancestor of these two values. For example:
if strArr is ["[12, 5, 9, 6, 2, 0, 8, #, #, 7, 4, #, #, #, #]", "6", "4"]
then this tree looks like the following: 

 

For the input above, your program should return 5 because that is the
value of the node that is the LCA of the two nodes with values 6 and 4.
You can assume the two nodes you are searching for in the tree will exist
somewhere in the tree. 

### Test Cases:
Input:"[5, 2, 6, 1, #, 8, #]", "2", "6"
Output:5

Input:"[5, 2, 6, 1, #, 8, 12, #, #, #, #, #, #, 3, #]", "3", "12"
Output:12
