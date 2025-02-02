// 2.2.2.3 I don't see reasons to change anything. But we can consider using binary search.
#include <stdio.h>
#include <unordered_map>
using namespace std;

int main() {
    int v, n, a;
    scanf("%d%d", &v, &n);
    unordered_map<int, int> data;
    for (int i = 0; i < n; i++) {
        scanf("%d", &a);
        if (data.find(a) != data.end()) {
            data.insert({a, data.at(a) + 1});
        } else {
            data.insert({a, 1});
        }
    }
    for (auto i = data.begin(); i != data.end(); i++) {
        if (data.find(v - (i -> first)) != data.end()) {
            printf("Found %d + %d = %d", i -> first, v - (i -> first), v);
            break;
        }
    }
}