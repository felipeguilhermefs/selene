
---
# Books Page

**URL:** `GET /books`

*User should be logged in

## Components
- Bookshelf
- Book Search
- New book

---
## Bookshelf

Pageable Table roughly like:

| Title | Author | Tags | Actions |
| ----- | ------ | ---- | ------- |
| The Hobbit | JRR Tolkien | adventure, fantasy | Edit , Delete, View |
| The Raven | Edgar Allan Poe | terror | Edit , Delete, View |
| | | < 1/10 > | |

Default Page Size: 10

### Flow

1. Clicking in Edit button should direct to [Edit Book Page](./edit_book)
1. Clicking in View button should direct to [View Book Page](./view_book)
1. Clicking in View button should send `DELETE /api/books/:id`
   - *OK:* should show an sucess message
   - *Error:* should show error message
1. Clicking in *Previous Page Button (<)* should show previous page. If user on first page it should be disabled
1. Clicking in *Next Page Button (>)* should show next page. If user on last page it should be disabled

---
## Book Search

### Form Fields
- Search Text

### Flow

1. Clicking in Search button should send `GET /books?search={text}`
1. Server should filter books based on search text:
   - **free text:** books should be found if *Title*, *Author* or *Tags* **contains** text.
   - text starts with `author:`, *Author:* contains text
   - text starts with `title:`, *Title:* contains text
   - text starts with `tag:`, *Tags:* contains text
   - blank text should not filter books
1. User will be shown [Books Page](./books) with filtered books in *Bookshelf* component

### Requirements
- *Search Text*
  - 200 max length

---
## New book

### Flow

1. Clicking in *New Book* button should redirect to [New Book Page](./new_book)

---
