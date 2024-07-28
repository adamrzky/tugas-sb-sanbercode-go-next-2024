// src/components/Navbar.jsx
import React, { useContext } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { FaUserCircle } from 'react-icons/fa';
import './Navbar.css';
import { UserContext } from '../contexts/UserContext';

const Navbar = () => {
    const { user, logout } = useContext(UserContext);
    const navigate = useNavigate();  // Pastikan useNavigate diinisialisasi di sini

    const handleLogout = () => {
        logout(() => navigate("/login"));  // Gunakan arrow function untuk memastikan bahwa ini adalah fungsi
    };

    return (
        <nav className="navbar">
            <div className="navbar-container">
                <Link to="/" className="navbar-logo">Culinary Reviews</Link>
                <div className="menu">
                    <Link to="/" className="nav-item">Home</Link>
                    {/* <Link to="/reviews" className="nav-item">Reviews</Link> */}
                    {/* <Link to="/manage" className="nav-item">Manage Users</Link> */}
                    {user ? (
                        <div className="dropdown">
                            <div className="user-info">
                                <FaUserCircle className="user-icon" />
                                <span className="user-email">{user.email}</span>
                            </div>
                            <div className="dropdown-content">
                                <Link to="/profile">Profile</Link>
                                <Link to="/change-password">Management User</Link>
                                <button onClick={handleLogout} className="button-logout">Logout</button>
                            </div>
                        </div>
                    ) : (
                        <Link to="/login" className="nav-item">Login</Link>
                    )}
                </div>
            </div>
        </nav>
    );
};

export default Navbar;
