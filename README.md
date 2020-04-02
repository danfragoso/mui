# MUI
This is the MUI markup language, its main purpose is to describe user interfaces.

```
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
```

The most important mui lang element is the node element. Nodes can have three elements inside them: Content, Properties and Nodes. 

They are optional, but they are positional, the order in which they are represented matter.

### Content

Content is just the value of the node it's represented by a string characters enclosed by single quotes. A node can only have one content element.
```
'Hello World'
```

### Properties

Properties are key value pairs that are used to give a node any property, they are represented by the key, a colon and a value enclosed by single quotes.  A node can have multiple properties, separated by comma or new line.

```
key: 'value'
```

```
keyOne: 'valueOne', keyTwo: 'valueTwo'
```

```
keyOne: 'valueOne'
keyTwo: 'valueTwo'
```

### Nodes

Nodes are the most important element, they are represented by the node type and parenthesis.

```
node ()
```

A node with content, properties and nodes would look like this.

```
button (
	ref: 'saveButton'
	background: '#ccc'
	color: '#000'
	
	icon (srcPath: 'icons/floppy.svg')
	label ('SAVE')
)
```