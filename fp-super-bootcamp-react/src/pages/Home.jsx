// File: src/pages/Home.jsx
import React from 'react';
import Navbar from '../components/Navbar'; // Sesuaikan dengan path yang benar
import RestaurantList from '../components/RestaurantList'; // Sesuaikan dengan path yang benar
import './Home.css'; // Import CSS file for styling

const Home = () => {
    return (
        <div>
            <Navbar />
            <div className="content">
                <h1>Welcome to Culinary Reviews</h1>
                <RestaurantList />
            </div>
        </div>
    );
};

export default Home;
