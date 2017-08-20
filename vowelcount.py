#  Have the function VowelCount(str) take the str string parameter being
#  passed and return the number of vowels the string contains
#  (ie. "All cows eat grass and moo" would return 8). Do not count y as a
#  vowel for this challenge. 

def VowelCount(s):
    cnt = 0
    li = list(s)
    vw = ['a','e','i','o','u']
    for c in li:
        if c in vw:
            cnt += 1
    return cnt

def test1():
	assert VowelCount("hello") == 2
	assert VowelCount("coderbyte") == 3
	assert VowelCount("quaaaaiizick") == 8
