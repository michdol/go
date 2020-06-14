
def selection_sort(array):
  comparisons = 0
  exchanges = 0
  length = len(array)
  for i in range(0, length):
    minimum = i
    for j in range(i + 1, length):
      comparisons += 1
      if array[j] < array[minimum]:
        minimum = j
    exchanges += 1
    array[minimum], array[i] = array[i], array[minimum]
  print("selection_sort -> comparisons: {}, exchanges: {}".format(comparisons, exchanges))


def insertion_sort(array):
  comparisons = 0
  exchanges = 0
  length = len(array)
  for i in range(0, length):
    j = i - 1
    key = array[i]
    while j >= 0 and key < array[j]:
      array[j + 1] = array[j]
      j -= 1
      comparisons += 1
    # Initial comparison
    comparisons += 1
    exchanges += 1
    array[j + 1] = key
  print("insertion_sort -> comparisons: {}, exchanges: {}".format(comparisons, exchanges))


def selection_sort_test():
  input_string = "SORTEXAMPLE"
  print(input_string)
  array = [ch for ch in input_string]
  selection_sort(array)
  result = "".join(array)
  print(result)
  assert "".join(sorted(input_string)) == result


def insertion_sort_test():
  input_string = "SORTEXAMPLE"
  print(input_string)
  array = [ch for ch in input_string]
  insertion_sort(array)
  result = "".join(array)
  print(result)
  assert "".join(sorted(input_string)) == result


def main():
  selection_sort_test()
  insertion_sort_test()


if __name__ == "__main__":
  main()
