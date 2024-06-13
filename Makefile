# #build docker image
# image:
# 	@echo "Building docker image"
# 	@docker build -t goapp .

# clear:
# 	@echo "Clearing docker image"
# 	@docker rmi goapp

# run:
# 	@echo "Running docker image"
# 	@docker run --name goweb -p 8000:8000 goapp

# stop:
# 	@echo "Stopping docker image"
# 	@docker stop goweb
# 	@docker rm goweb


run:
	@echo "Running main.go"
	@air

tailwind:
	@echo "Running tailwindcss"
	@pnpm run watch

template:
	@echo "Running template"
	@templ generate --watch