A. Prepare a single database "{username}_access", A. single table "{username}_access_log" B. with single column named "timestamp" with data type timestamp or datetime
![membuat database putri_access](https://github.com/putribow/learning-docker/assets/82745332/739a1d30-c5e1-40b6-9025-ac35ee4f9a29)
![membuat tabel access_log](https://github.com/putribow/learning-docker/assets/82745332/70dbcd1c-721b-438b-8aac-0dbfd24af2a3)
![timestamp kolom](https://github.com/putribow/learning-docker/assets/82745332/6eafcac8-2de3-40cc-bac9-52a3cca3d41c)

B. Use existing golang project (from Project 1) that serve http and try to connect to postgres (or any) databases with additional:
- Path "/ping", Do ping to database and show the result
- Output or return to browser as string or html Success = "Pong!", Failed = "Ouch!”
- Print to console or log Success = "Ping successful", Failed = "Ping failed"
- Insert or record a new data timestamp with "NOW()" or current timestamp whenever /ping accessed. E.g., new record current timestamp 2024-01-31 19:55:55
- Start HTTP on port 77 (or 78 if your chrome blocked the port)
![codingan](https://github.com/putribow/learning-docker/assets/82745332/53089cd0-0430-43cc-b324-f562237c1d20)
![menjalankan app golang](https://github.com/putribow/learning-docker/assets/82745332/d89d2503-48be-48fe-8463-6e10c293e4fe)
![menjalankan app golang dengan ping](https://github.com/putribow/learning-docker/assets/82745332/a32a662e-c40a-4ab4-85d2-5c032e6afc50)


C. Create Dockerfile to build and run your golang project dockerfile

- run
  ![image](https://github.com/putribow/learning-docker/assets/82745332/add1cd04-5f63-4085-8463-4ca7552e77d5)

D. Create docker-compose.yml and instruct to run 2 containers
Golang Project (container name: "gosvc_container"), expose container port 77 (or 78) to host port 4331 Postgres image postgres:latest (container name: "dbsvc_container")
Pastikan env postgres database "{username}_access"
![docker-compose yml](https://github.com/putribow/learning-docker/assets/82745332/03481e3f-76ba-44bd-8ff7-60ed0cbb95d1)
![docker-compose yml 2](https://github.com/putribow/learning-docker/assets/82745332/866fd755-3603-4ccb-8680-d50d32e93f18)


E. Make sure the golang project can connect to postgresql and postgresql data is persistent
network name = "net_mycompose_{username}" and volume name = "vol_mycompose_{username}"
You can check via pgadmin and expose the port 5432 (bebas) to host.


F. Run Docker compose with name "project-my-compose"
Make sure you can access http://localhost:4331 and show a wise word
Make sure you can access http://localhost:4331/ping and show Success (with success ping to database) and new data inserted in your database
![docker compose ](https://github.com/putribow/learning-docker/assets/82745332/d3f64231-1a0c-48e9-943a-fe5c7318dd6c)


G. Check logs of your each container in "project-my-compose"
![logs gosvc_container](https://github.com/putribow/learning-docker/assets/82745332/d67c75f0-bb76-469d-8706-65d651012b66)

H. Masukkan semua files project & screenshot hasilnya ke repository github kalian {username}/learning-docker Folder project-my-compose

