package main

import (
	"fmt"

	mui "./mui"
)

var source = `
html (
  head (
    style ('
      body {
        width: 100vw;
        height: 100vh;
        background: #ccc;
      }

      .text {
        text-align: center;
        font-size: 200px;
      }
    ')
  )
  body (
    span (
      'tipos de carinhas sao'
      class: 'text'
    )
  )
)
`

func main() {
	tree, err := mui.Parse(source)
	if err != nil {
		panic("memes")
	}

	fmt.Println(tree)
	node := mui.NewNode("memes")
	fmt.Println(node)
}
