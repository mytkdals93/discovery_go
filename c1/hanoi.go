package c1

import "strconv"

var result = []string{}

func move(disks, from, to, via int) {
	if disks == 0 {
		return
	}
	move(disks-1, from, via, to)
	result = append(result, ""+strconv.Itoa(from)+"->"+strconv.Itoa(to))
	move(disks-1, via, to, from)
	return

}

func Hanoi(disks int) []string {
	move(disks, 1, 2, 3)
	return result
}
