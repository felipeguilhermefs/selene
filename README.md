# Pages:

*FE will be changed to a rect app in the future*

## Signup Page

**URL:** `GET /signup`

### Form Fields:
- Name *
- Email *
- Password *

### Flow

1. Clicking in Signup button should send `POST /api/users`
1. Model validations should be done in server
1. *OK:* User should be logged in and redirected to **Books Page**
1. *Error:* Error should pop up in **Signup Page**

### Requirements
- *Name:*
  - Non blank
  - 200 max length
- *Email:*
  - Unique
  - email format
  - 200 max length
- *Password:*
  - 8 min length
  - 200 max length
  - Server should salt and pepper

## GET /books
## GET /books/create
## GET /books/:id/edit
## GET /books/:id/view
## GET /login

# Next Steps:
- Use external Auth provider like Auth0
- Drop custom user CRUD
