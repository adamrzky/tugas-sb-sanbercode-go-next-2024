import React, { useState, useEffect, useContext } from "react";
import axios from "axios";
import Modal from "react-modal";
import "./RestaurantList.css";
import { UserContext } from "../contexts/UserContext";

Modal.setAppElement("#root");

// const RestaurantList = () => {
//   const [restaurants, setRestaurants] = useState([]);
//   const [selectedRestaurant, setSelectedRestaurant] = useState(null);
//   const [rating, setRating] = useState(5);
//   const [comment, setComment] = useState("");

const renderStars = (rating) => {
  let stars = [];
  for (let i = 1; i <= 5; i++) {
    if (i <= rating) {
      stars.push(
        <span key={i} className="star">
          &#9733;
        </span>
      ); // Full star
    } else {
      stars.push(
        <span key={i} className="star">
          &#9734;
        </span>
      ); // Empty star
    }
  }
  return stars;
};

const RestaurantList = () => {
  const [restaurants, setRestaurants] = useState([]);
  const [selectedRestaurant, setSelectedRestaurant] = useState(null);
  const [detailsModalIsOpen, setDetailsModalIsOpen] = useState(false);
  const [reviewModalIsOpen, setReviewModalIsOpen] = useState(false);
  const [rating, setRating] = useState(5);
  const [comment, setComment] = useState("");
  const { user } = useContext(UserContext);

  useEffect(() => {
    const fetchRestaurants = async () => {
      try {
        const response = await axios.get(
          "https://fp-super-bootcamp-go.vercel.app/restaurants/"
        );
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
    // if (!isLoggedIn) return; // Cek apakah user sudah login
    setSelectedRestaurant(restaurant);
    setReviewModalIsOpen(true);
  };

  const closeReviewModal = () => {
    setReviewModalIsOpen(false);
    setSelectedRestaurant(null);
    setRating(5); // Reset the rating
    setComment(""); // Clear the comment
  };

  const submitReview = async () => {
    // Cek apakah restaurant dan user dipilih
    if (!selectedRestaurant || !user) {
      console.error("Restaurant or user data is missing.");
      return;
    }

    // Mendapatkan token dari localStorage
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No token available. Please log in.");
      return;
    }

    // Payload untuk request
    const payload = {
      restaurant_id: selectedRestaurant.ID,
      rating,
      comment,
    };

    // Konfigurasi headers untuk Axios
    const config = {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    };

    try {
      // Mengirim POST request dengan Axios
      const response = await axios.post(
        "https://fp-super-bootcamp-go.vercel.app/reviews/",
        payload,
        config
      );
      console.log("Review submitted successfully:", response.data);
      closeReviewModal(); // Tutup modal setelah pengiriman berhasil
      // Opsi: Refresh review atau update UI disini
    } catch (error) {
      console.error(
        "Error submitting review:",
        error.response ? error.response.data : error.message
      );
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
                {restaurant.Reviews.length > 0
                  ? renderStars(
                      restaurant.Reviews.reduce(
                        (acc, review) => acc + review.Rating,
                        0
                      ) / restaurant.Reviews.length
                    )
                  : "No rating"}
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
              {selectedRestaurant.Foods &&
              selectedRestaurant.Foods.length > 0 ? (
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
          {user ? (
            <>
              <h3>Add Your Review</h3>
              <select
                value={rating}
                onChange={(e) => setRating(e.target.value)}
              >
                {[1, 2, 3, 4, 5].map((num) => (
                  <option key={num} value={num}>
                    {num} Stars
                  </option>
                ))}
              </select>
              <textarea
                placeholder="Write your comment here..."
                value={comment}
                onChange={(e) => setComment(e.target.value)}
              ></textarea>
              <button onClick={submitReview}>Submit Review</button>
            </>
          ) : (
            <p>Please login to write a review.</p>
          )}
          <button onClick={closeReviewModal} className="close-button">
            Close
          </button>
          <div className="modal-section">
            <h3>Reviews:</h3>
            {selectedRestaurant.Reviews &&
            selectedRestaurant.Reviews.length > 0 ? (
              <ul>
                {selectedRestaurant.Reviews.map((review) => (
                  <li key={review.ID} className="review-item">
                    <p>
                      <strong>Rating:</strong> {review.Rating}/5
                    </p>
                    <p>
                      <strong>Comment:</strong> {review.Comment}
                    </p>
                  </li>
                ))}
              </ul>
            ) : (
              <p>No reviews yet.</p>
            )}
          </div>
        </Modal>
      )}
    </div>
  );
};

export default RestaurantList;
