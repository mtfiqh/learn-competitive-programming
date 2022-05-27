// https://www.hackerrank.com/challenges/2d-array/problem?isFullScreen=true&h_r=next-challenge&h_v=zen
#include <iostream>
using namespace std;

int main() {
    int arr[6][6];

    for (int i=0; i<6; i++) {
        for (int j=0; j<6; j++){
            cin>>arr[i][j];
        }
    }
    bool first = true;
    int max;

    for (int i=0; i<=6-3; i++) {
        for(int j=0; j<=6-3; j++) {
            int a = arr[i][j];
            int b = arr[i][j+1];
            int c = arr[i][j+2];
            int d = arr[i+1][j+1];
            int e = arr[i+2][j];
            int f = arr[i+2][j+1];
            int g = arr[i+2][j+2];
            int sum = a+b+c+d+e+f+g;
            if (first) {
                first = false;
                max = sum;
            }else {
                if (max<sum) {
                    max = sum;
                }
            }

        }
    }

    cout << max;

}