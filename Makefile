# Makefile
.PHONY: build

BINARY_NAME=turtlor

# build builds the tailwind css sheet, and compiles the binary into a usable thing.
build:
	go mod tidy && \
   	templ generate && \
	go generate && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}

# dev runs the development server where it builds the tailwind css sheet,
# and compiles the project whenever a file is changed.
dev:
	npx tailwindcss -i ./static/css/style.css -o ./static/css/tailwind.css --watch & \
	templ generate	--watch --proxy="http://localhost:9090"  --cmd="sh -c ./dev-run"

clean:
	go clean

tailwind:
	npx tailwintdcss -i ./static/css/style.css -o ./static/css/tailwind.css --watch
