package continuity

func major(dev uint) int {
	return int(uint(dev>>24) & 0xff)
}

func minor(dev uint) int {
	return int(dev & 0xffffff)
}

func makedev(major int, minor int) int {
	return ((major << 24) | minor)
}
