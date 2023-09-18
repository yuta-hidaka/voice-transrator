dev: 
	cd ./plugin && pnpm run dev

build:
	cd ./plugin && pnpm run build

evans:
	docker compose exec api bash -c 'evans --host localhost --port 8080 --reflection repl'

api:
	docker compose exec api bash
	