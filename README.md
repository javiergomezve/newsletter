# Newsletter

This project is a full-stack application built with Next.js (frontend) and Golang (backend), including a PostgreSQL database. You can easily run the entire project using Docker Compose.

## Prerequisites

Make sure you have Docker and Docker Compose installed on your machine.

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## How to Run

To run this project, use the following command:

```bash
docker-compose up --build
```

This command will build and start the Docker containers for the frontend (Next.js), backend (Golang), and PostgreSQL database.

The frontend will be accessible at [http://localhost:3000](http://localhost:3000), the backend at [http://localhost:8080](http://localhost:8080), and the PostgreSQL database at [localhost:5431](localhost:5431).

## Database Connection

The PostgreSQL database is included in the Docker Compose setup. You can configure the database connection in your backend code or environment variables.

- **Host:** localhost
- **Port:** 5431
- **Username:** postgres
- **Password:** postgres
- **Database:** postgres

## Notes

- Ensure that no other services are already using the specified ports on your machine before running the `docker-compose` command.

- If you encounter any issues or need further customization, please refer to the project's documentation or consult the respective documentation for Next.js, Golang, and PostgreSQL.
