pull:
	git pull origin main

db:
	docker run --name my-postgres -d -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=mydatabase postgres
