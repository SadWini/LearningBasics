#include <iostream>
#include <string>
using namespace std;

int main() {
    string data;
    long n, k;
    while (cin >> data, data != "END") {
        n = data.size();
        if (n > 1) k = 2;
        if (n == 1) k = 1;
        while(data.size() > 1) {
            k = k + 1;
            data = to_string(n);
            n = data.size();
        }
        printf("%ld\n", k);
    }
}