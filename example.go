package main

import (
  "log"

  "github.com/matzhouse/go-data-structures/list"
  "github.com/matzhouse/go-data-structures/set"
)

func main() {

  showList()
  showSet()

}

func showList() {

  log.Println("List Example:")

  l := list.NewList()

  log.Println("We add some values to the list")
  l.InsertAtBeginning("hello")
  l.InsertAtEnd("world")

  log.Println("does the list contain 'hello' ?")
  log.Println(l.Contains("hello"))

}

func showSet() {

  log.Println("Set Example:")

  s := set.NewSet()

  log.Println("Add a value..")
  err := s.Add("hello")

  if err != nil {
    log.Println(err)
    return
  }

  log.Println("And another..")
  err = s.Add("world!")

  if err != nil {
    log.Println(err)
    return
  }

  log.Println("Oops - we try the same one..")
  err = s.Add("world!")

  log.Println("We now find we can't add 'world!' again")
  log.Println(err)

}
