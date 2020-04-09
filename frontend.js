const fs = require("fs");
const stdin = fs.readFileSync(0).toString();

const tokenizerOutput = JSON.parse(stdin);

function createNode(token) {
  let cleanToken = token.value.replace(/^\s+|\s+$/g, '');
  return {
    name: cleanToken.split('(')[0],
    content: '',
    props: [],
    children: [],
  }
}

function createProp(token) {
  let cleanToken = token.value.replace(/^\s+|\s+$/g, '');

  return {
    key: cleanToken.split(':')[0],
    value: ''
  }
}

function buildTree(tokens) {
  let currentNode;
  let lastToken;
  let lastValidToken;
  
  for (let i = 0; i < tokens.length; i++) {
    const token = tokens[i];
    
    switch (token.type) {
      case 'nodeOpener':
        if (!currentNode) {
          currentNode = createNode(token)
        } else {
          let childNode = buildTree(tokens.slice(i))
          i = i + childNode[1]
          currentNode.children.push(childNode[0])
        }
        break;

      case 'string':
        if (lastValidToken.type == 'nodeOpener') {
          currentNode.content = token.value.replace(/'/g, '');
        } else if (lastValidToken.type == 'prop') {
          let lastPropIdx = currentNode.props.length - 1;
          currentNode.props[lastPropIdx].value = token.value.replace(/'/g, '');
        }
        break;
      
      case 'nodeCloser':
        return [currentNode, i];
      
      case 'prop':
        currentNode.props.push(createProp(token))
      default:
        break;
    }

    lastToken = token;
    if (lastToken.type != 'newLine') {
      lastValidToken = lastToken
    }
  }
};

let resultTree = buildTree(tokenizerOutput.tokens)
console.log(JSON.stringify(resultTree[0]));