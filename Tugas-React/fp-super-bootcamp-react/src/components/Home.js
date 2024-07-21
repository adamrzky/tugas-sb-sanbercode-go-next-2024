import React, { useEffect, useState } from 'react';
import axios from 'axios';
import RestaurantCard from './RestaurantCard';
import './App.css'; // Untuk styling umum

function Home() {
  const [restaurants, setRestaurants] = useState([]);

  useEffect(() => {
    const fetchRestaurants = async () => {
      try {
        const response = await axios.get('http://localhost:8080/restaurants');
        setRestaurants(response.data);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchRestaurants();
  }, []);

  return (
    <div className="container">
      <h1>Daftar Restoran</h1>
      <div className="restaurants-list">
        {restaurants.map((restaurant) => (
          <RestaurantCard key={restaurant.ID} restaurant={restaurant} />
        ))}
      </div>
    </div>
  );
}

export default Home;
