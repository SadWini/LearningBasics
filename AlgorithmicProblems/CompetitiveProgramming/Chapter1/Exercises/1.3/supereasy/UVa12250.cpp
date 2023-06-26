#include <bits/stdc++.h>
using namespace std;

int main() {
    int k = 1;
    string str;
    while(cin >> str && str != "#") {
        if (str == "HELLO") puts("ENGLISH");
        else if (str == "HOLA") puts("SPANISH");
        else if (str == "HALLO")  puts("GERMAN");
        else if (str == "BONJOUR")  puts("FRENCH");
        else if (str == "CIAO") puts("ITALIAN");
        else if (str == "ZDRAVSTVUJTE") puts("RUSSIAN");
        else puts("UNKNOWN");
    }
}