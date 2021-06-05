module donation-playlist-server

go 1.16

require internal/users v1.0.0

replace internal/users => ./internal/users

require internal/database v1.0.0

replace internal/database => ./internal/database

require (
	github.com/gorilla/mux v1.8.0
	go.mongodb.org/mongo-driver v1.5.3
)
