def mergeSort(arr):
    if len(arr) == 1:
        return arr
    mid = len(arr) // 2
    return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))


def merge(left, right):
    size, i, j = len(left) + len(right), 0, 0
    result = []
    
    for _ in range(size):
        if i > len(left) - 1 and j < len(right):
            result.append(right[j])
            j += 1
        elif i < len(left) and j > len(right) - 1:
            result.append(left[i])
            i += 1
        elif left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    return result


def main():
    arr = [1, 3, 6, 8, 1, 12312, 4, 3, 123]
    sorted = mergeSort(arr)
    print(f"Initial array:\t{arr}")
    print(f"Final array:\t{sorted}")


if __name__ == "__main__":
    main()