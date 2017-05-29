#include <iostream>
#include<cassert>
#include<cmath>
#include<cassert>
#include<string>
#include<ctime>
#include<algorithm>
#include<typeinfo>
#include"InsertionSort.h"
#include"MergeSort.h"
#include"QuickSort.h"
#include"SortTestHelper.h"
#include"HeapSort.h"
#include"Heap.h"
using namespace std;
template<typename Item>
class IndexMaxHeap{
private:
    Item *data;
    int *indexes;
    int *reverse;
    int count;
    int capacity;
    void shiftUp(int k){
        Item ret=data[k];
        while(k>1&&data[indexes[k/2]]<data[indexes[k]]){
            indexes[k/2]=indexes[k];
            reverse[indexes[k/2]]=k/2;
            reverse[indexes[k]]=k;
            k/=2;
        }
        data[k]=ret;
    }
    void shiftDown(int k){
        Item ret=data[k];
        while(2*k<=count) {
            int j = 2 * k;
            if(j+1<=count&&data[indexes[j+1]]>data[indexes[j]])
                j++;
            if(data[indexes[k]]>data[indexes[j]])break;
            swap(indexes[k],indexes[j]);
            reverse[indexes[k]]=k;
            reverse[indexes[j]]=j;
            k=j;
        }
    }
public:
    IndexMaxHeap(int capacity){
        data=new Item[capacity+1];
        indexes=new int[capacity+1];
        reverse=new int[capacity+1];
        for(int i=0;i<capacity;i++)
            reverse[i]=0;
        count=0;
        this->capacity=capacity;
    }
    ~MaxHeap(){
        delete[] data;
        delete[] indexes;
        delete[] reverse;
    }
    int size()
    {
        return count;
    }
    bool isEmpty(){
        return count==0;
    }
    void insert(int i,Item item){
        assert(count+1<=capacity);
        assert(i+1>=1&&i+1<=capacity);
        i+=1;
        data[count+1]=item;
        indexes[count+1]=i;
        reverse[i]=count+1;
        count++;
        shiftUp(count);
    }
    Item extractMax(){
        assert(count>0);
        Item ret=data[indexes[1]];
        swap(indexes[1],indexes[count]);
        reverse[indexes[1]]=1;
        reverse[indexes[count]]=0;
        count--;
        shiftDown(1);
        return ret;
    }
    int extractMaxIndex(){
        assert(count>0);
        int ret=indexes[1]-1;
        swap(indexes[1],indexes[count]);
        reverse[indexes[1]]=1;
        reverse[indexes[count]]=0;
        count--;
        shiftDown(1);
        return ret;
    }
    bool contain(int i){
        assert(i+1>=1&&i+1<=capacity);
        return reverse[i+1]!=0;
    }
    Item getItem(int i){
        assert(contain(i));
        return data[i+1];
    }
    void change(int i,Item newItem){
        assert(contain(i));
        i+=1;
        data[i]=newItem;
//        for(int j=1;j<count;j++){
//            if(indexes[j]==i){
//                shiftUp(j);
//                shiftDown(j);
//
//            }
//        }
        int j=reverse[i];
        shiftUp(j);
        shiftDown(j);
    }
};


int main() {

    int n = 1000000;
    // 测试1 一般性测试
    cout<<"Test for Random Array, size = "<<n<<", random range [0, "<<n<<"]"<<endl;
    int* arr1 = SortTestHelper::generateRandomArray(n,0,n);
    int* arr2 = SortTestHelper::copyIntArray(arr1, n);
    int* arr3 = SortTestHelper::copyIntArray(arr1, n);
    int* arr4 = SortTestHelper::copyIntArray(arr1, n);
    int* arr5 = SortTestHelper::copyIntArray(arr1, n);
    int* arr6 = SortTestHelper::copyIntArray(arr1, n);

    SortTestHelper::testSort("Merge Sort", mergeSort, arr1, n);
    SortTestHelper::testSort("Quick Sort", quickSort, arr2, n);
    SortTestHelper::testSort("Quick Sort 3 Ways", quickSort3Way, arr3, n);
    SortTestHelper::testSort("Heap Sort 1", heapSort1, arr4, n);
    SortTestHelper::testSort("Heap Sort 2", heapSort2, arr5, n);
    SortTestHelper::testSort("Heap Sort ", heapSort, arr5, n);

    delete[] arr1;
    delete[] arr2;
    delete[] arr3;
    delete[] arr4;
   // delete[] arr5;

    cout<<endl;


    // 测试2 测试近乎有序的数组
    int swapTimes = 100;
    cout<<"Test for Random Nearly Ordered Array, size = "<<n<<", swap time = "<<swapTimes<<endl;
    arr1 = SortTestHelper::generateNearlyOrderedArray(n,swapTimes);
    arr2 = SortTestHelper::copyIntArray(arr1, n);
    arr3 = SortTestHelper::copyIntArray(arr1, n);
    arr4 = SortTestHelper::copyIntArray(arr1, n);
    arr5 = SortTestHelper::copyIntArray(arr1, n);
    arr6 = SortTestHelper::copyIntArray(arr1, n);

    SortTestHelper::testSort("Merge Sort", mergeSort, arr1, n);
    SortTestHelper::testSort("Quick Sort", quickSort, arr2, n);
    SortTestHelper::testSort("Quick Sort 3 Ways", quickSort3Way, arr3, n);
    SortTestHelper::testSort("Heap Sort 1", heapSort1, arr4, n);
    SortTestHelper::testSort("Heap Sort 2", heapSort2, arr5, n);
    SortTestHelper::testSort("Heap Sort ", heapSort, arr5, n);

    delete[] arr1;
    delete[] arr2;
    delete[] arr3;
    delete[] arr4;
    delete[] arr5;
    delete[] arr6;

    cout<<endl;


    // 测试3 测试存在包含大量相同元素的数组
    cout<<"Test for Random Array, size = "<<n<<", random range [0,10]"<<endl;
    arr1 = SortTestHelper::generateRandomArray(n,0,10);
    arr2 = SortTestHelper::copyIntArray(arr1, n);
    arr3 = SortTestHelper::copyIntArray(arr1, n);
    arr4 = SortTestHelper::copyIntArray(arr1, n);
    arr5 = SortTestHelper::copyIntArray(arr1, n);
    arr6 = SortTestHelper::copyIntArray(arr1, n);

    SortTestHelper::testSort("Merge Sort", mergeSort, arr1, n);
    SortTestHelper::testSort("Quick Sort", quickSort, arr2, n);
    SortTestHelper::testSort("Quick Sort 3 Ways", quickSort3Way, arr3, n);
    SortTestHelper::testSort("Heap Sort 1", heapSort1, arr4, n);
    SortTestHelper::testSort("Heap Sort 2", heapSort2, arr5, n);
    SortTestHelper::testSort("Heap Sort ", heapSort, arr5, n);

    delete[] arr1;
    delete[] arr2;
    delete[] arr3;
    delete[] arr4;
    delete[] arr5;

    return 0;
}