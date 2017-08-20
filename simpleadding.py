#  Have the function SimpleAdding(num) add up all the numbers from 1 to num.
#  For example: if the input is 4 then your program should return 10 because
#  1 + 2 + 3 + 4 = 10. For the test cases, the parameter num will be any number
#  from 1 to 1000.

def SimpleAdding(num):
	sum = 0
	for x in range(1, num + 1):
		print (x)
		sum += x
		print (sum)
	return sum

def test1():
	assert SimpleAdding(4) == 10
	assert SimpleAdding(12) == 78
	assert SimpleAdding(140) == 9870
