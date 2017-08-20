import math

#  Have the function TimeConvert(num) take the num parameter being passed and
#  return the number of hours and minutes the parameter converts to
#  (ie. if num = 63 then the output should be 1:3). Separate the number of
#  hours and minutes with a colon. 

def TimeConvert(num): 
	h = math.floor(num / 60)
	m = num % 60
	t = str(h) + ':' + str(m)
	return t

def test1():
	assert TimeConvert(63) == "1:3"
	assert TimeConvert(45) == "0:45"
	assert TimeConvert(126) == "2:6"
