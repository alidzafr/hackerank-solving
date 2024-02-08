def topView(root):
    q=[]            #Queue
    d=dict()
    root.level=0    ##vari(-1 +1) level
    q.append(root)
    while len(q) !=0:
        root=q.pop(0)
        if root.level not in d:
            d[root.level]=root.info 
        if root.left is not None:
            q.append(root.left)
            root.left.level = root.level-1
        if root.right is not None:
            q.append(root.right)
            root.right.level = root.level+1
    print(d.values)
    # print(*d.values)

    for i in sorted(d):         #left to right
        print(d[i], end=' ')