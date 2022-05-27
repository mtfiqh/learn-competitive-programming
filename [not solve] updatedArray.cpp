// https://www.hackerrank.com/challenges/crush/problem?isFullScreen=true&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

#include <iostream>
using namespace std;

int main() {
    unsigned long long int n, m;
    cin >> n >> m;
    unsigned long long int arr[n+1];
    for (long long int i = 0 ; i<n; i ++) {
        arr[i] = 0;
    }

    unsigned long long int max = 0;
    for (unsigned long long int i=0; i<m; i++) {
        unsigned long long int a,b,c;
        cin >> a >> b >> c;

        for (unsigned long long int j = a-1 ; j < b; j++) {
            arr[j]+=c;
            if (arr[j] > max) {
                max = arr[j];
            }
        }
    }
     for (long long int i = 0 ; i<n; i ++) {
       cout<<arr[i]<<" ";
    }
    cout<<endl;

    cout<<max<<endl;

}