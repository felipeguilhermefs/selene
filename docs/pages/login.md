
---
# Login Page

**URL:** `GET /login`

## Components
- Login Form

---
## Signup Form

### Form Fields
- Name *
- Email *
- Password *

### Flow

1. Clicking in Login button should send `POST /api/login`
1. Model validations should be done in server
1. *OK:* User should be logged in and redirected to **Books Page**
1. *Error:* Error should pop up in **Login Page**

### Requirements
- *Email*
  - 200 max length
  - Email format
  - should exist in DB
- *Password*
  - 8 min length
  - 200 max length
  - Server should compare raw with salt and pepper
