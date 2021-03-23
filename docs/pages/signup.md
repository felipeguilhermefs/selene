
---
# Signup Page

**URL:** `GET /signup`

## Components
- Signup Form

---
## Signup Form

### Form Fields
- Name *
- Email *
- Password *

### Flow

1. Clicking in Signup button should send `POST /api/users`
1. Model validations should be done in server
1. *OK:* User should be logged in and redirected to [Books Page](./books.md)
1. *Error:* Error should pop up in [Signup Page](./signup.md)

### Requirements
- *Name:*
  - Non blank
  - 200 max length
- *Email:*
  - Unique
  - Email format
  - 200 max length
- *Password:*
  - 8 min length
  - 200 max length
  - Server should salt and pepper
---
