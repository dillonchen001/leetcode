package main

import "fmt"

func main() {
	path := "/..hidden"
	fmt.Println(simplifyPath(path))
}

func simplifyPath(path string) string {
	lp := len(path)

	var ret string

	if lp == 0 {
		return ret
	}

	point := 0
	chars := 0
	for i := 0; i < lp; i++ {
		if path[i] == '.' {
			point++
		} else if path[i] == '/' {
			if chars > 0 || point > 2 {
				ret += "/"
				ret += path[i-chars-point : i]
				point = 0
				chars = 0
			} else if point < 2 {
				point = 0
				continue
			} else if point == 2 {
				if len(ret) > 0 {
					m := len(ret)
					for j := m - 1; j >= 0; j-- {
						if ret[j] == '/' {
							if j == 0 {
								ret = ""
							} else {
								ret = ret[0:j]
							}
							break
						}
					}
				}
				point = 0

			}
		} else {
			chars++
		}
	}
	fmt.Println(chars)

	if chars > 0 || point > 2 {
		ret += "/"
		ret += path[lp-chars-point : lp]
		point = 0
		chars = 0
	}

	if point == 2 {
		m := len(ret)
		for j := m - 1; j >= 0; j-- {
			if ret[j] == '/' {
				if j == 0 {
					ret = ""
				} else {
					ret = ret[0:j]
				}
				break
			}
		}
	}

	if len(ret) == 0 {
		return "/"
	}

	return ret
}
