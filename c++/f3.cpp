#include<iostream>
#include<map>

using namespace std;

void solve(string s)
{
    int del=0;
    map<int,bool> skipIndex;
    int size = s.size()-1;

    int i =0 ;
    int j=1;

    while(i<size && j<size+1)
    {
        if (skipIndex[j])
        {
            j++;
        }

        if (s[i]==s[j])
            {
                cout<<j<<"\n";
                skipIndex[j]=true;
                j++;
            }
        if (s[i]!=s[j])
        {
            i=j;
            j++;
        }
    }
    for (auto i = skipIndex.begin(); i != skipIndex.end(); i++)
    {
        cout<<i->first<<"\n";
        /* code */
    }
    
    cout<<skipIndex.size()<<"\n";

}

int main()
{
    solve("BABABA");
    return 0;
}