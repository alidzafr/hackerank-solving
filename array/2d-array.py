tempArr = list()
    for i in range(4):
        for j in range(4):
            sum = arr[i][0+j] + arr[i][1+j] + arr[i][2+j]+arr[i+1][1+j]+arr[i+2][0+j] + arr[i+2][1+j] + arr[i+2][2+j]
            tempArr.append(sum)
    print(arr)
    print(tempArr)
    return max(tempArr)