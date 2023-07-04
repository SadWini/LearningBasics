#include <iostream>
#include <stdio.h>
#include <string>
using namespace std;

int main() {
    string s;
    int n, r, d, tempr = -1, tempd = -1, min = 100000000;
    while(scanf("%d", &n), n > 0) {
        cin >> s;
        min = 100000000;
        tempd = -1;
        tempr = -1;
        for (int i = 0; i < n; i++) {
            if (s[i] == 'Z') {
                min = 0;
            }
            if (s[i] == 'R') {
                tempr = i;
                if (tempd > -1 && min > tempr - tempd) {
                    min = tempr - tempd;
                }
            }
            if (s[i] == 'D') {
                tempd = i;
                if (tempr > -1 && min > tempd - tempr) {
                    min = tempd - tempr;
                }
            }
        }
        printf("%d\n", min);
    }
}