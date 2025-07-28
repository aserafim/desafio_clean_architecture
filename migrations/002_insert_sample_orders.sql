-- migrations/002_insert_sample_orders.sql
INSERT INTO orders (id, product, price, created_at)
VALUES
  ('550e8400-e29b-41d4-a716-446655440001', 'Notebook Dell XPS', 8999.90, NOW() - INTERVAL '2 days'),
  ('550e8400-e29b-41d4-a716-446655440002', 'Monitor LG Ultrawide', 1299.50, NOW() - INTERVAL '1 day'),
  ('550e8400-e29b-41d4-a716-446655440003', 'Teclado Mecânico', 499.99, NOW());
