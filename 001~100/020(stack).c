// URL : https://leetcode.com/problems/valid-parentheses
//4 ms	7 MB
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
    return;
}

void push(stack** s, int val){
    stack* new=malloc(sizeof(stack));
    new->next=*s;
    new->token=val;
    (*s)=new;
}

int pop(stack** s){
    if(!(*s)){
        fprintf(stderr, "Stack is empty\n");
        exit(1);
    }
    int token=(*s)->token;
    stack* garbage=(*s);
    (*s)=(*s)->next;
    free(garbage);
    return token;
}

bool isValid(char* s){
    int i=0, getTokenInt;
    stack* rec=NULL;
    push(&rec,0);  //0 is the end of rec
    while(s[i]!='\0'){
        getTokenInt=getToken(s[i++]);
        if(rec->token + getTokenInt == 0){
            pop(&rec);
        }else if((rec->token) * getTokenInt < 0){
            return false;
        }else{
            push(&rec,getTokenInt);
        }
    }

    if(pop(&rec)==0){
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
