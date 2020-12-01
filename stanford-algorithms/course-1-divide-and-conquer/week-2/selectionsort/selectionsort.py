def selectionsort(arr):
    for i in range(len(arr)):
        min = i
        for j in range(i, len(arr)):
            if arr[j] < arr[min]:
                min = j
        arr[min], arr[i] = arr[i], arr[min]
    return arr

def main():
    arr = [134, 6, 24, 1, 3213]
    print(f"Initial:\t\t{arr}")
    selectionsort(arr)
    print(f"Final:\t\t\t{arr}")


if __name__ == "__main__":
    main()