const fs = require('fs');
const inputFile = process.argv.slice(2)[0];

const lang = require('./lang');
const inputString = fs.readFileSync(inputFile).toString();

const tokenizerOutput = {
  version: lang.version,
  filePath: inputFile, 
  tokens: []
};

function compileRegex(tokens) {
  tokens.forEach(token => {
    token.regex = new RegExp(token.pattern, token.flags); 
  });
};

function tokenize(inputString, offset) {
  let currentToken;
  let lastToken = {index: Infinity}; 
  let slicedString = inputString.slice(offset);

  for (let i = 0; i < lang.tokens.length; i++) {
    const token = lang.tokens[i];
    const tokenMatch = slicedString.match(token.regex);

    if (tokenMatch != null) {
      currentToken = tokenMatch;
      currentToken.type = token.type;

      if (currentToken.index < lastToken.index) {
        lastToken = currentToken;
      }
    }
  }

  
  const tokenValue = lastToken[0];
  
  if (tokenValue) {
    const tokenIdx = lastToken.index;
    const inputOffset = tokenIdx + tokenValue.length + offset;
    
    tokenizerOutput.tokens.push({
      value: tokenValue,
      type: lastToken.type,
      index: offset,
    });

    tokenize(inputString, inputOffset)
  }
};

console.log(`MUI Lang v${lang.version}\n-----------`);

compileRegex(lang.tokens);
tokenize(inputString, 0);
console.log(tokenizerOutput);
