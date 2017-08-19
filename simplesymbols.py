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
