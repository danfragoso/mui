const fs = require("fs");
const stdin = fs.readFileSync(0).toString();

const muiTree = JSON.parse(stdin);
const tab = '  ';

function generateHTML(node, n) {
  let childrenHtml = '';
  let propString = '';
  
  node.children.forEach(child => {
    childrenHtml += generateHTML(child, n+1);
  });

  node.props.forEach(prop => {
    propString += `${prop.key}="${prop.value}" `;
  });

  let htmlString = '';
  htmlString = n < 1 ? '<!DOCTYPE html>' : '';
  htmlString += '\n' + tab.repeat(n) + `<${node.name.trim()}${propString ? ' ' + propString.trim() : ''}>`;

  if (node.content) {htmlString += '\n' + tab.repeat(n + 1) + node.content};
  
  htmlString += childrenHtml;
  htmlString += '\n' + tab.repeat(n) + `</${node.name.trim()}>`;
  
  return htmlString;
}

console.log(generateHTML(muiTree, 0));