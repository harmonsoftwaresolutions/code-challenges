#  Have the function SwapCase(str) take the str parameter and swap the case
#  of each character. For example: if str is "Hello World" the output should
#  be hELLO wORLD. Let numbers and symbols stay the way they are.

def SwapCase(s):
	li = list(s)
	newli = []
	for c in li:
		a = c
		if c.isalpha():
			a = c.swapcase()
		newli.append(a)
	s = ''.join(newli)
	return s

def test1():
	assert SwapCase("Hello-LOL") == "hELLO-lol"
	assert SwapCase("Sup DUDE!!?") == "sUP dude!!?"
