package min_len_encoding

type trie struct {
	children []*trie
	count    int
}

func newTrie() *trie {
	return &trie{
		children: make([]*trie, 26),
		count:    0,
	}
}

func (t *trie) get(c byte) *trie {
	if t.children[c-'a'] == nil {
		t.children[c-'a'] = newTrie()
		t.count++
	}
	return t.children[c-'a']
}

func minimumLengthEncoding(words []string) int {
	t := newTrie()
	nodes := make(map[*trie]int)

	for i := 0; i < len(words); i++ {
		word := words[i]
		cur := t
		for j := len(word) - 1; j >= 0; j-- {
			cur = cur.get(word[j])
		}
		nodes[cur] = i
	}

	ans := 0

	for node := range nodes {
		if node.count == 0 {
			ans += len(words[nodes[node]]) + 1
		}
	}

	return ans
}
