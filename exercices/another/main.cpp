#include <bits/stdc++.h>
using namespace std;
const uint m = 10e9 + 7;
uint calculating(uint a, uint b)
{
    if (b == 0)
    {
        return 1;
    }
    else
    {
        if (b % 2 == 0)
        {
            uint x = calculating(a, b >> 1);
            return (x * x) % m;
        }
        else
        {
            uint x = calculating(a, b - 1);
            return (a * x) % m;
        }
    }
}
int main()
{
    uint a, b;
    cin >> a >> b;
    cout << calculating(a, b);
}