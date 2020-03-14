#include<iostream>
#include<ctime>
#include<cstdlib>
using namespace std;
const int n = 10;

template <class T1>
void swap(T1*a,T1*b)
{
    T1*temp = a;
    a = b;
    b = temp;
}

template <class T>
void quicksort(T *arr,int left,int right)
{

    if(left < right)
    {
        int i = left;
        int j=right+1;
        int pviot = arr[left];

        do
        {
            do
            {
                i++;
            } while (arr[i]<pviot);

            do
            {
                j--;
            } while (arr[j]>pviot);

            if(i<j)
            {
                swap(arr[i],arr[j]);
            }

        } while (i<j);

        swap(arr[left],arr[j]);

        quicksort(arr,left,j-1);
        quicksort(arr,j+1,right);

    }
}


int main()
{
    srand(time(NULL));
    int *arr = new int[n];
    for(int i = 0;i<n;i++)
    {
        arr[i] = rand()%1000+1;
        cout<< arr[i] << endl;
    }
    quicksort(arr,0,n-1);
    cout << endl;
    for(int i = 0;i<n;i++)
    {
        cout<< arr[i] << endl;
    }
    return 0;
}