A. Prepare a single database "(yourname)_access", single table "access _log" with single column named timestamp with data type timestamp or datetime
![membuat database putri_access](https://github.com/putribow/learning-docker/assets/82745332/357124b5-f28d-4a34-b9eb-4762f56f3503)
![membuat tabel dan kolom](https://github.com/putribow/learning-docker/assets/82745332/a8dd3533-785b-4790-ad58-3aaa03e31fe7)


B. Create or use existing golang project that serve hip and connect to postgres (or any) databases.
-Path "/", print html text of your favorite wise word
![image](https://github.com/putribow/learning-docker/assets/82745332/8323d579-5ed8-4e8d-a9f6-e06fc5b3802f)
- Also when access to "/ print to console using fmt.Printin with text "Ouch!*
![image](https://github.com/putribow/learning-docker/assets/82745332/b8da2f1c-1082-4875-a688-9c7ca1c0c9c7)

Path "/ping* -> Do ping to database and show the result (success connect or failed) in browser. Also when access to "/ping" Insert or record "NOW" or timestamp whenever /ping accessed.
![image](https://github.com/putribow/learning-docker/assets/82745332/8b20a6ca-d4b7-4a02-b11b-e5c6310fac0d)
  
- Expose to port 78
  ![image](https://github.com/putribow/learning-docker/assets/82745332/19e814b7-d506-47d3-b2cb-d0d0e42bb620)

C. Create Dockerfile to build and run your golang project
![image](https://github.com/putribow/learning-docker/assets/82745332/4385b921-be95-4066-83e5-b314663485aa)

D. Create docker-compose.yml and instruct to run 2 containers Golang Project (give name: "go service"), expose container port 78 to host port 9911
Postgres image postgreslatest (give name: "database")

![image](https://github.com/putribow/learning-docker/assets/82745332/f748568a-7b9c-4644-b778-921c1b10fa93)

![image](https://github.com/putribow/learning-docker/assets/82745332/981d7452-8659-4d57-8f0c-22185038723f)


E. Make sure the golang project can connect to postgresql and postgresql data is persistent network name = "netwwc" and volume name = "volwwc"

![image](https://github.com/putribow/learning-docker/assets/82745332/9113056a-233d-4979-b1e1-007146b25352)


F. Run Docker compose with name 'my-wise-word-compose"
![image](https://github.com/putribow/learning-docker/assets/82745332/359fc317-90e2-45fc-b081-a658440b5f22)
![image](https://github.com/putribow/learning-docker/assets/82745332/29fa8519-3a35-40f6-a765-e135ad6f6d96)

Make sure you can access http://localhost:9911 and show a wise word
![image](https://github.com/putribow/learning-docker/assets/82745332/fce67358-082b-4426-9abc-152b9d659b8c)

Make sure you can access http://localhost 9911/ping and show Success [with success ping to database) and new data inserted in your database
![image](https://github.com/putribow/learning-docker/assets/82745332/27f45fba-d620-4e1d-a7f5-a8ef2a9771be)


G. Check logs of your each container in "my-wise-word-compose"
![image](https://github.com/putribow/learning-docker/assets/82745332/f8e9cfb6-1f09-47a5-b41d-3bf2ebd6f36a)
![image](https://github.com/putribow/learning-docker/assets/82745332/6e8fbb2a-c1a2-4b0b-8563-b84516e4ac27)


H. Masukkan files & screenshot step dan hasil ke repository github kalian {username}/learning-docker Folder my-wise-word-compose

