#  Have the function SimpleSymbols(str) take the str parameter being passed and
#  determine if it is an acceptable sequence by either returning the string
#  true or false. The str parameter will be composed of + and = symbols with
#  several letters between them (ie. ++d+===+c++==a) and for the string to be
#  true each letter must be surrounded by a + symbol. So the string to the left
#  would be false. The string will not be empty and will have at least one
#  letter.

def SimpleSymbols1(s):
	test = None
	ln = len(s) - 1
	for i, c in enumerate(s):
		if c.isalpha() == True and test != "false":
			if i == 0 or s[i-1] != '+':
				test = "false"
			elif i == ln or s[i+1] != '+':
				test = "false"
			else:
				test = "true"
	return test

def test1():
	assert SimpleSymbols1("f++d+") == "false"
	assert SimpleSymbols1("+d===+a+") == "false"
	assert SimpleSymbols1("+d+") == "true"

def SimpleSymbols2(s):
	# append to string to avoid any out of index errors
	s = '=' + s + '='
	li = list(s)
	for i in range(0, len(li)):

		if li[i].isalpha():
			if li[i-1] != '+' or li[i+1] != '+':
				return "false"
	return "true"

def test2():
	assert SimpleSymbols2("f++d+") == "false"
	assert SimpleSymbols2("+d===+a+") == "false"
	assert SimpleSymbols2("+d+") == "true"
