package main

import (
	. "linkedlist/list"
	"fmt"
)

func main() {
	var monsters = IdList();
	monsters.AddToList(100);
	monsters.AddToList(150);
	monsters.AddToList(111);
	monsters.AddToList(325);
	monsters.AddToList(666);
	monsters.AddToList(999);
	monsters.AddToList(1010101010);
	
	//monsters.RemoveById(100);
	monsters.RemoveById(150);
	//monsters.RemoveById(111);
	//monsters.RemoveById(325);
	//monsters.RemoveById(666);
	//monsters.RemoveById(999);
	monsters.RemoveById(1010101010);

	var current *IdNode_t = monsters.First;
	for !(current == nil) {
		fmt.Println(GetId_s(current.Parent), "<-", current.Id, "->", GetId_s(current.Child));
		current = current.Child;
	}

}