# 1 = append to stack
# 2 = delete from stack
# 3 = print

def getMax(operations):
    stack = []
    tempArr = []

    for _ in operations:
        q = list(map(int, _.split()))     # q = [1 27]

        # check the query
        if q[0] == 1:
            if stack:
                stack.append(max(stack[-1], q[1]))
            else:
                stack.append(q[1])

        elif q[0] == 2:
            stack.pop()
        elif q[0] == 3:
            tempArr.append(stack[-1])
    return tempArr

def equalStack(h1, h2, h3):

    h1 = h1[::-1]
    h2 = h2[::-1]
    h3 = h3[::-1]

    sum1 = sum(h1)
    sum2 = sum(h2)
    sum3 = sum(h3)
    
    while True:
        minHeight = min(sum1, sum2, sum3)

        if minHeight == 0:
            return 0
        
        if sum1 > minHeight:
            sum1 -= h1.pop()
        if sum2 > minHeight:
            sum2 -= h2.pop()
        if sum3 > minHeight:
            sum3 -= h3.pop()

        if sum1 == sum2 == sum3:
            return sum2

    print(sum1, sum2, sum3)

    