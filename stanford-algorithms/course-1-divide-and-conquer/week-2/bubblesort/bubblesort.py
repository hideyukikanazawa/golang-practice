def bubblesort(arr):
    result = arr.copy()
    size = len(arr)
    changed = False
    for i in range(size):
        for j in range(size-i):
            if result[j] > result[j+1]:
                result[j], result[j+1] = result[j+1], result[j]
                changed = True
        if not changed:
            break
    return result


def main():
    arr = [1, 46, 7865, 23, 12, 3]
    final = bubblesort(arr)
    print(f"Initial: {arr}")
    print(f"Final: {final}")

if __name__ == "__main__":
    main()