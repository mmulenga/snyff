up:
	sudo docker compose up -d --build

down:
	sudo docker compose down

check:
	sudo docker compose exec db psql -U postgres -d postgres -c "SELECT * FROM requests;"
	