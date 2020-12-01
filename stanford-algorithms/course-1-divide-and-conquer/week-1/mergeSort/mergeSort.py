def mergeSort(arr):
    size = len(arr)
    if size < 2:
        return arr

    mid = size // 2
    return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))


def merge(left, right):
    size, i, j = len(left) + len(right), 0, 0
    result = []

    for _ in range(size):
        if i > len(left) - 1 and j <= len(right) - 1:
            result.append(right[j])
            j+=1
        elif j > len(right) - 1 and i <= len(left) - 1:
            result.append(left[i])
            i+=1
        elif left[i] <= right[j]:
            result.append(left[i])
            i+=1
        else:
            result.append(right[j])
            j+=1
    
    return result 


def main():
    array = [1, 2, 5, 7, 890, 10000, 123, 3, 32, 5, 75, 97, 975, 3234, 123]
    sorted = mergeSort(array)
    print(f"Initial array:\t{array}")
    print(f"Final array:\t{sorted}")


if __name__ == "__main__":
    main()