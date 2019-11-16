// URL : https://leetcode.com/problems/add-two-numbers/
function ListNode(val) {
    this.val = val;
    this.next = null;
}

function printLink(link){
    var current=link;
    while(current!=null){
        console.log(current.val);
        current=current.next;
    }
}

//----------------------------------- Answer Begin -----------------------------------//
function push(link,val){
    var newLink = new ListNode(val);
    while(link.next!=null){
        link=link.next;
    }
    link.next=newLink;
}

var addTwoNumbers = function(l1, l2) {
    var carry=0;
    let v=l1.val + l2.val;
    var answer=new ListNode(v%10), currentAnswer=answer;
    carry=Math.floor(v/10);

    var current1=l1.next, current2=l2.next;
    while(current1!=null && current2!=null){
        let v=current1.val + current2.val + carry;
        push(currentAnswer, v%10);
        carry=Math.floor(v/10);
        currentAnswer=currentAnswer.next;  //加速 push
        current1=current1.next;
        current2=current2.next;
    };
    //以上步驟已將最小位數相加，比如 3 位數 + 5 位數 那麼已完成 3 位數

    while(current1 != null){
        let v=current1.val + carry;
        push(currentAnswer, v%10);
        carry=Math.floor(v/10);
        currentAnswer=currentAnswer.next;
        current1=current1.next;
    }

    while(current2 != null){
        let v=current2.val + carry;
        push(currentAnswer, v%10);
        carry=Math.floor(v/10);
        currentAnswer=currentAnswer.next;
        current2=current2.next;
    }

    while(carry!=0){    //避免進位超過l1和l2的長度
        let v=carry;
        push(currentAnswer, v%10);
        carry=Math.floor(v/10);
        currentAnswer=answer.next;  //加速 push
    }

    return answer;
};
//----------------------------------- Answer END -----------------------------------//

var l1 = new ListNode(7);
push(l1,8);
push(l1,9);

var l2 = new ListNode(7);
push(l2,8);
push(l2,9);
push(l2,9);
push(l2,9);

printLink(addTwoNumbers(l1,l2));
