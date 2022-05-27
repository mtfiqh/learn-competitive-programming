// https://www.hackerrank.com/challenges/array-left-rotation/problem?isFullScreen=true
#include <iostream>
using namespace std;

int main() {
    int n, d;
    cin>>n>>d;
    int arr[n];

    for(int i=0; i<n;i++) {
        cin>>arr[i];
    }

    int newArr[n];
    for(int i = 0; i<n; i++) {
        int c = d;
        while (c+i > n-1) {
            c-=n;
        }
        
        newArr[i] = arr[i+c];
        cout << newArr[i]<<" ";
    }
}