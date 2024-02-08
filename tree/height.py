def height(root):
    if root is None:
        return -1
    if root is not None:
        temp_hl = height(root.left)
        temp_hr = height(root.right)
        return max(temp_hl, temp_hr)+1