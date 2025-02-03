#include <stdio.h>
#include <string>
#include <iostream>

using namespace std;
int main() {
    int n, rate[10], max;
    string data[10];
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        max = 0;
        for (int j = 0; j < 10; j++) {
            cin >> data[j] >> rate[j];
            if (rate[j] > max) max = rate[j];
        }
        printf("Case #%d:\n", i + 1);
        for (int j = 0; j < 10; j++) {
            if (rate[j] == max) {
                cout << data[j] << "\n";
            }
        }
    }
    return 0;
}