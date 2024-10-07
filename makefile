run_build:
	chmod +x ./database/auth_scripts/schema.sql
	chmod +x ./database/memory_scripts/schema.sql
	docker-compose up --build
	
run:
	chmod +x ./database/auth_scripts/schema.sql
	chmod +x ./database/memory_scripts/schema.sql
	docker-compose up

stop:
	docker-compose down

reset:
	docker-compose down -v