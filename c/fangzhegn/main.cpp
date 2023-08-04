// 本题为考试多行输入输出规范示例，无需提交，不计分。
#include <iostream>
#include <cstdio>

using namespace std;

int main(){
    //freopen("1.in","r",stdin);
    int n,ans = 0;
    cin >> n;
    for(int i = 0; i < n; i++){
        for(int j = 0; j < n; j++){
            int x; scanf("%d",&x);
            ans += x;
        }
    }
    cout << ans << endl;
    return 0;
}