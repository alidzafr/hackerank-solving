# self.info
# left right
class Node:
    def __init__(self, info):
        self.info = info  
        self.left = None  
        self.right = None 
        self.level = None 

    def __str__(self):
        return str(self.info)
    
    def preOrder(root):
        if root:
            print(root.info, end=" ")
            preOrder(root.left)
            preOrder(root.right)
        return
    
class BinarySearchTree:
    def __init__(self): 
        self.root = None
    
    def insert(self, val):
        if self.root is None:
            self.root = Node(val)

        temp = self.root
        while self.root:
            if temp < val:
                if temp.right is None:
                    temp.right = Node(val)
                    return self.root
                else:
                    temp = temp.right
            else:
                if temp.left is None:
                    temp.left = Node(val)
                    return self.root
                else:
                    temp = temp.left
        return self.root
