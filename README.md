# mui

This is mui its a markup language, its mains purpose is to describe user interfaces.

The center piece of the mui lang is a node, nodes are described this way:

node ()

Nodes can have tree kinds of things inside their parentesis, content, properties and nodes, they are not required for a node to be valid but they are positional

content

content is the value of the node, ex in the code bellow the content of the button node could be its label

button(
  'Click Me!'

  color: '#666'
  background: '#ff00ff'
)

properties

properties are key value pairs they are defined like this

key: 'value'

So, a mui node with content, props and nodes would look like

screen (
  'mainScreen'

  navBar (
    height: '50px'
    background: '#000'

    button (
      ref: 'backButton'
      iconPath: 'icons/arrowLeft.svg'
      iconFill: '#fff'
    )

    button (
      ref: 'homeButton'
      iconPath: 'icons/menu.svg'
      iconFill: '#fff'
    )
  )
)