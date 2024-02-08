from collections import deque

def minimumMoves(grid, startX, startY, goalX, goalY):
    c = 0
    visit = set()
    q = deque([startX, startY, c]) #stack moves FIFO
    moves = [[1,0], [-1,0], [0,1], [0,-1]]

    # Main logic
    while q:
        pathX, pathY, c = q.popleft()
        c += 1

        for xi, yi in moves:
            x, y = pathX, pathY
            while True:
                x, y = x+xi, y+yi
                # Check Condition
                if 0 <= x < len(grid) and 0 <= y < len(grid[0]) and grid[x][y] == '.':
                    if x == goalX and y == goalY:
                        return c
                    elif (x,y) not in visit:
                        visit.add((x,y))
                        q.append([x,y,c])
                else:
                    break
    return -1