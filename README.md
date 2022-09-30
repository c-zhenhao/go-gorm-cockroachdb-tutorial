## Golang + GORM (GO Object Relational Mapping) + CockroachDB tutorial

https://www.cockroachlabs.com/docs/stable/build-a-go-app-with-cockroachdb-gorm.html?

## readme from git clone https://github.com/cockroachlabs/example-app-go-gorm

This repo contains the source code for a simple CRUD application, written in Go. The application uses [GORM](https://gorm.io/) to connect to and communicate with an existing [CockroachDB](https://www.cockroachlabs.com/docs/stable/) cluster.

For instructions on starting CockroachDB and running the code, see [this tutorial](https://www.cockroachlabs.com/docs/stable/build-a-go-app-with-cockroachdb-gorm.html).

## my notes

- using Beekeeper Studio to see the data
- <img width="879" alt="image" src="https://user-images.githubusercontent.com/16322250/193227978-72de44a5-96d3-45d1-9f17-897c785276cc.png">

- need to setup CA cert
  - ![image](https://user-images.githubusercontent.com/16322250/193229111-52fd305e-7d6d-47cb-98f5-a527c84cc559.png)
  - enable SSL (since encryption is enabled (sslmode=verify-full) in connection string)
- dont get scammed by Beekeeper's options parameters hints -- dont need
  - <img width="618" alt="image" src="https://user-images.githubusercontent.com/16322250/193229227-58ec90a1-c97e-4dfb-b379-d367a870859d.png">
  - ![image](https://user-images.githubusercontent.com/16322250/193229665-ea33c587-1085-46c6-81f8-9167e6941d89.png)
