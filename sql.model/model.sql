CREATE TABLE book (
    id UUID DEFAULT gen_random_uuid(),
    title varchar(24) NOT NULL,
    price INT NOT NULL,
    author_name varchar NOT NULL,
    published_year INT NOT NULL,
    page INT NOT NULL,
    genre varchar

);

INSERT INTO book (title, price, author_name, published_year, page, genre) VALUES
('Harry Potter', 200000,'JK Rowling', 2000, 600, 'fiction' ),
('A Little Life', 255000,'Hanya Yanagihara', 2015, 359, 'historical' ),
('The Siege', 400000,'Helen Dunmore', 2001, 250, 'romance' ),
('Thinking, Fast and Slow', 600000,'Daniel Kahneman ', 2011, 1000, 'science' ),
('Days Without End', 500000,' Sebastian Barry', 2016, 678, 'fantasy' )
;
