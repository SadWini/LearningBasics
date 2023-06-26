#include <stdio.h>
using namespace std;

int main() {
    char c;
    bool begin = true;
    while(scanf("%c", &c), c != EOF) {
        if (c == '"') {
            if (begin) {
                printf("``");
            } else {
                printf("''");
            }
            begin = !begin;
        } else {
            printf("%c", c);
        }
    }
}