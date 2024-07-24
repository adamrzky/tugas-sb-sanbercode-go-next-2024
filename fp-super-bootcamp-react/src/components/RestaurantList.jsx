import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Modal from 'react-modal';
import './RestaurantList.css';

Modal.setAppElement('#root');

const RestaurantList = () => {
    const [restaurants, setRestaurants] = useState([]);
    const [selectedRestaurant, setSelectedRestaurant] = useState(null);

    useEffect(() => {
        const fetchRestaurants = async () => {
            try {
                const response = await axios.get('http://localhost:8080/restaurants/');
                if (Array.isArray(response.data)) {
                    setRestaurants(response.data);
                } else {
                    setRestaurants([response.data]);
                }
            } catch (error) {
                console.error("Error fetching restaurants data:", error);
            }
        };

        fetchRestaurants();
    }, []);

    const openModal = (restaurant) => {
        setSelectedRestaurant(restaurant);
    };

    const closeModal = () => {
        setSelectedRestaurant(null);
    };

    return (
        <div className="restaurant-list">
            {restaurants.length > 0 ? (
                restaurants.map((restaurant) => (
                    <div key={restaurant.ID} className="restaurant-card">
                        <img 
                            src={`https://picsum.photos/400/300?random=${restaurant.ID}`} 
                            alt={restaurant.Name} 
                            className="restaurant-image" 
                        />
                        <div className="restaurant-details">
                            <h2 className="restaurant-name">{restaurant.Name}</h2>
                            <p className="restaurant-address">{restaurant.Address}</p>
                            <div className="restaurant-rating">
                                <span>{restaurant.Reviews.length > 0 ? (restaurant.Reviews.reduce((acc, review) => acc + review.Rating, 0) / restaurant.Reviews.length).toFixed(1) : 'No rating'}</span>
                            </div>
                            <button onClick={() => openModal(restaurant)} className="view-details-button">View Details</button>
                        </div>
                    </div>
                ))
            ) : (
                <p>No restaurants found.</p>
            )}

            {selectedRestaurant && (
                <Modal
                    isOpen={!!selectedRestaurant}
                    onRequestClose={closeModal}
                    contentLabel="Restaurant Details"
                    className="modal"
                    overlayClassName="overlay"
                >
                    <img 
                        src={`https://picsum.photos/600/400?random=${selectedRestaurant.ID}`} 
                        alt={selectedRestaurant.Name} 
                        className="modal-restaurant-image" 
                    />
                    <h2>{selectedRestaurant.Name}</h2>
                    <p>{selectedRestaurant.Address}</p>

                    <div className="modal-scrollable-content">
                        <div className="modal-section">
                            <h3>Menu:</h3>
                            {selectedRestaurant.Foods && selectedRestaurant.Foods.length > 0 ? (
                                <ul>
                                    {selectedRestaurant.Foods.map((food) => (
                                        <li key={food.ID} className="food-item">
                                            <img 
                                                src={`https://picsum.photos/100/100?random=${food.ID}`} 
                                                alt={food.Name} 
                                                className="food-image" 
                                            />
                                            <div>
                                                <h4>{food.Name}</h4>
                                                <p>{food.Description}</p>
                                                <p>Price: ${food.Price}</p>
                                            </div>
                                        </li>
                                    ))}
                                </ul>
                            ) : (
                                <p>No menu items available.</p>
                            )}
                        </div>

                        <div className="modal-section">
                            <h3>Reviews:</h3>
                            {selectedRestaurant.Reviews && selectedRestaurant.Reviews.length > 0 ? (
                                <ul>
                                    {selectedRestaurant.Reviews.map((review) => (
                                        <li key={review.ID} className="review-item">
                                            <p><strong>Rating:</strong> {review.Rating}/5</p>
                                            <p><strong>Comment:</strong> {review.Comment}</p>
                                        </li>
                                    ))}
                                </ul>
                            ) : (
                                <p>No reviews yet.</p>
                            )}
                        </div>
                    </div>

                    <button onClick={closeModal} className="close-button">Close</button>
                </Modal>
            )}
        </div>
    );
};

export default RestaurantList;
