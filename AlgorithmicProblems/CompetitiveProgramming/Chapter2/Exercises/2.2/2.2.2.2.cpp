#include <stdio.h>

int main() {
    int a;
    scanf("%d", &a);
    printf("%s", a == a & -a ? "yes" : "no");
}