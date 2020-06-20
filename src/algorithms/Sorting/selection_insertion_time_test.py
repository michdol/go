import sys
import timeit

from typing import Tuple


def selection_sort(array):
  length = len(array)
  for i in range(0, length):
    minimum = i
    for j in range(i + 1, length):
      if array[j] < array[minimum]:
        minimum = j
    array[minimum], array[i] = array[i], array[minimum]


def insertion_sort(array):
  length = len(array)
  for i in range(0, length):
    j = i - 1
    key = array[i]
    while j >= 0 and key < array[j]:
      array[j + 1] = array[j]
      j -= 1
    array[j + 1] = key


def main(algorithm: str, input_size: str):
  setup = '''
from input_integer_arrays import input_100, input_1000, input_10000
input = input_{input}()
from __main__ import selection_sort, insertion_sort
alg = {algorithm}
'''.format(input=input_size, algorithm=algorithm)
  print(timeit.Timer('arr = input[:];alg(arr)', setup=setup).repeat(1, 1))


def get_args(*args) -> Tuple[str, str]:
  if len(args) < 3:
    raise ValueError("Need algorithm name and test input size\nAlg(insert, select), InputSize(100, 1000)")
  algorithm: str = ""
  if args[1] == "i":
    algorithm = "insertion_sort"
  else:
    algorithm = "selection_sort"
  input_size: str = args[2]
  return algorithm, input_size


if __name__ == "__main__":
  input_args = get_args(*sys.argv)
  main(*input_args)
