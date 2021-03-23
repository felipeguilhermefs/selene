
---
# Login Page

**URL:** `GET /login`

## Components
- Login Form

---
## Login Form

### Form Fields
- Email *
- Password *

### Flow

1. Clicking in Login button should send `POST /api/login`
1. Model validations should be done in server
1. *OK:* User should be logged in and redirected to [Books Page](./books.md)
1. *Error:* Error should pop up in [Login Page](./login.md)

