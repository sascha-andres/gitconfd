package main

import "livingit.de/code/gitconfd"

func main() {
	gitconfd.Execute("../", "commit-msg")
}
