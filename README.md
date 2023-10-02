# VMA-FinalProject

Hello,
Welcome to my project, this is my first personal golang project at VMA, here's something about my project.

**About the project:**

The project describes product management on an online sales website, consisting of 4 main tables: user, role, category, product

**Diagrams**
<pre>
                             ┌──────────┐
                             │ Products │
                             │          │
                ┌───────── n │          │ n ─────────┐
                │            │          │            │
                │            │          │            │
                │            └──────────┘            │
                │                                    │
                │1                                   │1
           ┌────┴──────┐                        ┌────┴─────┐
           │ Categories│                        │   Users  │
           │           │                        │          │
           │           │                        │          │
           │           │                        │          │
           └───────────┘                        └──────────┘
                                                     │n
                                                     │
                                                     │
                                                     │
                                                     │1
                                                ┌────┴─────┐
                                                │   Roles  │
                                                │          │
                                                │          │
                                                │          │
                                                └──────────┘
</pre>

**Project functions:**
- Register, log in, log out and authenticate with JWT, decentralize with RBAC
- CRUD with user and admin roles, specifically: user role can only CRUD themselves, can only view products but cannot add, edit, delete products. Meanwhile, the admin has full rights to users, categories, and products but cannot CRUD other admins.
- Reset new password randomly and send to email
- Post photos to Azamon S3 Bucket
- Uses Redis in project



**Project structure:**
<pre>
VMA-FINALPROJECT/
│
│
├── controllers/
│
├── database/
│   ├── migrations/
│   └── seeder/
│
├── middleware/
│
├── models/
│   
├── routers/
│
├── services/
│
├── utils/
│
├── .env
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md
</pre>


**How to run it?**
1. Clone this project to your computer by CLI git clone
2. Enter website: https://go.dev/dl/ to download golang package 1.20.3
3. Enter website: https://www.docker.com/ to download docker and run docker desktop
4. Open the project with any IDE you have (VScode, intelIJ...) and step by step following part: 
- Ccreate a ".env" file and setup your environment based on the template in the ".env.example" file available in the project
- Terminal: typing "docker-compose up -d" to pull Images docker
- Terminal: typing "go mod download" to pull Dependencies in project
5. ALL SET ARE DONE, typing "go run main.go" to run your project


**Bugs during project starting on:**
1. If an error "failed to initialize database..." try running this command: "docker-compose down" at the terminal of the project directory
2. If there is a redis initialization error, try changing the following value in the .env file: "localhost:6379" to "redis:6379". Or try changing to another port
3. If there is any error with go mod or go sum, run the command "go mod tidy" (not recommended)
4. If you have any questions, please contact me by email: nguyenhalinh.1709@gmail.com

**Attention:**

I have put sample body/raw if you want to use for testing on postman, swagger or any other API testing tool, you will find them in "postman.txt" inside the project.

**THANK YOU**

