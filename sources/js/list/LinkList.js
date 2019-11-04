/**
 * 单链表
 */

class Node {
    constructor(){
        this.data = null;
        this.next = null;
    }
}

class LinkList {
    constructor(){
        this.len = 0;
        this.first = null;
    }
    length(){
        return this.len;
    }
    display(){
        if (this.len == 0) {
            console.log("数据结构内无元素");
            return;
        }
        let currentNode = this.first;
        for (let i = 1; i <= this.len; i++) {
            console.log("位置 " + i + "元素为 " + currentNode.data);
            if (i == this.len) {
                break
            }
            currentNode = currentNode.next;
        }
    }
    append(e){
        let insertNode = new Node();
        insertNode.data = e;
        // 第一次追加
        if (this.len == 0) {
            this.first = insertNode;
            this.len++;
            return;
        }

        let currentNode = this.first;
        for (let i = 1; i <= this.len; i++ ) {
            if (i == this.len) {
                currentNode.next = insertNode;
            }
        }
        this.len++;
    }
    insert(index, e){
        if (index < 1 || index > this.len) {
            console.log("位序不正确");
            return
        }
        // 构造要插入的节点
        let insertNode = new Node();
        insertNode.data = e;

        // 插入：找到插入位置的前一个节点
        let currentNode = this.first;// 此处等于已经循环了一次，所以是 - 2
        for (let i = 1; i <= index - 2; i++) {
            currentNode = currentNode.next;
        }
        insertNode.next = currentNode.next;
        currentNode.next = insertNode;
        this.len++;
    }
    delete(index){      // 按照位序删除
        if (index < 1 || index > this.len + 1) {
            console.log("位序不正确");
            return
        }
        // 找到要删除元素的前一个元素
        let currentNode = this.first;// 此处等于已经循环了一次，所以是 - 2
        for (let i = 1; i <= index - 2; i++) {
            currentNode = currentNode.next;
        }
        // 如果要删除的是最后一个元素
        if (index == this.len) {
            currentNode.next = null;
            this.len--
            return
        }
        currentNode.next = currentNode.next.next;
        this.len--;
    }
    update(index, e){
        if(index < 1 || index > this.len) {
            console.log("位序不正确");
            return
        }
        let currentNode = this.first;
        for (let i = 1; i <= index; i++) {
            if (i == index) {
                currentNode.data = e
                break;
            }
            currentNode = currentNode.next;
        }
    }
    getElem(index){          // 按照位序查询
        if (index < 1 || index > this.len) {
            console.len("位序不合法");
            return
        }
        let currentNode = this.first
        for (let i = 1; i <= index; i++) {
            if (i == index) {
                return currentNode.data;
            }
            currentNode = currentNode.next;
        }
    }
    locateElem(e){      // 按照值查询位序
        let currentNode = this.first;
        for (let i = 1; i <= this.len; i++) {
            if (currentNode.data == e) {
                return i;
            }
        }
        console.log("未找到元素");
    }
    priorElem(e) {      // 查找前驱
        if (this.len <= 1) {
            console.log("数据结构元素不足，无法查询");
            return
        }
        if (this.first.data == e) {
            console.log("首元素无前驱");
            return
        }
        let currentNode = this.first;
        for (let i = 1; i<= this.len; i++) {
            if (currentNode.next.data == e) {
                return currentNode.data;
            }
            currentNode = currentNode.next;
        }
        console.log("未找到元素");
    }
    nextElem(e){             // 查询后继
        if (this.len <= 1) {
            console.log("数据元素不足，无法查询");
            return
        }
        let currentNode = this.first;
        for (let i = 1; i <= this.len; i++ )  {
            if (i == this.len) {
                console.log("最后一个元素无后继");
                return
            }
            if (currentNode.data == e) {
                return currentNode.next.data;
            }
            currentNode = currentNode.next;
        }
        console.log("未找到传入元素");
    }
    clear(){
        this.first = null;
        this.len = 0;
    }
}

export default LinkList