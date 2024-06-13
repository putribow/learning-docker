1. Create a new Golang project that serve http
    - Path "/", do print string/html text of your favorite wise word
    - When you access path "/", also print to log/console using fmt.Println with text "Ouch!" Start HTTP on port 77
2. Create file "AUTHORS.md" and "LINKS.md" within same folder with "Dockerfile
    - Edit "AUTHORS.md" file and fill it with your first name or your fake name
    - Edit "LINKS.md" file and fill it with your github profile link
3. Create Dockerfile to build and run your golang project
    - Use Base Image "golang:1.21"
    - Set WORKDIR with /myapp
    - RUN some command "go version" after WORKDIR
    - COPY "AUTHORS.md" to image BEFORE run build golang (go build) COPY "LINKS.md" to image BEFORE run build golang (go build)
    - Make golang build output with name "my-go-app" and make sure it will run correctly
4. Build Dockerfile image with name "my-go-app:v2" docker run --name go-app -p 5555:77 my-go-app:v2
5. Run the image into container with name "go-app" and expose to port host 5555 docker build -t my-go-app:v2 .
6. Access your golang inside docker via localhost:5555 and see logs of your container "go-app"