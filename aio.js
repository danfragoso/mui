const fs = require('fs');
const muiString = fs.readFileSync('tests/01.mui').toString();

function parseMui(muiString, pointer) {
  muiNode = {
    name: '',
    content: '',
    props: [],
    children: []
  }

  stringIsOpen = false
  nodeIsOpen = false
  propIsOpen = false

  currentProp = ''
  currentPropValue = ''

  for(undefined; pointer < muiString.length; pointer++) {
    currentChar = muiString[pointer]
    specialChar = isReservedChar(currentChar)
    
    if (!specialChar) {
      if (stringIsOpen) {
        if (!propIsOpen) {
          muiNode.content += currentChar
        } else {
          currentPropValue += currentChar
        }
      } else {
        if (nodeIsOpen) {
          currentProp += currentChar
        } else {
          muiNode.name += currentChar
        }
      }
    } else if (specialChar == '\'') {
      stringIsOpen = stringIsOpen ? false : true;
    } else if (specialChar == '(') {
      if (nodeIsOpen) {
        muiNode.children.append(parseMui(muiString, pointer))
      } else {
        nodeIsOpen = true
      }
    } else if (specialChar == ')') {
      return muiNode
    } else if (specialChar == ':') {
      propIsOpen = true
    }
  }

  return muiNode
}

function isReservedChar(char) {
  if (char == '\'' || char == '(' || char == ')' || char == ':' || char == ',' || char == ' ' || char == '\n') {
    return char
  }

  return false
}

console.log(parseMui(muiString, 0))
