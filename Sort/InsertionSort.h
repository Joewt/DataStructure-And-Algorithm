//
// Created by joe on 17-5-27.
//

#ifndef MERGE_SORT_INSERTIONSORT_H
#define MERGE_SORT_INSERTIONSORT_H
#include<iostream>
#include<cassert>
using namespace std;
template<typename T>
void insertionSort(T arr[],int n){
    for(int j=1;j<n;++j) {
        T e = arr[j];
        int i;
        for (i = j; i > 0 && arr[i - 1] > e; --i)
            arr[i] = arr[i - 1];
        arr[i] = e;
    }
    return;
}
template<typename T>
void insertionSort(T arr[],int l,int r){
    for(int j=l+1;j<=r;++j){
        T e=arr[j];
        int i;
        for(i=j;i>l&&arr[i-1]>e;--i)
            arr[i]=arr[i-1];
        arr[i]=e;
    }
    return;
}
#endif //MERGE_SORT_INSERTIONSORT_H
