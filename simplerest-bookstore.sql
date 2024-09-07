-- CRUD operations can be performed from the frontend too 

use simplerest;
select * from books;


INSERT INTO books (title, author, genre, price)
VALUES ('Book Title', 'Author Name', 'Genre', 19.99);


UPDATE books
SET title = 'New Title', price = 25.99
WHERE id = 1;



DELETE FROM books
WHERE id = 1;



