# Used to automatically regenerate & refresh when a template file changes
run/templ:
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v

# Used to automatically regenerate & refresh when a css file changes
run/tailwind:
	tailwindcss -i ./input.css -o ./dist/output.css --minify --watch

# Used to automatically recompile & refresh on a go file change
run/server:
	air

# Move the required file to the dist folder, to allow compatibility with Flowbite UI library
init/flowbite:
	cp ./node_modules/flowbite/dist/flowbite.min.js ./dist/flowbite.min.js

# initialize the project
init:
	npm install && make init/flowbite 

# run the project
run:
	make -j5 run/server run/templ run/tailwind init/flowbite