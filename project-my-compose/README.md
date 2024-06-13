A. Prepare a single database "{username}_access", A. single table "{username}_access_log" B. with single column named "timestamp" with data type timestamp or datetime
B. Use existing golang project (from Project 1) that serve http and try to connect to postgres (or any) databases with additional:
- Path "/ping", Do ping to database and show the result
- Output or return to browser as string or html Success = "Pong!", Failed = "Ouch!”
- Print to console or log Success = "Ping successful", Failed = "Ping failed"
- Insert or record a new data timestamp with "NOW()" or current timestamp whenever /ping accessed. E.g., new record current timestamp 2024-01-31 19:55:55
- Start HTTP on port 77 (or 78 if your chrome blocked the port)
C. Create Dockerfile to build and run your golang project dockerfile 
- build
- run
D. Create docker-compose.yml and instruct to run 2 containers
Golang Project (container name: "gosvc_container"), expose container port 77 (or 78) to host port 4331 Postgres image postgres:latest (container name: "dbsvc_container")
Pastikan env postgres database "{username}_access"

E. Make sure the golang project can connect to postgresql and postgresql data is persistent
network name = "net_mycompose_{username}" and volume name = "vol_mycompose_{username}"
You can check via pgadmin and expose the port 5432 (bebas) to host.

F. Run Docker compose with name "project-my-compose"
Make sure you can access http://localhost:4331 and show a wise word
Make sure you can access http://localhost:4331/ping and show Success (with success ping to database) and new data inserted in your database

G. Check logs of your each container in "project-my-compose"

H. Masukkan semua files project & screenshot hasilnya ke repository github kalian {username}/learning-docker Folder project-my-compose

