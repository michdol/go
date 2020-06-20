

def merge_sort(array, left, right):
  if left >= right:
    return
  middle = (left + right) // 2
  merge_sort(array, left, middle)
  merge_sort(array, middle + 1, right)
  merge(array, left, right, middle)


def merge(array, left, right, middle):
  left_half = array[left:middle + 1]
  right_half = array[middle + 1:right + 1]
  length_left = len(left_half)
  length_right = len(right_half)

  i = 0
  j = 0
  k = 0

  while i < length_left and j < length_right:
    if left_half[i] <= right_half[j]:
      array[k] = left_half[i]
      i += 1
    else:
      array[k] = right_half[j]
      j += 1
    k += 1

  while i < length_left:
    array[k] = left_half[i]
    i += 1
    k += 1

  while j < length_right:
    array[k] = right_half[j]
    j += 1
    k += 1


def main():
  arr1 = [5, 3, 6, 2, 4, 1]
  # arr2 = [5, 3, 6, 2, 4, 1, 7]
  merge_sort(arr1, 0, len(arr1))
  print(arr1)


if __name__ == "__main__":
  main()
