// https://www.hackerrank.com/challenges/dynamic-array/problem?isFullScreen=true

#include <iostream>
#include <vector> 

using namespace std;

int main() {
    long long int n, q;
    cin >> n >> q;
    vector<long long int> arr[n];
    long long int lostAnswer = 0;

    for (long long int i = 0 ; i < q; i++) {
        long long int qType, x, y, idx;
        cin >> qType >> x >> y;
        
        idx = (x ^ lostAnswer) %n;

        switch (qType) {
            case 1:
                arr[idx].push_back(y);
                break;
            case 2:
                lostAnswer = arr[idx].at(y % arr[idx].size());
                cout<<lostAnswer<<endl;
            break;
        }
    }
}