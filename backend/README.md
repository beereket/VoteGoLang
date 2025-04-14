# Online Elections (GO Project)

## Overview

This project provides an online election system built with Go. The application includes voting functionalities for various candidates (President, Deputy, Session Deputy) and petitions. It uses JWT (JSON Web Token) for authentication and authorization.

### Key Features:
- Voting for candidates (President, Deputy, Session Deputy)
- Voting on petitions (In favor, Against)
- Viewing and creating petitions
- Viewing general news
- JWT-based authentication for secure access

## Endpoints

### Authentication:
- **Login**: `/login`  
  - Method: `Get`
  - Description: Authenticates the user and returns a JWT token.
  
- **Register**: `/register`  
  - Method: `POST`
  - Description: Registers a new user and returns a JWT token.

### Candidates:
- **Get Candidates by Type**: `/candidate?type="type"`  
  - Method: `GET`
  - Query Parameter: `type` (can be `president`, `deputy`, `session_deputy`)
  - Description: Retrieves a list of all candidates based on the specified type.

### Voting:
- **Vote for Candidate**: `/vote`  
  - Method: `POST`
  - Headers: `Authorization: Bearer <JWT Token>`  
  - Request Body:
    ```json
    {
      "candidate_id": "candidate_id",
      "candidate_type": "candidate_type"
    }
    ```
  - Description: Allows a user to vote for a candidate. The user’s ID is extracted from the JWT token.

- **Get General News**: `/general_news`  
  - Method: `GET`
  - Description: Retrieves a list of all general news.

### Petition:
- **Create Petition**: `/petition`  
  - Method: `POST`
  - Request Body:
    ```json
    {
      "title": "Petition title",
      "description": "Petition description"
    }
    ```
  - Description: Allows users to create a new petition.

- **Get All Petitions**: `/petition`  
  - Method: `GET`
  - Description: Retrieves a list of all petitions.

- **Vote on Petition**: `/petition_vote`  
  - Method: `POST`
  - Request Body:
    ```json
    {
      "petition_id": "petition_id",
      "vote": "infavor/against"
    }
    ```
  - Description: Allows users to vote on a petition, either "In favor" or "Against".

### JWT Authentication:
- JWT tokens are used for authentication. Upon successful login or registration, the user receives a JWT token.
- For protected endpoints like voting, the JWT token should be included in the `Authorization` header with the prefix `Bearer`.

### Types of Votes:
- **Candidates**: Users can vote for candidates in the following categories:
  - President
  - Deputy
  - Session Deputy

- **Petition**: Users can vote on petitions with the following options:
  - In favor
  - Against
---

## Use Cases

This project supports the following use cases. Each use case is documented in the respective file in the GitHub repository.

### [User Registration/Login](https://github.com/DarkhanTastanov/VoteGolang/blob/master/internals/usecases/auth_usecase.go)

### [Candidates](https://github.com/DarkhanTastanov/VoteGolang/blob/master/internals/usecases/candidate_usecase.go)

### [Petition](https://github.com/DarkhanTastanov/VoteGolang/blob/master/internals/usecases/petition_usecase.go)

### [News](https://github.com/DarkhanTastanov/VoteGolang/blob/master/internals/usecases/general_news_usecase.go)

-
