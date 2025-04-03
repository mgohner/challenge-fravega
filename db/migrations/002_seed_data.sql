-- Migration: 002_seed_data
-- Insert initial seed data

-- Insert sample drivers
INSERT INTO driver (id, name, phone_number, email, address, identification, license_number) VALUES 

('e3b57a7a-fb4f-45bb-8fa6-81a406c1c596', 'John Doe', '555-123-4567', 'john.doe@example.com', '123 Main St', 'ID12345', 'LIC98765'),
('609549eb-8c70-40ea-9dc2-d7bb5bcee63e', 'Jane Smith', '555-987-6543', 'jane.smith@example.com', '456 Oak Ave', 'ID67890', 'LIC54321');

-- Insert sample vehicles
INSERT INTO vehicle (id, plate_number) VALUES 
('171f1ef5-1b5b-4fed-a4b4-9b3d2845893c', 'ABC123'),
('98fe6948-7adf-41b4-b096-2250e2ecb8db', 'XYZ789');

-- Insert sample routes
INSERT INTO route (id, name, description, status, vehicle_id, driver_id) VALUES 
('3e609a33-9bf6-4bce-9ed5-a3b1c55e34c7', 'Morning Delivery Route', 'Downtown delivery route for morning shifts', 'pending', '171f1ef5-1b5b-4fed-a4b4-9b3d2845893c', 'e3b57a7a-fb4f-45bb-8fa6-81a406c1c596'),
('356764f7-f984-437f-923d-b673d167d73b', 'Afternoon Delivery Route', 'Suburban delivery route for afternoon shifts', 'pending', '171f1ef5-1b5b-4fed-a4b4-9b3d2845893c', 'e3b57a7a-fb4f-45bb-8fa6-81a406c1c596');

-- Insert sample route points
INSERT INTO route_point (id, purchase_order_id, route_id, status, latitude, longitude, address) VALUES 
('9b7a41f2-4061-44e9-b772-30845c3f9e93', '85b01dae-d210-4ccf-a709-9ff7ba528abf', '3e609a33-9bf6-4bce-9ed5-a3b1c55e34c7', 'pending', -34.603722, -58.381592, 'Florida 165, Buenos Aires'),
('0c266de7-0a76-4d56-8828-2ac6e485d59a', 'a8093181-194d-4f16-802a-92cda4a7ef49', '3e609a33-9bf6-4bce-9ed5-a3b1c55e34c7', 'pending', -34.606761, -58.370527, 'Lavalle 750, Buenos Aires'),
('53af236a-eccc-4966-a802-3abea337ed4c', '3ab0aad5-201c-4fdf-8070-1d7df41e4db1', '3e609a33-9bf6-4bce-9ed5-a3b1c55e34c7', 'pending', -34.617737, -58.368873, 'Av. Córdoba 1111, Buenos Aires'),
('34a58f9e-1860-45ad-9572-cbe6fc8edc3c', '40d566d6-0761-4641-b261-780c3c07d283', '3e609a33-9bf6-4bce-9ed5-a3b1c55e34c7', 'pending', -34.587488, -58.419188, 'Av. Cabildo 2000, Buenos Aires'); 