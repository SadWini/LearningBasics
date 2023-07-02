#include <stdio.h>
#include <string>
#include <iostream>
#include <map>
using namespace std;

int main() {
    map<string, int> data;
    int n, amount, x;
    string name;
    while (cin >> n) {
        data.clear();
        while(n--) {
            cin >> name;
            data.insert({name, 0});
        }
        for (auto it = data.begin(); it != data.end(); it++) {
            cin >> name >> amount >> x;
            while (x) {
                cin >> name;
                it -> second -= amount;
                data.find(name) -> second += amount;
                x--;
            }
        }
        for (auto it = data.begin(); it != data.end(); it++){
            cout << it -> first << " " << it -> second << "\n";
        }
        cout << "\n";
    }
}