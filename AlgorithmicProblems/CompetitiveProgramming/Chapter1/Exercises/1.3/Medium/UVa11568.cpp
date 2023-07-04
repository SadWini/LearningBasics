#include <stdio.h>
#include <string>
#include <algorithm>
#include <iostream>
using namespace std;

int main() {
    int n, m, f;
    string s;
    scanf("%d", &n);
    while(n > 0) {
        n--;
        m = 0;
        f = 0;
        getline(cin, s);
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == 'M') {
                m++;
            }
            if (s[i] == 'F') {
                f++;
            }
        }
        printf("%sLOOP\n", m == f ? "" : "NO ");
    }
}