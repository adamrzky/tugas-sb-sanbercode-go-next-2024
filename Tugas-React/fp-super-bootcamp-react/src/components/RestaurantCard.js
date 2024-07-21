import React from 'react';
import './RestaurantCard.css'; // Styling khusus untuk RestaurantCard

function RestaurantCard({ restaurant }) {
  return (
    <div className="restaurant-card">
      <h2>{restaurant.Name}</h2>
      <p><strong>Alamat:</strong> {restaurant.Address}</p>

      <div className="foods">
        <h3>Daftar Makanan:</h3>
        <ul>
          {restaurant.Foods.map((food) => (
            <li key={food.ID}>
              <strong>{food.Name}</strong> - {food.Description} - Rp{food.Price}
            </li>
          ))}
        </ul>
      </div>

      <div className="reviews">
        <h3>Ulasan:</h3>
        <ul>
          {restaurant.Reviews.map((review) => (
            <li key={review.ID}>
              <strong>Rating:</strong> {review.Rating} - {review.Comment}
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default RestaurantCard;
