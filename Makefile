#build docker image
image:
	@echo "Building docker image"
	@docker build -t goapp .

clear-image:
	@echo "Clearing docker image"
	@docker rmi goapp

runim:
	@echo "Running docker image"
	@docker run --name goweb -e PORT="localhost:8000" -e SECRET_KEY="secret-taco" -p 8000:8000 goapp

stopim:
	@echo "Stopping docker image"
	@docker stop goweb
	@docker rm goweb


run:
	@echo "Running main.go"
	@air

tailwind:
	@echo "Running tailwindcss"
	@pnpm run watch

template:
	@echo "Running template"
	@templ generate --watch