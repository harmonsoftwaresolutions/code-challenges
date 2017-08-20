#  Have the function WordCount(str) take the str string parameter being passed
#  and return the number of words the string contains
#  (e.g. "Never eat shredded wheat or cake" would return 6). Words will be
#  separated by single spaces. 

def WordCount(s):
	li = s.split()
	return len(li)

def test1():
	assert WordCount("hello world") == 2
	assert WordCount("one 22 three") == 3
