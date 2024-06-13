1. Create a new Golang project that serve http
    - Path "/", do print string/html text of your favorite wise word
    - When you access path "/", also print to log/console using fmt.Println with text "Ouch!" Start HTTP on port 77
     ![image](https://github.com/putribow/learning-docker/assets/82745332/5e121413-2c56-4089-a764-3c5e7a6891cb)
    
2. Create file "AUTHORS.md" and "LINKS.md" within same folder with "Dockerfile
    - Edit "AUTHORS.md" file and fill it with your first name or your fake name
      ![isi dari AUTHORS md](https://github.com/putribow/learning-docker/assets/82745332/563a41da-d381-42e0-88b8-9703dcecc85a)
    - Edit "LINKS.md" file and fill it with your github profile link
      ![isi dari LINKS md](https://github.com/putribow/learning-docker/assets/82745332/a39fa8a5-70a5-4aba-8e08-1ea4cf858138)

3. Create Dockerfile to build and run your golang project
    - Use Base Image "golang:1.21"
    - Set WORKDIR with /myapp
    - RUN some command "go version" after WORKDIR
    - COPY "AUTHORS.md" to image BEFORE run build golang (go build) COPY "LINKS.md" to image BEFORE run build golang (go build)
    - Make golang build output with name "my-go-app" and make sure it will run correctly
    ![isi dari dockerfile](https://github.com/putribow/learning-docker/assets/82745332/641cc862-ac94-47b7-a7a6-9378d8242556)
    
4. Build Dockerfile image with name "my-go-app:v2" docker run --name go-app -p 5555:77 my-go-app:v2
    ![running dockerfile](https://github.com/putribow/learning-docker/assets/82745332/725fdb94-f676-48da-b117-8972752c3f44)
    ![docker image berhasil dibuat](https://github.com/putribow/learning-docker/assets/82745332/22b6ae17-b7be-4f49-bd22-955bda8b3a96)

6. Run the image into container with name "go-app" and expose to port host 5555 docker build -t my-go-app:v2 .
    ![menjalankan docker image menjadi kontainer](https://github.com/putribow/learning-docker/assets/82745332/63a25c86-21b2-41d9-bde4-7d85fe135023)
    ![kontainer go-app](https://github.com/putribow/learning-docker/assets/82745332/2b7c5f2f-bcff-47e7-b4ac-714975121168)

    
8. Access your golang inside docker via localhost:5555 and see logs of your container "go-app"
