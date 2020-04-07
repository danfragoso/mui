#! /bin/bash
node tokenizer.js $1 | node frontend.js | node htmlBackend.js > out.html