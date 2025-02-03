#include <stdio.h>
#include <unordered_map>
using namespace std;

int main() {
    int n, a;
    scanf("%d", &n);
    unordered_map<int, int> data;
    for (int i = 0; i < n; i++) {
        scanf("%d", &a);
        if (data.find(a) != data.end()) {
            printf("Duplicate exist %d", a);
            break;
        } else {
            data.insert({a, 1});
        }
    }
}
