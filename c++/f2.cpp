#include<iostream>
#include<vector>
#include<iterator>
#include <sstream>
#include<ostream>
using namespace std;

bool isPal(string s){
    int mid = s.size()/2;
    int same =0;
    for (int i =0;i <mid;i++){
        if (s[i]==s[s.size()-1-i])
        same++;
    }
    return same==mid;
}
void solve( vector<int> arr, int n, int k )
{
    int mid = arr.size()/2;
    int chances = k;
   
    // cout<<k;
    for( int i= 0;i< mid;i++)
    {
        if (chances !=0)
        {
            if (arr[i]!=arr[arr.size()-i-1])
            {
                if (chances>=2){
                    // if (arr[i]<9){
                    arr[i]=9;
                    chances--;
                    // }
                    // if (arr[arr.size()-i-1]<9){
                    
                    arr[arr.size()-i-1]=9;
                    chances--;
                    // }
                }else{
                    arr[i]<arr[arr.size()-1-i]?
                
                    arr[i]=arr[arr.size()-1-i]
                    :
                    arr[arr.size()-i-1]= arr[i];
                    chances--;
                }
                    
            }
            else if (arr[i]== arr[arr.size()-i-1])
            {
                if (chances>=2){
                   if (arr[i]<9){
                    arr[i]=9;
                    chances--;
                    }
                    if (arr[arr.size()-i-1]<9){
                    
                    arr[arr.size()-i-1]=9;
                    chances--;
                    };
                }
               
            }
        }else{
            break;
        }
    }
    if (chances>0){
if (n%2==1&&(arr.size()>=5|| arr.size()==1))
    {
        arr[mid]=9;
        chances--;
    }
    }
     
    
    ostringstream oss;

  if (!arr.empty())
  {
    // Convert all but the last element to avoid a trailing ","
    copy(arr.begin(), arr.end()-1,
        ostream_iterator<int>(oss, ""));

    // Now add the last element with no delimiter
    oss << arr.back();
  }
if (!isPal(oss.str())){
    cout<<"-1"<<"\n";
    return;
}
  cout << oss.str() << std::endl;
    // string s(arr.begin(),arr.end());
    // ostream_iterator<int> os(cout);
    // // os=ostream_iterator<int>(cout); 
    // // copy(arr.begin(), arr.end(), os);
    // cout<<s<<"\n";
}



int main(){
    int n,k;
    string s;
    cin>>n;
    cin>>k;
    cin>>s;
    vector<int> arr;
    for (int i= 0;i<s.length();i++)
    {
        arr.push_back(int(s[i])-'0');
    }
    // for (auto i= arr.begin();i!= arr.end();i++)
    // {
    //     cout<<*i<<"\n";
    // }
    solve(arr,n,k);
}