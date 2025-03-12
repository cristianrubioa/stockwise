# 🚀 Stockwise

![Screenshot from 2025-03-11 22-26-54](https://github.com/user-attachments/assets/8ae22332-1f9e-4be6-8c32-78e43ba91061)


## 🛠️ Technologies Used

- **Backend:** Golang
- **UI:** Vue 3 with TypeScript, Pinia, and Tailwind CSS
- **Database:** CockroachDB



## 🚧 Installation and Setup

### Prerequisites

- [CockroachDB](https://www.cockroachlabs.com/docs/stable/install-cockroachdb)
- [Go (1.21+)](https://golang.org/dl/)

### Steps to Set Up the Project

Clone the repository:
```bash
git clone git@github.com:cristianrubioa/stockwise.git
cd stockwise
```

### 1. Start the Database

Run CockroachDB locally:

```bash
cockroach start-single-node --insecure \
  --listen-addr=localhost \
  --http-addr=localhost:8081 \
  --store=cockroach-data
```

### 2. Initialize Database and Data

In another terminal, execute the initial setup script:

```bash
./setup.sh
```

This will initialize the database, create the necessary table, and load initial data from the API.

> **Note**: The script checks if data already exists to avoid duplication.

---

## 🚀 Running the Project

After setting up and initializing the database:

### Backend

Start the backend server:

```bash
go run cmd/main.go
```

This exposes the API to query and manage data.

### Frontend

```bash
npm run dev
```

-----

![Screenshot from 2025-03-11 22-28-26](https://github.com/user-attachments/assets/5b64a2de-f86f-4e48-b980-14edd1682859)


