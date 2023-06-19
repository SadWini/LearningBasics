#include <bits/stdc++.h>
using namespace std;

void convertStrtoArr(string str)
{
    int str_length = str.length();
    int arr[str_length] = { 0 };
    int j = 0, i, sum = 0;
    for (i = 0; i<str.length(); i++) {
        if (str[i] == ' '){
            j++;
        } else {
            arr[j] = arr[j] * 10 + (str[i] - 48);
        }
    }

    for (i = 0; i <= j; i++) {
        sum += arr[i];
    }
    cout<<endl;
    cout<<sum<<endl;
}

int main()
{
    string str = "100 34 25 6 3 14";
    convertStrtoArr(str);
    return 0;
}