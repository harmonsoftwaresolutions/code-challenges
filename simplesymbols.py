#  Have the function SimpleSymbols(str) take the str parameter being passed and
#  determine if it is an acceptable sequence by either returning the string
#  true or false. The str parameter will be composed of + and = symbols with
#  several letters between them (ie. ++d+===+c++==a) and for the string to be
#  true each letter must be surrounded by a + symbol. So the string to the left
#  would be false. The string will not be empty and will have at least one
#  letter.

def SimpleSymbols(s):
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
	assert SimpleSymbols("f++d+") == "false"
	assert SimpleSymbols("+d===+a+") == "false"
	assert SimpleSymbols("+d+") == "true"
