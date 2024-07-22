// File: src/components/Navbar.jsx
import React from 'react';
import { Link } from 'react-router-dom';
import { FaUserCircle } from 'react-icons/fa'; // Import the user icon
import './Navbar.css'; // Import your CSS file

const Navbar = () => {
    return (
        <nav className="navbar">
            <div className="navbar-container">
                <Link to="/" className="navbar-logo">
                    Culinary Reviews
                </Link>
                <div className="menu">
                    <Link to="/" className="nav-item">Home</Link>
                    <Link to="/reviews" className="nav-item">Reviews</Link>
                    <Link to="/manage" className="nav-item">Manage Users</Link>
                    <Link to="/login" className="nav-item">Login</Link>
                    <div className="dropdown">
                        <FaUserCircle className="user-icon"/>
                        <div className="dropdown-content">
                            <Link to="/profile">Profile</Link>
                            <Link to="/change-password">Change Password</Link>
                        </div>
                    </div>
                </div>
            </div>
        </nav>
    );
};

export default Navbar;
