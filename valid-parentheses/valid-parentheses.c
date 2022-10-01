#include<stdio.h>
#include<stdlib.h>

typedef enum{
    false,
    true
}bool;

struct s {
    int token;
    struct s* next;
};
typedef struct s stack;

bool isValid(char* s) {
    int i = 0, getTokenInt;
    stack* rec = NULL;

    // push
    stack* new = malloc(sizeof(stack));
    new->next = rec;
    new->token = 0;
    rec = new;

    while (s[i] != '\0') {
        switch (s[i++]) {
            case '(':
                getTokenInt = 1;
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
        if (rec->token + getTokenInt == 0) {
            stack* garbage = rec;
            rec = rec->next;
            free(garbage);
        } else if ((rec->token) * getTokenInt < 0) {
            return false;
        } else {
            new = malloc(sizeof(stack));
            new->next = rec;
            new->token = getTokenInt;
            rec = new;
        }
    }

    return rec->token == 0;
}

int main(){
    char* str1 = "()";
    char* str2 = "()[]{}";
    char* str3 = "(]";
    printf("isValid('%s') = %s\n", str1, isValid(str1) ? "true" : "false");
    printf("isValid('%s') = %s\n", str2, isValid(str2) ? "true" : "false");
    printf("isValid('%s') = %s\n", str3, isValid(str3) ? "true" : "false");
}