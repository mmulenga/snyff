up:
	docker compose up -d --build

down:
	docker compose down

check:
	docker compose exec db psql -U postgres -d postgres -c "SELECT * FROM requests;"
