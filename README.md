# VoteGolang 🎓

A full-stack **Online Voting System** built with:

- 🛠️ **Golang** (Backend API)
- 💃 **Vue.js** (Frontend SPA)
- 💃 **Chart.js** (for beautiful charts)
- 📃 **MySQL** (Database)

---

## 🌟 Features

- User Registration & Login (JWT Authentication)
- Admin & User Roles (Role-based access control)
- Create & Manage Elections
- Vote for Candidates (only once per election)
- View Live Election Results (Charts)
- Create & Manage Petitions
- Vote on Petitions (In Favor / Against)
- View Petition Results
- Admin Dashboard with Analytics:
    - Total Votes, Users, Petitions
    - Top Candidates
- Activity Logs for Admins
- Secure Backend Middleware (Authorization, Ban Checking)
- Auto-close Elections after election date

---

## 🚀 How to Run Locally

### Backend (Golang API)

```bash
# 1. Clone the repository
git clone https://github.com/your-username/VoteGolang.git

# 2. Navigate into project
cd VoteGolang

# 3. Install dependencies
go mod tidy

# 4. Set environment variables (e.g., in a `.env` file or manually):
#    - DB_USER
#    - DB_PASSWORD
#    - DB_NAME
#    - JWT_SECRET

# 5. Run the server
go run main.go
```

Server will run on `http://localhost:8080`.

---

### Frontend (Vue.js)

```bash
# 1. Navigate into frontend folder
cd web

# 2. Install dependencies
npm install

# 3. Run the development server
npm run serve
```

Frontend will run on `http://localhost:8081`.

---

## 🌍 Technologies Used

| Layer      | Tech             |
|------------|------------------|
| Backend    | Go, Gorilla Mux   |
| Frontend   | Vue.js (Vue 2)    |
| Charts     | Chart.js          |
| Database   | MySQL             |
| Auth       | JWT Tokens        |

---

# 🎯 Author

Made with ❤️ by **Bereket Yergali**

---

