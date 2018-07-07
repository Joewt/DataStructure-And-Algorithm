cmake_minimum_required(VERSION 3.3)
project(Search)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11")

set(SOURCE_FILES main.cpp BinarySearch.h FileOps.h SequenceST.h)
add_executable(Search ${SOURCE_FILES})