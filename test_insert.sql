-- ============================================================
--  Restaurant API Database — Full Setup & Seed Script
--  Run this from the default 'postgres' database:
--    psql -U postgres -f gorest_setup.sql
-- ============================================================

-- --------------------------------------------------------
--  1. CREATE DATABASE & CONNECT
-- --------------------------------------------------------
DROP DATABASE IF EXISTS gorest;
CREATE DATABASE gorest;
\c gorest

-- --------------------------------------------------------
--  2. CREATE TABLES
-- --------------------------------------------------------

CREATE TABLE burgers (
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(100) NOT NULL,
    price    NUMERIC(6,2) NOT NULL,
    calories INT          NOT NULL,
    is_meal  BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE TABLE drinks (
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(100) NOT NULL,
    price    NUMERIC(6,2) NOT NULL,
    calories INT          NOT NULL,
    is_iced  BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE TABLE sides (
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(100) NOT NULL,
    price    NUMERIC(6,2) NOT NULL,
    calories INT          NOT NULL
);

CREATE TABLE desserts (
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(100) NOT NULL,
    price    NUMERIC(6,2) NOT NULL,
    calories INT          NOT NULL
);

CREATE TABLE orders (
    order_number BIGSERIAL    PRIMARY KEY,
    total_price  NUMERIC(8,2) NOT NULL,
    date_time    TIMESTAMP    NOT NULL DEFAULT NOW(),
    status       VARCHAR(20)  NOT NULL CHECK (status IN ('pending', 'completed', 'cancelled'))
);

CREATE TABLE order_items (
    id           SERIAL       PRIMARY KEY,
    order_number BIGINT       NOT NULL REFERENCES orders(order_number) ON DELETE CASCADE,
    item_name    VARCHAR(100) NOT NULL,
    item_type    VARCHAR(20)  NOT NULL CHECK (item_type IN ('burger', 'drink', 'side', 'dessert')),
    quantity     INT          NOT NULL DEFAULT 1,
    price        NUMERIC(6,2) NOT NULL
);

-- --------------------------------------------------------
--  3. SEED MENU TABLES
-- --------------------------------------------------------

-- BURGERS (15 items)
INSERT INTO burgers (name, price, calories, is_meal) VALUES
('Classic Smash Burger',            8.99,   540, FALSE),
('Classic Smash Burger Meal',      12.49,   980, TRUE),
('Double Beast Burger',            11.99,   820, FALSE),
('Double Beast Burger Meal',       15.49,  1260, TRUE),
('Crispy Chicken Sandwich',         9.49,   610, FALSE),
('Crispy Chicken Sandwich Meal',   13.49,  1050, TRUE),
('BBQ Bacon Burger',               10.99,   740, FALSE),
('BBQ Bacon Burger Meal',          14.49,  1180, TRUE),
('Mushroom Swiss Burger',          10.49,   680, FALSE),
('Spicy Jalapeño Burger',          10.49,   690, FALSE),
('Veggie Black Bean Burger',        9.99,   420, FALSE),
('Fish Fillet Sandwich',            9.99,   530, FALSE),
('Triple Stack Burger',            14.99,  1150, FALSE),
('Truffle Butter Burger',          13.49,   760, FALSE),
('Hawaiian Teriyaki Burger',       11.49,   710, FALSE);

-- DRINKS (12 items)
INSERT INTO drinks (name, price, calories, is_iced) VALUES
('Coca-Cola',              2.49,  250, TRUE),
('Sprite',                 2.49,  230, TRUE),
('Fanta Orange',           2.49,  260, TRUE),
('Iced Lemonade',          2.99,  180, TRUE),
('Strawberry Milkshake',   4.99,  580, TRUE),
('Chocolate Milkshake',    4.99,  620, TRUE),
('Vanilla Milkshake',      4.99,  560, TRUE),
('Brewed Coffee',          1.99,    5, FALSE),
('Hot Chocolate',          2.99,  320, FALSE),
('Iced Tea',               2.49,  120, TRUE),
('Sparkling Water',        1.79,    0, TRUE),
('Mango Smoothie',         5.49,  340, TRUE);

-- SIDES (12 items)
INSERT INTO sides (name, price, calories) VALUES
('Regular Fries',                3.49, 380),
('Large Fries',                  4.49, 520),
('Onion Rings',                  4.29, 450),
('Mozzarella Sticks',            5.49, 490),
('Mac & Cheese Bites',           5.29, 460),
('Coleslaw',                     2.99, 180),
('Side Salad',                   3.49, 120),
('Sweet Potato Fries',           4.49, 360),
('Loaded Nachos',                6.49, 680),
('Chicken Tenders (4pc)',        5.99, 520),
('Jalapeño Poppers',             5.29, 410),
('Garlic Parmesan Wings (6pc)',  7.49, 590);

-- DESSERTS (10 items)
INSERT INTO desserts (name, price, calories) VALUES
('Chocolate Brownie',            3.99, 410),
('Warm Cookie Sundae',           5.49, 620),
('Churros (3pc)',                4.29, 340),
('New York Cheesecake Slice',    5.99, 480),
('Apple Pie Slice',              4.49, 350),
('Soft Serve Ice Cream',         2.99, 240),
('Oreo Cookie Crumble',          5.49, 560),
('Cinnamon Sugar Donuts',        3.99, 380),
('Salted Caramel Tart',          5.29, 440),
('Banana Split',                 6.49, 720);

-- --------------------------------------------------------
--  4. SEED ORDERS & ORDER ITEMS
--     Total price = SUM(price * quantity) per order
-- --------------------------------------------------------

/*
  ORDER 1 — completed
  Classic Smash Burger  x1   8.99
  Regular Fries         x1   3.49
  Coca-Cola             x1   2.49
  ──────────────────────── 14.97
*/
INSERT INTO orders (total_price, date_time, status)
VALUES (14.97, '2025-02-24 11:15:00', 'completed');

INSERT INTO order_items (order_number, item_name, item_type, quantity, price) VALUES
(1, 'Classic Smash Burger', 'burger',  1,  8.99),
(1, 'Regular Fries',        'side',    1,  3.49),
(1, 'Coca-Cola',            'drink',   1,  2.49);

/*
  ORDER 2 — completed
  Double Beast Burger Meal  x1  15.49
  Chocolate Milkshake       x1   4.99
  ──────────────────────────── 20.48
*/
INSERT INTO orders (total_price, date_time, status)
VALUES (20.48, '2025-02-24 11:42:00', 'completed');

INSERT INTO order_items (order_number, item_name, item_type, quantity, price) VALUES
(2, 'Double Beast Burger Meal', 'burger', 1, 15.49),
(2, 'Chocolate Milkshake',      'drink',  1,  4.99);

/*
  ORDER 3 — completed
  BBQ Bacon Burger     x2  10.99 each = 21.98
  Onion Rings          x1   4.29
  Iced Lemonade        x2   2.99 each =  5.98
  ─────────────────────────────────── 32.25
*/
INSERT INTO orders (total_price, date_time, status)
VALUES (32.25, '2025-02-24 12:05:00', 'completed');

INSERT INTO order_items (order_number, item_name, item_type, quantity, price) VALUES
(3, 'BBQ Bacon Burger', 'burger', 2, 10.99),
(3, 'Onion Rings',      'side',   1,  4.29),
(3, 'Iced Lemonade',    'drink',  2,  2.99);

/*
  ORDER 4 — cancelled
  Triple Stack Burger       x1  14.99
  Large Fries               x1   4.49
  Strawberry Milkshake      x1   4.99
  Chocolate Brownie         x1   3.99
  ─────────────────────────────── 28.46
*/
INSERT INTO orders (total_price, date_time, status)
VALUES (28.46, '2025-02-24 12:30:00', 'cancelled');

INSERT INTO order_items (order_number, item_name, item_type, quantity, price) VALUES
(4, 'Triple Stack Burger',   'burger',  1, 14.99),
(4, 'Large Fries',           'side',    1,  4.49),
(4, 'Strawberry Milkshake',  'drink',   1,  4.99),
(4, 'Chocolate Brownie',     'dessert', 1,  3.99);

/*
  ORDER 5 — completed
  Crispy Chicken Sandwich Meal  x1  13.49
  Mozzarella Sticks             x1   5.49
  Sprite                        x1   2.49
  Warm Cookie Sundae            x1   5.49
  ──────────────────────────────── 26.96
*/
INSERT INTO orders (total_price, date_time, status)
VALUES (26.96, '2025-02-24 13:10:00', 'completed');

INSERT INTO order_items (order_number, item_name, item_type, quantity, price) VALUES
(5, 'Crispy Chicken Sandwich Meal', 'burger',  1, 13.49),
(5, 'Mozzarella Sticks',            'side',    1,  5.49),
(5, 'Sprite',                       'drink',   1,  2.49),
(5, 'Warm Cookie Sundae',           'dessert', 1,  5.49);

/*
  ORDER 6 — completed
  Veggie Black Bean Burger  x1   9.99
  Side Salad                x1   3.49
  Sparkling Water           x1   1.79
  ──────────────────────────── 15.27
*/
INSERT INTO orders (total_price, date_time, status)
VALUES (15.27, '2025-02-24 13:35:00', 'completed');

INSERT INTO order_items (order_number, item_name, item_type, quantity, price) VALUES
(6, 'Veggie Black Bean Burger', 'burger', 1, 9.99),
(6, 'Side Salad',               'side',   1, 3.49),
(6, 'Sparkling Water',          'drink',  1, 1.79);

/*
  ORDER 7 — pending
  Truffle Butter Burger      x1  13.49
  Sweet Potato Fries         x1   4.49
  Mango Smoothie             x1   5.49
  Salted Caramel Tart        x1   5.29
  ──────────────────────────────── 28.76
*/
INSERT INTO orders (total_price, date_time, status)
VALUES (28.76, '2025-02-24 14:02:00', 'pending');

INSERT INTO order_items (order_number, item_name, item_type, quantity, price) VALUES
(7, 'Truffle Butter Burger', 'burger',  1, 13.49),
(7, 'Sweet Potato Fries',    'side',    1,  4.49),
(7, 'Mango Smoothie',        'drink',   1,  5.49),
(7, 'Salted Caramel Tart',   'dessert', 1,  5.29);

/*
  ORDER 8 — pending
  Hawaiian Teriyaki Burger  x2  11.49 each = 22.98
  Loaded Nachos             x1   6.49
  Fanta Orange              x2   2.49 each =  4.98
  ──────────────────────────────────────── 34.45
*/
INSERT INTO orders (total_price, date_time, status)
VALUES (34.45, '2025-02-24 14:28:00', 'pending');

INSERT INTO order_items (order_number, item_name, item_type, quantity, price) VALUES
(8, 'Hawaiian Teriyaki Burger', 'burger', 2, 11.49),
(8, 'Loaded Nachos',            'side',   1,  6.49),
(8, 'Fanta Orange',             'drink',  2,  2.49);

/*
  ORDER 9 — cancelled
  Fish Fillet Sandwich     x1   9.99
  Coleslaw                 x1   2.99
  Iced Tea                 x1   2.49
  ──────────────────────────── 15.47
*/
INSERT INTO orders (total_price, date_time, status)
VALUES (15.47, '2025-02-24 14:55:00', 'cancelled');

INSERT INTO order_items (order_number, item_name, item_type, quantity, price) VALUES
(9, 'Fish Fillet Sandwich', 'burger', 1, 9.99),
(9, 'Coleslaw',             'side',   1, 2.99),
(9, 'Iced Tea',             'drink',  1, 2.49);

/*
  ORDER 10 — pending
  Mushroom Swiss Burger         x1  10.49
  Spicy Jalapeño Burger         x1  10.49
  Chicken Tenders (4pc)         x1   5.99
  Jalapeño Poppers              x1   5.29
  Vanilla Milkshake             x2   4.99 each = 9.98
  Churros (3pc)                 x1   4.29
  New York Cheesecake Slice     x1   5.99
  ──────────────────────────────────────── 52.52
*/
INSERT INTO orders (total_price, date_time, status)
VALUES (52.52, '2025-02-24 15:20:00', 'pending');

INSERT INTO order_items (order_number, item_name, item_type, quantity, price) VALUES
(10, 'Mushroom Swiss Burger',       'burger',  1, 10.49),
(10, 'Spicy Jalapeño Burger',       'burger',  1, 10.49),
(10, 'Chicken Tenders (4pc)',        'side',    1,  5.99),
(10, 'Jalapeño Poppers',            'side',    1,  5.29),
(10, 'Vanilla Milkshake',           'drink',   2,  4.99),
(10, 'Churros (3pc)',               'dessert', 1,  4.29),
(10, 'New York Cheesecake Slice',   'dessert', 1,  5.99);

-- --------------------------------------------------------
--  5. VERIFICATION QUERY
--     Run this to confirm all order totals match
-- --------------------------------------------------------
SELECT
    o.order_number,
    o.total_price                                    AS stored_total,
    SUM(oi.price * oi.quantity)                      AS computed_total,
    o.total_price = SUM(oi.price * oi.quantity)      AS totals_match,
    o.status,
    COUNT(oi.id)                                     AS item_count
FROM orders o
JOIN order_items oi ON oi.order_number = o.order_number
GROUP BY o.order_number, o.total_price, o.status
ORDER BY o.order_number;