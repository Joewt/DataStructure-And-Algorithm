//
// Created by joe on 17-5-30.
//

#ifndef SEARCH_BINARYSEARCH_H
#define SEARCH_BINARYSEARCH_H

template<typename T>
int binarySearch(T arr[],int n,T target){
    int l=0,r=n-1;
    while(l<=r){
        int mid=l+(r-l)/2;
        if(arr[mid]==target){
            return mid;
        }
        if(target>arr[mid])
            l=mid+1;
        else
            r=mid-1;
    }
    return -1;
}
template<typename T>
int __binarySearch2(T arr[],int l,int r,int target){
    if(l>r)
        return -1;
    int mid=l+(r-l)/2;
    if(target==arr[mid])
        return mid;
    else if(target>arr[mid])
        return __binarySearch2(arr,mid+1,r,target);
    else
        return __binarySearch2(arr,l,mid-1,target);
}

template<typename T>
int binarySearch2(T arr[],int n,T target){
    return __binarySearch2(arr,0,n-1,target);
}

#endif //SEARCH_BINARYSEARCH_H
