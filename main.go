package main

import (
	"log"

	"github.com/MasterEvarior/gize/cmd/git"
)

func main() {
	repositories, _ := git.GetAllRepositories("/home/giannin/Documents/Github")
	for _, repo := range repositories {
		log.Println(repo.Name)
	}
}
