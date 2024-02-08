import collections


def levelOrderTraversal(self, root: TreeNode) -> List[List[int]]:
    res = []

    q = collections.dequeue()
    q.append(root)

    while q:
        qlen = len(q)
        level = []
        for i in range(qlen):
            node = q.popleft()
            if node:
                level.append(node.val)
                q.append(node.left)
                q.append(node.right)
        if level:
            res.append(level)
    return res

def levelOrder(root):
    q = deque([root])
    while q:
        node = q.popleft()
        if node:
            print(node.info, end=" ")
            q.append(node.left)
            q.append(node.right)