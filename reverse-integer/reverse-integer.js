// URL : https://leetcode.com/problems/reverse-integer/
// 116 ms 37.4 MB
var reverse = function(x) {
    let y = 0;
    while(x!=0){
        y = y * 10 + x % 10;
        x = (x - x % 10) / 10;
    }
    return (y > 2147483647 || y < -2147483648)? 0 : y;
};

console.log(reverse(1234))
console.log(reverse(12345))
console.log(reverse(-12345))
console.log(reverse(2147483647))
console.log(reverse(1534236469))
console.log(reverse(-2147483648))
