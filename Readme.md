# build an image from Dockerfile
- > docker build -t docker-test .

# check image
- > docker image ls

# run docker image 
- > docke run -p 8080:8081 -it docker-test
=>8080 is localhost and 8081 where server is running)
=> open any browser, hit url localhost:8080 

# list the container
- > docker container ls --al

# run an interactive sheel in container
- >  docker exec -it  sh