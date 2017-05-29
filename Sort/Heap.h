//
// Created by joe on 17-5-28.
//

#ifndef HEAP_HEAP_H
#define HEAP_HEAP_H
#include<cassert>
#include<algorithm>
using namespace std;
template<typename Item>
class MaxHeap{
private:
    Item *data;
    int count;
    int capacity;
    void shiftUp(int k){
        Item ret=data[k];
        while(k>1&&data[k/2]<data[k]){
            data[k/2]=data[k];
            k/=2;
        }
        data[k]=ret;
    }
    void shiftDown(int k){
        Item ret=data[k];
        while(2*k<=count) {
            int j = 2 * k;
            if(j+1<=count&&data[j+1]>data[j])
                j++;
            if(data[k]>data[j])break;
            swap(data[k],data[j]);
            k=j;
        }
    }
public:
    MaxHeap(int capacity){
        data=new Item[capacity+1];
        count=0;
        this->capacity=capacity;
    }
    MaxHeap(Item arr[],int n){
        data=new Item[n+1];
        capacity=n;
        for(int i=0;i<n;i++)
            data[i+1]=arr[i];
        count=n;
        for(int i=count/2;i>=1;i--)
            shiftDown(i);
    }
    ~MaxHeap(){
        delete[] data;
    }
    int size()
    {
        return count;
    }
    bool isEmpty(){
        return count==0;
    }
    void insert(Item item){
        assert(count+1<=capacity);
        data[count+1]=item;
        count++;
        shiftUp(count);
    }
    Item extractMax(){
        assert(count>0);
        Item ret=data[1];
        swap(data[1],data[count]);
        count--;
        shiftDown(1);
        return ret;
    }
};
#endif //HEAP_HEAP_H
