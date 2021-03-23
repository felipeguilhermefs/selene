
---
# Book Page

**URL:** `GET /books/:id`

*User should be logged in

## Components
- Edit Book Form

---
## Edit Book Form

### Form Fields
- Title *
- Author
- Comments
- Tags

### Flow

1. Clicking in Save button should send `PUT /api/books/:id`
1. Model validations should be done in server
1. *OK:* User should be redirected to [Books Page](./books.md)
1. *Error:* Error should pop up in [Book Page](./book.md)

---
