package main

import (
	"./customsmt"
	"log"
)


func main() {
	//Build list of Content to build tree
	var content []string
	content = append(content, "a")
	content = append(content, "b")
	//content = append(content, "c")
	//content = append(content, "d")

	sContent := customsmt.CreateContent(content)


	t := customsmt.CreateTree(sContent)

	log.Println(customsmt.Hashes(t))

	var content2 []string
	content2 = append(content2, "hv1ly")
	content2 = append(content2, "cl444ivv7l")
	content2 = append(content2, "cl4iv6vl")
	content2 = append(content2, "cl1iv4v4l")
	content2 = append(content2, "cl2ivv5l")
	content2 = append(content2, "4444")

	sContent2 := customsmt.CreateContent(content)

	customsmt.Test(sContent2, t)

	log.Println(customsmt.Hashes(t))
	//log.Println(customsmt.GetMerkleRoot(t))
	//log.Println(customsmt.VerifySpecificLeaf(t, sContent[1]))


	//////Create a new Merkle Tree from the list of Content
	//t, err := merkletree.NewTree(list)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//////Get the Merkle Root of the tree
	//mr := t.MerkleRoot()
	//log.Println(mr)
	//
	//////Verify the entire tree (hashes for each node) is valid
	//vt, err := t.VerifyTree()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Verify Tree: ", vt)
	//
	////Verify a specific content in in the tree
	//vc, err := t.VerifyContent(list[0])
	//if err != nil {
	//	log.Fatal(err)
	//}
	////
	//log.Println("Verify Content: ", vc)
	//
	////String representation
	//log.Println(t)
	//
	//st := t.String()
	//
	//log.Println(st)
}

