#include<iostream>
#include<ctime>
#include<cstdlib>
using namespace std;

template<class T>
void merge(T*arr,T*temp,const int l, const int m ,const int n)
{
    int i , iResult, j;
    for(i = l,iResult = l,j = m+1; (i <= m)&&(j<=n);iResult++)
    {
        if(arr[i]<=arr[j])
        {
            temp[iResult] = arr[i];
            i++;
        }
        else
        {
            temp[iResult] = arr[j];
            j++;
        }
        
    }

    copy(arr + i,arr+m+1,temp+iResult);
    copy(arr + j,arr+n+1,temp+iResult);

}

template<class T>
void mergepass(T*arr, T*temp,const int n,const int s)
{   
    int i;
    for( i = 0;i<=(n-2*s+1);i+=2*s)
    {
        merge(arr,temp,i,i+s-1,i+2*s-1);
    }

    if((i+s-1)<n)// only in n<2*s
    {
        merge(arr,temp,i,i+s-1,n);
    }   
    else
    {
        copy(arr+i,arr+n+1,temp+i);
    }
    
}

template<class T>
void mergesort(T*arr,const int n)
{
    T*temp = new int[n+1];
    for(int i = 1;i<n;i*=2)//size of a record
    {
        mergepass(arr,temp,n,i);
        i*=2;
        mergepass(temp,arr,n,i);
    }
    delete[] temp;
}

int main()
{
    srand(time(NULL));
    int m;
    int *arr = new int[100];
    int *temp = new int [100];
    for(int i = 0;i<100;i++)
    {
        if(i<=6)
        {
            m = i;
            arr[i] = rand()%100+1;
        }
        else
        {

           arr[i] = rand()%100+1; 
        }
       // temp[i] = NULL;
        
    }
   // merge(arr,temp,0,6,9);
    for(int i = 0;i<100;i++)
    {
         cout<< arr[i] << " ";
    }
    cout << endl;
    //mergepass(arr,temp,9,1);
    mergesort(arr,99);
    cout << endl;
    for(int i = 0;i<100;i++)
    {
        cout<< arr[i] << " ";
        
    }
    return 0;
}