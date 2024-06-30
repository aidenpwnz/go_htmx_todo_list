# todo_list_go
A todo list app made with Go, HTMX, Templ and tailwindcss. Echo was used as a web framwork.

A good exercise to learn Go from scratch as a first timer, while also peeking into HTMX.

## Getting Started

1. Clone the repository:

	git clone https://github.com/aidenpwnz/go_htmx_todo_list.git

2. Run the application:

	task run

	This command will first check if mongo is rnning, if not it will pull the image and run it. Then, it will run the required commands to init the project and finally start the Go server, compile the Tailwind CSS styles, and watch for changes in the project files.

3. Open your web browser and navigate to `http://localhost:7173` (proxied by templ) to access the Todo List app.

