#include <stdio.h>
#include <string>
#include <iostream>

using namespace std;

int main() {
    int i = 1;
    string str;
    while (cin >> str && str != "*") {
        printf("Case %d: ", i++);
        if (str == "Hajj") {
            puts("Hajj-e-Akbar");
        }
        if (str == "Umrah") {
            puts("Hajj-e-Asghar");
        }
    }
    return 0;
}