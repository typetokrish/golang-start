
all:

	echo "RUN"
dev-build:
	echo "Building for Development"
	docker build \
         --build-arg USER_ID=$(shell id -u) \
         --build-arg GROUP_ID=$(shell id -g) \
         -t welcome-app-dev . 
dev-run:
	echo "Running Development"
	docker run -it -d --rm -p 8080:8080 welcome-app-dev	 

