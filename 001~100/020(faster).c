// URL : https://leetcode.com/problems/valid-parentheses
//0 ms	7.1 MB (拆開所有 fucntion)
#include<stdio.h>
#include<stdlib.h>
#include<string.h>

typedef enum{
    false,true
}bool;

//----------------------------------- Answer Begin -----------------------------------//
struct s{
    int token;
    struct s* next;
};
typedef struct s stack;

bool isValid(char* s){
    int i=0, getTokenInt;
    stack* rec=NULL;

    stack* new=malloc(sizeof(stack));
    new->next=rec;
    new->token=0;
    rec=new;

    while(s[i]!='\0'){
        switch (s[i++]) {
            case '(':
                getTokenInt= 1;
                break;
            case ')':
                getTokenInt = -1;
                break;
            case '[':
                getTokenInt = 2;
                break;
            case ']':
                getTokenInt = -2;
                break;
            case '{':
                getTokenInt = 3;
                break;
            case '}':
                getTokenInt = -3;
                break;
        }
        if(rec->token + getTokenInt == 0){
            stack* garbage=rec;
            rec=rec->next;
            free(garbage);
        }else if((rec->token) * getTokenInt < 0){
            return false;
        }else{
            new=malloc(sizeof(stack));
            new->next=rec;
            new->token=getTokenInt;
            rec=new;
        }
    }

    if(rec->token==0){
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
    system("pause");
    return 0;
}
