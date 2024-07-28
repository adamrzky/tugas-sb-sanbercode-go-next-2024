import React, { useState, useEffect, useContext } from "react";
import axios from "axios";
import Modal from "react-modal";
import "./RestaurantList.css";
import { UserContext } from "../contexts/UserContext";
import Swal from "sweetalert2";

Modal.setAppElement("#root");

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
  const [hasReviewed, setHasReviewed] = useState(false);

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
    setSelectedRestaurant(restaurant);
    setReviewModalIsOpen(true);
    const alreadyReviewed = restaurant.Reviews.some(
      (review) => review.Email === user?.email
    );
    setHasReviewed(alreadyReviewed);
  };

  const closeReviewModal = () => {
    setReviewModalIsOpen(false);
    setSelectedRestaurant(null);
    setRating(5); // Reset the rating
    setComment(""); // Clear the comment
  };

  // Function to check if the user has already reviewed the restaurant
  const checkIfReviewed = (restaurant) => {
    const hasReviewed = restaurant.Reviews.some(
      (review) => review.Email === user.email
    );
    setHasReviewed(hasReviewed);
  };

  const StarRating = ({ rating, setRating, editable = true }) => {
    const handleRating = (rate) => {
      if (editable && setRating) {
        setRating(rate);
      }
    };

    return (
      <div className="star-rating">
        {[...Array(5)].map((_, index) => {
          const ratingValue = index + 1;
          return (
            <span
              key={ratingValue}
              className={`star ${ratingValue <= rating ? "filled" : ""}`}
              onClick={() => handleRating(ratingValue)}
              style={{ cursor: editable ? "pointer" : "default" }}
            >
              {ratingValue <= rating ? "★" : "☆"}
            </span>
          );
        })}
      </div>
    );
  };

  const submitReview = async () => {
    if (!selectedRestaurant || !user) {
      Swal.fire({
        icon: "error",
        title: "Oops...",
        text: "Restaurant or user data is missing.",
      });
      return;
    }

    const token = localStorage.getItem("token");
    if (!token) {
      Swal.fire({
        icon: "error",
        title: "Oops...",
        text: "No token available. Please log in.",
      });
      return;
    }

    const payload = {
      restaurant_id: selectedRestaurant.ID,
      rating: parseInt(rating, 10),
      comment,
    };

    const config = {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    };

    try {
      const response = await axios.post(
        "https://fp-super-bootcamp-go.vercel.app/reviews/",
        payload,
        config
      );
      Swal.fire({
        icon: "success",
        title: "Success!",
        text: "Review submitted successfully.",
      });
      const updatedRestaurants = restaurants.map((restaurant) =>
        restaurant.ID === selectedRestaurant.ID
          ? { ...restaurant, Reviews: [...restaurant.Reviews, response.data] } // Assuming the server responds with the new review
          : restaurant
      );
      setRestaurants(updatedRestaurants);
      setHasReviewed(true);
      closeReviewModal();
    } catch (error) {
      const errorMessage =
        error.response && error.response.data && error.response.data.error
          ? error.response.data.error
          : "An unexpected error occurred.";
      Swal.fire({
        icon: "error",
        title: "Failed to submit review",
        text: errorMessage,
      });
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
                Menu
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
        <p>Loading Data....</p>
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
          <div className="modal-header">
            <h2>{selectedRestaurant.Name}</h2>
            {/* <h5>{selectedRestaurant.Address}</h5> */}
          </div>
          <div className="modal-body">
            {user ? (
              !hasReviewed ? (
                <div className="review-form">
                  {/* <h4>Add Your Review</h4> */}
                  <div className="star-rating">
                    {/* <p className="rating-instruction">Pilih bintang:</p> */}
                    {[...Array(5)].map((_, index) => (
                      <span
                        key={index}
                        className={`star ${index < rating ? "filled" : ""}`}
                        onClick={() => setRating(index + 1)}
                      >
                        &#9733;
                      </span>
                    ))}
                  </div>
                  <textarea
                    className="review-textarea"
                    style={{ width: "100%", minHeight: "100px" }} // Adjust the size of textarea
                    placeholder="Write your comment here..."
                    value={comment}
                    onChange={(e) => setComment(e.target.value)}
                  ></textarea>
                  <br></br>
                  <button className="submit-button" onClick={submitReview}>
                    Submit Review
                  </button>
                </div>
              ) : (
                <p className="already-reviewed">
                  You have already reviewed this restaurant.
                </p>
              )
            ) : (
              <p className="login-prompt">Please login to write a review.</p>
            )}
          </div>

          <div className="modal-footer">
            <button className="close-button" onClick={closeReviewModal}>
              Close
            </button>
          </div>
          <div className="modal-section">
            <h3>Reviews:</h3>
            {selectedRestaurant.Reviews &&
            selectedRestaurant.Reviews.length > 0 ? (
              <div
                className="review-list"
                style={{ maxHeight: "300px", overflowY: "auto" }}
              >
                <ul>
                  {selectedRestaurant.Reviews.map((review) => (
                    <li key={review.ID} className="review-item">
                      <p>
                        <strong>Rating:</strong>{" "}
                        <StarRating rating={review.Rating} editable={false} />
                      </p>
                      <div className="review-content">
                        <p>
                          <strong>Comment:</strong> {review.Comment}
                        </p>
                        <p className="review-email">
                          <strong>Email:</strong> {review.Email}
                        </p>
                      </div>
                    </li>
                  ))}
                </ul>
              </div>
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
