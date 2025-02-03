#include <stdio.h>
using namespace std;

int main() {
    long x, y;
    int n;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%ld%ld", &x, &y);
        printf("%c\n", x == y ? '=' : x > y ? '>' : '<');
    }
}