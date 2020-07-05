package homework

type Trie struct {
	data string
    children [26]*Trie
    isEndingChar bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    for _, v := range word {
        index := v - 'a'
        if this.children[index] == nil {
            newNode := new(Trie)
            this.children[index] = newNode
        }
        this = this.children[index]
    }
    this.isEndingChar = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    for _, v := range word {
        index := v - 'a'
        if this.children[index] == nil {
            return false
        }
        this = this.children[index]
    }
    if this.isEndingChar == false{
        return false
    }
    return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    for _, v := range prefix {
        index := v - 'a'
        if this.children[index] == nil {
            return false
        }
        this = this.children[index]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */