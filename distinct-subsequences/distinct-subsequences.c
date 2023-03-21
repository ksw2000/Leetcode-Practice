// https://leetcode.com/problems/distinct-subsequences/
// 4 ms 5.8 MB
struct el{
    int val;
    unsigned int counter;
};

struct list{
    struct el* l;
    int cap;
    int len;
};

struct list* List(int cap){
    struct list* l = malloc(sizeof * l);
    l->l = malloc(sizeof(struct el) * cap);
    l->cap = cap;
    l->len = 0;
    return l;
}

void append(struct list* l, int val, int counter){
    if(l->cap == l->len){
        struct el* new = realloc(l->l, sizeof(struct el) * l->cap << 1);
        l->l = new;
        l->cap = l->cap << 1;
    }
    l->l[l->len].val = val;
    l->l[l->len].counter = counter;
    l->len++;
}

int numDistinct(char * s, char * t){
    struct list* l1 = List(125);
    struct list* l2 = List(125);
    struct list* l3;
    int lenS = strlen(s);
    int lenT = strlen(t);
    int h, i, j;
    int k = 0;
    unsigned int counter = 0;
    append(l1, -1, 1);
    for(i = 0; t[i] != '\0'; i++){
        l2->len = 0;
        for(j = k; s[j] != '\0' && j < lenS-(lenT-i-1); j++){
            if(t[i] == s[j]){
                counter = 0;
                for(h = 0; h < l1->len && l1->l[h].val < j; h++){
                    counter += l1->l[h].counter;
                }
                if(counter) append(l2, j, counter);
            }
        }
        if(l2->len == 0){
            return 0;
        }
        k = l2->l[0].val+1;
        l3 = l1;
        l1 = l2;
        l2 = l3;
    }
    int res = 0;
    for(i = 0; i < l1->len; i++){
        res += l1->l[i].counter;
    }
    return res;
}