// URL : https://leetcode.com/problems/valid-parentheses
//4 ms	7 MB
#include<stdio.h>
#include<stdlib.h>
#include<string.h>

typedef enum{
    false,true
}bool;

//----------------------------------- Answer Begin -----------------------------------//
int getToken(char c){
    switch (c) {
        case '(':
            return 1;
        case ')':
            return -1;
        case '[':
            return 2;
        case ']':
            return -2;
        case '{':
            return 3;
        case '}':
            return -3;
    }
    return 0;
}

bool isValid(char* s){
    int i=0,getTokenInt;
    int top=-1;
    int* stack=malloc(sizeof(int)*(strlen(s)+1));
    stack[++top]=0; //0 is the end of this stack
    while(s[i]!='\0'){
        getTokenInt=getToken(s[i++]);
        if(getTokenInt + stack[top] == 0){
            top--;
        }else if(getTokenInt * stack[top] <0){
            return false;
        }else{
            stack[++top]=getTokenInt;
        }
    }
    if(stack[top]==0){
        return true;
    }else{
        return false;
    }
}
//----------------------------------- Answer END -----------------------------------//

int main(){
    char s[16];
    strcpy(s,"()");
    printf("%s \t\t %d\n",s,isValid(s));
    strcpy(s,"()[]{}");
    printf("%s \t\t %d\n",s,isValid(s));
    strcpy(s,"(]");
    printf("%s \t\t %d\n",s,isValid(s));
    strcpy(s,"{[]}");
    printf("%s \t\t %d\n",s,isValid(s));
    strcpy(s,"()[]{}");
    printf("%s \t\t %d\n",s,isValid(s));
    system("pause");
    return 0;
}
