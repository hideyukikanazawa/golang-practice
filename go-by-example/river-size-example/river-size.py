def getRiverSizes(matrix):
    marked = set()
    riverSizes = []
    for row in range(len(matrix)):
        for col in range(len(matrix[row])):
            if matrix[row][col] == 1 and (row, col) not in marked:
                marked.add((row, col))
                riverSize = 1
                nodes = [(row, col)]
                while nodes:
                    current = nodes.pop()
                    neighbours = getNeighbours(current, matrix)
                    for nb in neighbours:
                        y, x = nb
                        if nb not in marked and matrix[y][x] == 1:
                            marked.add(nb)
                            nodes.append(nb)
                            riverSize += 1
                riverSizes.append(riverSize)
    return riverSizes


def getNeighbours(coordinate, matrix):
    y, x = coordinate
    ns = []
    if y > 0:
        ns.append((y-1, x))
    if y < len(matrix) - 1:
        ns.append((y+1, x))
    if x > 0:
        ns.append((y, x-1))
    if x < len(matrix[y]) - 1:
        ns.append((y, x+1))
    return ns


def main():
    matrixA = [
        [1, 1, 1, 1, 0],
        [1, 0, 1, 0, 0],
        [1, 0, 1, 0, 1],
        [1, 0, 1, 0, 1],
        [1, 0, 1, 1, 0]
    ]
    riverSizes = getRiverSizes(matrixA)
    print(f"River sizes in matrix A are: {riverSizes}")

if __name__ == "__main__":
    main()