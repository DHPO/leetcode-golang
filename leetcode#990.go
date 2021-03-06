package main

type Union []byte

func (u *Union) Init() {
	*u = make([]byte, 26)
	for i := 0; i < 26; i++ {
		(*u)[i] = 'a' + byte(i)
	}
}

func (u *Union) GetRoot(b byte) (byte, int) {
	b -= 'a'
	origin := b
	rank := 0
	for (*u)[b] != b {
		b = (*u)[b]
		rank += 1
	}
	(*u)[origin] = b
	return b, rank
}

func (u *Union) Connect(a, b byte) {
	a -= 'a'
	b -= 'a'
	root1, rank1 := u.GetRoot(a)
	root2, rank2 := u.GetRoot(b)
	if root1 == root2 {
		return
	}
	if rank1 < rank2 {
		(*u)[root2] = root1
	} else {
		(*u)[root1] = root2
	}
}

func equationsPossible(equations []string) bool {
	u := Union{}
	u.Init()
	for _, e := range equations {
		if e[1] == '!' {
			continue
		}
		u.Connect(e[0], e[3])
	}
	for _, e := range equations {
		if e[1] == '!' {
			a, b := e[0], e[3]
			root1, _ := u.GetRoot(a)
			root2, _ := u.GetRoot(b)
			if root1 == root2 {
				return false
			}
		}
	}
	return true
}
