#include <stdio.h>
#include <cstring>

using namespace std;
int main() {
    char call[10];
    int n, balance, temp;
    balance = 0;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%s", &call);
        if (strcmp(call, "donate") == 0) {
            scanf("%d", &temp);
            balance += temp;
        }
        if (strcmp(call, "report")) {
            printf("%d\n", balance);
        }
    }
    return 0;
}