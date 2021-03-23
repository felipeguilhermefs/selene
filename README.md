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
- *Name:* should be non empty
- *Email:* should be unique
- *Password:* should be at least 8 characters long
- *Password:* should be stored with salt and pepper

## GET /books
## GET /books/create
## GET /books/:id/edit
## GET /books/:id/view
## GET /login
