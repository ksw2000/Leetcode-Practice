// URL : https://leetcode.com/problems/reverse-integer/
var reverse = function(x) {
    var answer=0;
    while(x!=0){
        tmp=x%10;
        x=(x-tmp)/10;
        answer=answer*10+tmp;
    }
    return (answer>2147483647 || answer < -2147483648)? 0 : answer;
};
