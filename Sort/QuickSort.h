//
// Created by joe on 17-5-28.
//

#ifndef MERGE_SORT_QUICKSORT_H
#define MERGE_SORT_QUICKSORT_H
#include<iostream>
#include<ctime>
#include<algorithm>
#include"InsertionSort.h"
using namespace std;
template<typename T>
int __partation(T arr[],int l,int r){
    swap(arr[l],arr[rand()%(r-l+1)+l]);
    T v=arr[l];
    int i=l+1,j=r;
    while(true){
        while(i<=r&&arr[i]<v)
            i++;
        while(j>=l+1&&arr[j]>v)
            j--;
        if(i>j)
            break;
        swap(arr[i],arr[j]);
        i++;
        j--;
    }
}
template<typename T>
int __partation2(T arr[],int l,int r){
    swap(arr[l],arr[rand()%(r-l+1)+l]);
    T v=arr[l];
    int j=l;
    for(int i=l+1;i<=r;i++)
        if(arr[i]<v){
            swap(arr[i],arr[j]);
            j++;
        }
    swap(arr[l],arr[j]);
    return j;

}
template<typename T>
void __quickSort(T arr[],int l,int r){

    if(r-l<15){
        insertionSort(arr,l,r);
        return;
    }
    int p=__partation(arr,l,r);
    __quickSort(arr,l,p-1);
    __quickSort(arr,p+1,r);

}
template<typename T>
void quickSort(T arr[],int n)
{
    srand(time(NULL));
    __quickSort(arr,0,n-1);
}

template<typename T>
void __quickSort3Way(T arr[],int l,int r){
    if(r-l<15){
        insertionSort(arr,l,r);
        return;
    }

    swap(arr[l],arr[rand()%(r-l+1)+l]);
    T v=arr[l];
    int lt=l;
    int gt=r+1;
    int i=l+1;
    while(i<gt){
        if(arr[i]<v){
            swap(arr[lt++],arr[i++]);
        }else if(arr[i]>v){
            swap(arr[gt--],arr[i++]);
        }else{
            i++;
        }
    }
    swap(arr[l],arr[lt]);
    __quickSort(arr,l,lt-1);
    __quickSort(arr,gt,r);

}
template<typename T>
void quickSort3Way(T arr[],int n){
    srand(time(NULL));
    __quickSort3Way(arr,0,n-1);
}


#endif //MERGE_SORT_QUICKSORT_H
