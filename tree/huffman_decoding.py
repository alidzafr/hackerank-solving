def decodeHuff(root, s):
	#Enter Your Code Here
    # Temporary pointer
    temp = root
    result = []
    
    # traverse binary coded string
    for char in s:
        if char == '0':
            temp = temp.left
        else:
            temp = temp.right
        
        # Check leaf node
        if temp.left is None and temp.right is None:
            result.append(temp.data)
            temp = root
    print("".join(result))

def lca(root, v1, v2):
    temp = root
    min = []
    result = []

    if root is None:
        return 0
    if root is not None:
        


