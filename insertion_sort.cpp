#include<iostream>
using namespace std;

void insertionsort(int arr[],int n)
{
    int i ,j,key;
    for(int i = 1;i<n;i++)
    {
        key = arr[i];
        j=i-1;

        while(j>=0&& arr[j] > key)
        {
            arr[j+1]=arr[j];
            j--;
        }
        arr[j+1] = key;
    }
}

void print(int arr[],int n)
{
    for(int i = 0;i<n;i++)
    {
        cout << arr[i] << endl;
    }
}

int main()
{
    int arr[] = {12,5,3,1,16,4};
    int n = sizeof(arr)/sizeof(arr[0]);

    insertionsort(arr,n);
    print(arr,n);
    return 0;
}