def mergesortinversions(arr):
    if len(arr) == 1:
        return arr, 0
    else:
        mid = len(arr) // 2
        left, right = arr[:mid], arr[mid:]

        left, ai = mergesortinversions(left)
        right, bi = mergesortinversions(right)
        result = []

        i, j = 0, 0
        inversions = 0 + ai + bi

        while i < len(left) and j < len(right):
            if left[i] < right[j]:
                result.append(left[i])
                i += 1
            else: 
                result.append(right[j])
                j += 1
                inversions += (len(left)-i)
        result += left[i:]
        result += right[j:]
    return result, inversions


def main():
    with open("D:/Downloads/IntegerArray.txt", "r") as istr:
        output = [int(x.strip("\n")) for x in istr.readlines()]
    final, inv = mergesortinversions(output)
    print(f"Final:\t\t{final}")
    print(f"Total no. of inversions: {inv}")

if __name__ == "__main__":
    main()