
---
# Add Book Page

**URL:** `GET /books/:id/edit`

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
1. *OK:* User should be redirected to [Books Page](./books)
1. *Error:* Error should pop up in [New Book Page](./new_book)

### Requirements
- *Title:*
  - Non blank
  - 200 max length
- *Author:*
  - 200 max length
- *Commments:*
  - 2000 max length
- *Tags:*
  - Should allow multiple tags
  - 20 tags max
  - Each tag should have 20 characters max

---
