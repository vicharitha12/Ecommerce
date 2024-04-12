import React, { useState, useEffect } from 'react';
import axios from 'axios';

const ListItems = () => {
  const [items, setItems] = useState([]);

  useEffect(() => {
    const fetchItems = async () => {
      const response = await axios.get('http://your-backend-url/items');
      setItems(response.data);
    };

    fetchItems();
  }, []);

  const handleAddToCart = (itemId) => {
    // Add item to cart logic
  };

  const handleCheckout = () => {
    // Convert cart to order logic
    // Show toast notification "Order successful"
  };

  const handleShowCart = () => {
    // Show all added items in the cart in a toast or alert
  };

  const handleShowOrderHistory = () => {
    // Show all placed Order ids in a toast or alert
  };

  return (
    <div>
      <h2>Items List</h2>
      <button onClick={handleCheckout}>Checkout</button>
      <button onClick={handleShowCart}>Cart</button>
      <button onClick={handleShowOrderHistory}>Order History</button>
      <ul>
        {items.map((item) => (
          <li key={item.id} onClick={() => handleAddToCart(item.id)}>{item.name}</li>
        ))}
      </ul>
    </div>
  );
};

export default ListItems;