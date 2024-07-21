import React from 'react';
import { Link } from 'react-router-dom';
import './Navbar.css'; // Tambahkan ini

function Navbar() {
  return (
    <nav className="navbar">
      <ul>
        <li><Link to="/">Home</Link></li>
        <li><Link to="/login">Login</Link></li>
        <li><Link to="/profile">Profile</Link></li>
        <li><Link to="/reviews">Reviews</Link></li>
      </ul>
    </nav>
  );
}

export default Navbar;
