// URL : https://leetcode.com/problems/two-sum/
// 4 ms 6.8 MB
#include<stdio.h>
#include<stdlib.h>

typedef struct hashmap *HashMap;
typedef struct entry *Entry;
typedef void (*Push)(HashMap, int, int);
typedef Entry (*Get)(HashMap, int);
typedef int (*HashCode)(HashMap, int);

struct entry{
    int key;
    int val;
    Entry next;
};

struct hashmap{
    int cap;
    Entry* list;
    Push push;
    Get get;
    HashCode hashcode;
};

void funcPush(HashMap , int, int);
Entry funcGet(HashMap , int);
int funcHashCode(HashMap , int);

HashMap __HashMap__(int initCap){
    HashMap new = malloc(sizeof(*new));
    new->cap = initCap;
    new->list = calloc(initCap, sizeof(Entry));
    new->push = funcPush;
    new->get = funcGet;
    new->hashcode = funcHashCode;
    return new;
}

void funcPush(HashMap self, int k, int v){
    int index = self->hashcode(self, k);
    Entry e = malloc(sizeof(*e));
    e->key  = k;
    e->val  = v;
    e->next = self->list[index];
    self->list[index] = e;
}

Entry funcGet(HashMap self, int k){
    int index = self->hashcode(self, k);
    Entry current;
    for(current = self->list[index]; current; current = current->next){
        if(current->key == k){
            return current;
        }
    }
    return NULL;
}

int funcHashCode(HashMap self, int k){
    return (k & 0x7fffffff) % self->cap;
}

int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    HashMap map = __HashMap__(numsSize);
    Entry tmp = NULL;
    int i;

    for(i=0; i<numsSize; i++){
        if(map->get(map, target - nums[i])){
            int* ret = malloc(2*sizeof(int));
            ret[0] = map->get(map, target - nums[i])->val;
            ret[1] = i;
            *returnSize = 2;
            return ret;
        }
        map->push(map, nums[i], i);
    }
    *returnSize = 0;
    return NULL;
}

int main(){
    int testcase[] = {5, 75, 25};
    int returnSize = 0;

    int* answer = twoSum(testcase, 3, 100, &returnSize);

    int i;
    for(i=0; i<returnSize; i++){
        printf("%d ",answer[i]);
    }

    printf("\n");
    return 0;
}
