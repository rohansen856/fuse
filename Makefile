pull:
	git pull origin main

db:
<<<<<<< Updated upstream
	docker run --name my-postgres -d -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=mydatabase postgres:16 & npx prisma db push
=======
	docker run --name my-postgres -d -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=mydatabase postgres:16
>>>>>>> Stashed changes
