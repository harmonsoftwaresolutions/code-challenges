#  Have the function AlphabetSoup(str) take the str string parameter being
#  passed and return the string with the letters in alphabetical order
#  (ie. hello becomes ehllo). Assume numbers and punctuation symbols will not
#  be included in the string.

def AlphabetSoup(s):
	li = list(s)
	for i, c in enumerate(li):
		if li[i].isalpha() != True:
			li.remove(c)
	li.sort()
	r = ''.join(li)
	return r

def test1():
	assert AlphabetSoup("coderbyte") == "bcdeeorty"
	assert AlphabetSoup("coderbyte4") == "bcdeeorty"
