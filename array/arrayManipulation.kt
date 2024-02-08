fun arrayManipulation(n: Int, queries: Array<Array<Int>>): Long {
    // Write your code here
    var resultArray : ArrayList<Long> = ArrayList()

    for (i in 0..n-1){
        resultArray.add(0)
    }
    for (array in queries){
        val firstIndex = array[0]-1
        val lastIndex = array[1]-1
        val valueK = array[2]
        for (index in firstIndex..lastIndex){
            val valueArray = resultArray[index]
            
            val value = valueArray + valueK.toLong()
            
            resultArray.set(index = index, element = value)
        }
    }
    
    return resultArray.maxOrNull() ?: 0
}