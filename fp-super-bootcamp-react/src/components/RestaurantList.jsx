import React, { useState, useEffect } from "react";
import axios from "axios";
import Modal from "react-modal";
import "./RestaurantList.css";

Modal.setAppElement("#root");

// const RestaurantList = () => {
//   const [restaurants, setRestaurants] = useState([]);
//   const [selectedRestaurant, setSelectedRestaurant] = useState(null);
//   const [rating, setRating] = useState(5);
//   const [comment, setComment] = useState("");

  const RestaurantList = () => {
    const [restaurants, setRestaurants] = useState([]);
    const [selectedRestaurant, setSelectedRestaurant] = useState(null);
    const [detailsModalIsOpen, setDetailsModalIsOpen] = useState(false);
    const [reviewModalIsOpen, setReviewModalIsOpen] = useState(false);
    const [rating, setRating] = useState(5);
    const [comment, setComment] = useState("");
  
    useEffect(() => {
      const fetchRestaurants = async () => {
        try {
          const response = await axios.get("http://localhost:8080/restaurants/");
          setRestaurants(response.data);
        } catch (error) {
          console.error("Error fetching restaurants data:", error);
        }
      };
  
      fetchRestaurants();
    }, []);
  
    const openDetailsModal = (restaurant) => {
      setSelectedRestaurant(restaurant);
      setDetailsModalIsOpen(true);
    };
  
    const closeDetailsModal = () => {
      setDetailsModalIsOpen(false);
      setSelectedRestaurant(null);
    };
  
    const openReviewModal = (restaurant) => {
      setSelectedRestaurant(restaurant);
      setReviewModalIsOpen(true);
    };
  
    const closeReviewModal = () => {
      setReviewModalIsOpen(false);
      setSelectedRestaurant(null);
      setRating(5); // Reset the rating
      setComment(''); // Clear the comment
    };
  
    const submitReview = async () => {
      if (!selectedRestaurant) return;
  
      const payload = {
        restaurant_id: selectedRestaurant.ID,
        rating,
        comment
      };
  
      try {
        await axios.post('http://localhost:8080/reviews', payload);
        closeReviewModal(); // Close modal after submission
      } catch (error) {
        console.error('Error submitting review:', error);
      }
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
                <span>
                  {restaurant.Reviews.length > 0
                    ? (
                        restaurant.Reviews.reduce(
                          (acc, review) => acc + review.Rating,
                          0
                        ) / restaurant.Reviews.length
                      ).toFixed(1)
                    : "No rating"}
                </span>
              </div>
              <button
                onClick={() => openDetailsModal(restaurant)}
                className="view-details-button"
              >
                View Details
              </button>
              <button
                onClick={() => openReviewModal(restaurant)}
                className="review-button"
              >
                Review
              </button>
            </div>
          </div>
        ))
      ) : (
        <p>No restaurants found.</p>
      )}

      {selectedRestaurant && (
        <Modal
          isOpen={detailsModalIsOpen}
          onRequestClose={closeDetailsModal}
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

          <button onClick={closeDetailsModal} className="close-button">
            Close
          </button>
        </Modal>
      )}

      {/* Review Modal */}
      {selectedRestaurant && (
        <Modal
          isOpen={reviewModalIsOpen}
          onRequestClose={closeReviewModal}
          contentLabel="Add Review"
          className="modal"
          overlayClassName="overlay"
        >
          <h2>Write a Review for {selectedRestaurant.Name}</h2>
          <select value={rating} onChange={(e) => setRating(e.target.value)}>
            {[1, 2, 3, 4, 5].map(num => (
              <option key={num} value={num}>{num} Star{num > 1 ? 's' : ''}</option>
            ))}
          </select>
          <textarea
            placeholder="Write your comment here..."
            value={comment}
            onChange={(e) => setComment(e.target.value)}
          ></textarea>
          <button onClick={submitReview}>Submit Review</button>
          <button onClick={closeReviewModal} className="close-button">Close</button>
        </Modal>
      )}
    </div>
  );

};

export default RestaurantList;
