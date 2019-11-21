// URL : https://leetcode.com/problems/implement-strstr/
//52 ms	33.8 MB
var strStr = function(haystack, needle) {
    return haystack.indexOf(needle);
};

var haystack, needle;
haystack="hello"; needle="ll";
console.log(strStr(haystack, needle));
haystack="aaaaa"; needle="bba"
console.log(strStr(haystack, needle));
