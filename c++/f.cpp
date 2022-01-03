#include<iostream>
#include<map>
#include<math.h>
using namespace std;

void util (map<int,int> mp){
    for (auto i = mp.begin(); i != mp.end(); i++)
    {
        cout<<i->first<<"\t"<<i->second<<"\n";
        /* code */
    }
    
}
void solve( string s){
    map<char,int> occ;
    map<int,int> occ2;
    int c=0;
    for(int i=0;i<s.length();i++)
    {
        char x = s[i];
        occ[x]++;
        // occ2[occ[x]]++;
 
    }
    int minv =INT8_MAX;
    for (auto i = occ.begin(); i != occ.end(); i++)
    {
        occ2[i->second]++;
        minv = min(minv,i->second);
    }
    
    // cout<<"min\t"<<minv<<"\n";
    // util(occ2);
    if (occ2.size()==1){
        cout<<"YES"<<"\n";
        return ;
    }
    if (occ2.size()>2){
        cout<<"NO"<<"\n";
        return ;
    }
    string ans="NO\n";
    for (auto itr= occ2.begin(); itr!= occ2.end();itr++)
    {
        if (itr->second==1 &&( itr->first-1<=1||itr->first-1==minv)){
            ans ="YES\n";
            break;
        }
  
    }

    cout<<ans;
}

int main(){
    string s;
    cin>>s;
   solve(s);
    
    return 0;
}
