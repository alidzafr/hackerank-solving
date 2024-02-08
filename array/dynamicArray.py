queries = [][]int{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}

n = 2

tempArr = [[]for i in range(n)]
returnSeq = []
lastAnswer = 0

for q in queries:
    tipe = q[0]
    x = q[1]
    y = q[2]

    sequenceNum = (x^lastAnswer)%n
    if tipe == 1:
        tempArr[sequenceNum].append(y)
    elif tipe == 2:
        indx = y % len(tempArr[sequenceNum])
        lastAnswer = tempArr[sequenceNum][indx]
        returnSeq = append(lastAnswer)
return returnSeq
