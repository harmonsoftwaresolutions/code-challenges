# Solutions
## Binary Search Tree LCA
```
function BinarySearchTreeLCA(strArr) {
    //take the first item in the argument array and turn it into an array of integers
    let nodeArray = strArr[0]
        .replace(/[[]]/g, '')
        .split(/,s/)
        .map(val => parseInt(val, 10));
        
    //take the other items in the argument array and convert into integers
    let num1 = parseInt(strArr[1], 10);
    let num2 = parseInt(strArr[2], 10);
    
    //find the positions of the given numbers in the number array
    let ind1 = nodeArray.findIndex(val => val === num1);
    let ind2 = nodeArray.findIndex(val => val === num2);
    
    //determine the farther of the two positions, we are not interested in elements past that
    let rightEdge = Math.max(ind1, ind2);
    
    //see if there are any items to the left of rightEdge that split the two given numbers, that will be the answer
    let result = nodeArray.filter((val, ind) => (val >= Math.min(num1, num2) && val <= Math.max(num1, num2) && ind <= rightEdge));
    
    //if not any, then return the item that is farthest to the left
    if (result.length === 0) return ind1 < ind2 ? strArr[1] : strArr[2];
    
    //if there is, then return it as a string
    return result[0].toString();
}
   
// keep this function call here 
BinarySearchTreeLCA(readline());
```
