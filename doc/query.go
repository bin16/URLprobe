package doc

type NamedNodeMap map[string]string

type DOMTokenList []string

type ClassList DOMTokenList

func (cl ClassList) Contains(className string) bool {
	for _, c := range cl {
		if c == className {
			return true
		}
	}

	return false
}

func (cl ClassList) Add(classNames ...string) {
	for _, c := range cl {
		if !cl.Contains(c) {
			cl = append(cl, c)
			return
		}
	}
}

func (cl ClassList) Remove(className string) {
	for i, c := range cl {
		if c == className {
			cl = append(cl[:i+1], cl[i+2:]...)
			return
		}
	}
}

func (cl ClassList) Replace(className, value string) {
	for i, c := range cl {
		if c == className {
			cl[i] = value
			return
		}
	}
}

func (cl ClassList) Toggle(className string) {
	if cl.Contains(className) {
		cl.Remove(className)
	} else {
		cl.Add(className)
	}
}
