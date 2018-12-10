
echo "Compiling Go..."
gopherjs build ./src/go -m --localmap -o dist/gopherjs/rsaCrypt.js
echo "webpack..."
npm run dev
echo "Copying static files..."
cp ./src/cryptochat.html ./dist/cryptochat.html

