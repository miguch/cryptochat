echo "Compiling TypeScript..."
tsc
echo "Compiling Go..."
gopherjs build ./src/go --localmap -o dist/gopherjs/rsaCrypt.js

