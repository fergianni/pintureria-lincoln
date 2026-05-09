CREATE TABLE IF NOT EXISTS categories (
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS products (
    id          SERIAL PRIMARY KEY,
    category_id INT          REFERENCES categories(id),
    name        VARCHAR(200) NOT NULL,
    description TEXT,
    price       NUMERIC(10,2) NOT NULL,
    image_url   VARCHAR(500),
    stock       INT          NOT NULL DEFAULT 0,
    active      BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS orders (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(200) NOT NULL,
    email      VARCHAR(200) NOT NULL,
    phone      VARCHAR(50),
    notes      TEXT,
    total      NUMERIC(10,2) NOT NULL,
    status     VARCHAR(50)  NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS order_items (
    id         SERIAL PRIMARY KEY,
    order_id   INT          NOT NULL REFERENCES orders(id),
    product_id INT          NOT NULL REFERENCES products(id),
    quantity   INT          NOT NULL,
    unit_price NUMERIC(10,2) NOT NULL
);

-- Datos de ejemplo
INSERT INTO categories (name, slug) VALUES
    ('Pinturas de interior', 'interior'),
    ('Pinturas de exterior', 'exterior'),
    ('Esmaltes', 'esmaltes'),
    ('Herramientas', 'herramientas')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO products (category_id, name, description, price, stock) VALUES
    (1, 'Latex Interior Blanco 4L',  'Latex lavable de alta calidad para interiores.',  3500.00, 20),
    (1, 'Latex Interior Color 1L',   'Latex interior en variedad de colores.',           1200.00, 35),
    (2, 'Latex Exterior Blanco 4L',  'Resistente a la humedad y rayos UV.',              4200.00, 15),
    (3, 'Esmalte Sintético Negro 1L','Esmalte brillante para madera y metal.',           1800.00, 25),
    (4, 'Rodillo 23cm',              'Rodillo de lana para pintura de paredes.',          450.00, 50),
    (4, 'Pincel Cerda Fina N°20',    'Pincel para terminaciones y detalles.',             280.00, 40)
ON CONFLICT DO NOTHING;
