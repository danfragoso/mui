const fs = require("fs");
const stdin = fs.readFileSync(0).toString();

const muiTree = JSON.parse(stdin);

function generateHTML(node) {
  let childrenHtml = '';
  let propString = '';
  
  node.children.forEach(child => {
    childrenHtml += generateHTML(child);
  });

  node.props.forEach(prop => {
    propString += `${prop.key}="${prop.value}" `;
  });

  return `
  <${node.name} ${propString}>
    ${node.content}
      ${childrenHtml}
  </${node.name}>`
}

console.log(generateHTML(muiTree));