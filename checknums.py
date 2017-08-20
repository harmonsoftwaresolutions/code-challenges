#  Have the function SimpleAdding(num) add up all the numbers from 1 to num.
#  For example: if the input is 4 then your program should return 10 because
#  1 + 2 + 3 + 4 = 10. For the test cases, the parameter num will be any number
#  from 1 to 1000.

def CheckNums(num1,num2): 
    r = None
    if num1 == num2:
        r = "-1"
    elif num1 > num2:
        r = "false"
    else:
        r = "true"
    return r

def test1():
	assert CheckNums(45, 154) == "true"
	assert CheckNums(63, 16) == "false"
	assert CheckNums(14, 14) == "-1"
