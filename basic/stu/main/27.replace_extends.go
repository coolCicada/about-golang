package main

import "fmt"

func main() {
	author1 := author{"l", "s", "go"}
	post1 := post{"go", "hh", author1}
	post2 := post{
		"go1",
		"hh1",
		author1,
	}
	post3 := post{"go2", "hh2", author1}
	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()
}

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println(p.title, p.content, p.fullName(), p.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("content of website")
	for _, v := range w.posts {
		v.details()
	}
}
