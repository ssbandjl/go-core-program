#include <iostream>
#include <sstream>
#include <vector>
 
using namespace std;

// g++ main.cpp -o main;./main
int main() {
    int start, end, n = 0;
    string str, word, nstr;
    getline(cin, str);
    stringstream ss(str);
    cin>>start>>end;
    cout << end <<endl;
    if (start < 0){
        cout << start <<endl;
        start=0;
    }
    if (end > 3){
        end=3;
    }
        cout << end <<endl;
    // cout << start<<endl;
    vector<string> arr;
    while (ss>>word) {
        arr.push_back(word);
    }
    if (end > arr.size()){
      cout<<end<<endl;
        end=arr.size();
    }
    if (start>=arr.size() || start>end || end>arr.size()) {
        cout<<str<<endl;
        return 0;
    }

    for (int i=0; i<start; ++i) {
        if (nstr.size() == 0) {
            nstr = arr[i];
        } else {
            nstr = nstr + " " + arr[i];
        }
    }
    for (int i=end; i>=start; --i) {
        if (nstr.size() == 0) {
            nstr = arr[i];
        } else {
            nstr = nstr + " " + arr[i];
        }
    }
    for (int i=end+1; i<arr.size(); ++i) {
        if (nstr.size() == 0) {
            nstr = arr[i];
        } else {
            nstr = nstr + " " + arr[i];
        }
    }
    cout<<nstr<<endl;
 
 
    return 0;
}