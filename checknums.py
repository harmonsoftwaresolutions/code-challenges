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
