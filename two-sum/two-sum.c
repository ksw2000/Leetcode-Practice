// https://leetcode.com/problems/two-sum/
// 8 ms 6.9 MB
#include <stdio.h>
#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
struct hashEl {
    int key;
    int val;
    struct hashEl* next;
};

struct hash {
    struct hashEl** list;
    int cap;
};

struct hash* Hash(int cap) {
    struct hash* h = malloc(sizeof *h);
    h->list = calloc(cap, sizeof(struct hashEl*));
    h->cap = cap;
    return h;
}

unsigned int hashcode(struct hash* h, int key) {
    return key & (h->cap - 1);
}

void set(struct hash* h, int key, int val) {
    int code = hashcode(h, key);
    struct hashEl* el = malloc(sizeof *el);
    el->key = key;
    el->val = val;
    el->next = h->list[code];
    h->list[code] = el;
}

int* get(struct hash* h, int key) {
    int code = hashcode(h, key);
    struct hashEl* current = h->list[code];
    for (; current; current = current->next) {
        if (current->key == key) {
            return &(current->val);
        }
    }
    return NULL;
}

int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    struct hash* h = Hash(256);
    int i;
    int* j;
    int* ans = malloc(sizeof(int) * 2);
    for (i = 0; i < numsSize; i++) {
        j = get(h, target - nums[i]);
        if (j) {
            ans[0] = *j;
            ans[1] = i;
            *returnSize = 2;
            return ans;
        }
        set(h, nums[i], i);
    }
    *returnSize = 0;
    return NULL;
}