#include <stdio.h>
using namespace std;

int main() {
    char c;
    while(scanf("%c", &c), c != EOF) {
        printf("%c", c);
    }
}