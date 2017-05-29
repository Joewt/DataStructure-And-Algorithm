cmake_minimum_required(VERSION 3.3)
project(Sort)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11")

set(SOURCE_FILES main.cpp InsertionSort.h SortTestHelper.h MergeSort.h QuickSort.h Heap.h HeapSort.h)
add_executable(Sort ${SOURCE_FILES})
