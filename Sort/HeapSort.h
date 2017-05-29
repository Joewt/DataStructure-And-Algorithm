//
// Created by joe on 17-5-28.
//

#ifndef HEAP_HEAPSORT_H
#define HEAP_HEAPSORT_H
#include"Heap.h"
template<typename T>
void heapSort1(T arr[],int n){
    MaxHeap<int> maxheap=MaxHeap<T>(n);
    for(int i=0;i<n;i++)
        maxheap.insert(arr[i]);
    for(int i=0;i<n;i++)
        arr[i]=maxheap.extractMax();
}
template<typename T>
void heapSort2(T arr[], int n){

    MaxHeap<T> maxheap = MaxHeap<T>(arr,n);
    for( int i = n-1 ; i >= 0 ; i-- )
        arr[i] = maxheap.extractMax();

}
template <typename T>
void __shiftDown(T arr[],int n,int k){
    T e=arr[k];
    while(2*k+1<n){
        int j=2*k+1;
        if(j+1<n&&arr[j+1]>arr[j])
            j++;
        if(e>=arr[j])break;
        arr[k]=arr[j];
        k=j;
    }
    arr[k]=e;
}
template<typename T>
void heapSort(T arr[],int n){
    for(int i=(n-1)/2;i>=0;i--)
        __shiftDown(arr,n,i);
    for(int i=n-1;i>0;i--){
        swap(arr[0],arr[i]);
        __shiftDown(arr,i,0);
    }
}
#endif //HEAP_HEAPSORT_H
