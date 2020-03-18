#include<iostream>
#include<ctime>
#include<cstdio>
using namespace std;
#define ARRAY_SIZE 100 

template <class T>
void swap(T*a,T*b)
{
    T*temp = a;
    a=b;
    b=temp;
}


template <class T>
void Adjust(T*arr,const int root,const int n)
{
    T r= arr[root];
    int j;
    for(j = root*2;j<=n;j*=2)
    {
        if((j<n)&&(arr[j] < arr[j+1]))
            j++;
        if(arr[j] <= r)
            break;
        arr[j/2] = arr[j];
    }
    arr[j/2] = r;
}

template <class T>
void heapsort(T*arr,const int n)
{
    for(int i = n/2;i>=0;i--)
    {
        Adjust(arr,i,n);
    }

    for(int i = n-1;i>=0;i--)
    {
        swap(arr[0],arr[i+1]);
        Adjust(arr,0,i);
    }

}

int main()
{
    srand(time(NULL));
    int *arr;
    arr = new int [ARRAY_SIZE];
    for(int i = 0;i<ARRAY_SIZE;i++)
    {
        arr[i] = rand()%1000+1;
    }
    heapsort(arr,99);
    for(int i = 0;i<ARRAY_SIZE;i++)
    {
        cout << arr[i] << " ";
    }
}