#include <stdio.h>
#include <math.h>
using namespace std;

int main() {
    int k, n;
    scanf("%d", &k);
    for (int i = 0; i < k; i++) {
        scanf("%d", &n);
        printf("%d\n", abs(((n * 63 + 7492) * 5 -498) % 100 / 10));
    }
}