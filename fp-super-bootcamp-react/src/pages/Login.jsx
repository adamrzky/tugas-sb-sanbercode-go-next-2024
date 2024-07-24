// src/pages/Login.jsx
import React, { useState, useContext } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../contexts/UserContext';
import './Login.css'; // Pastikan ini benar

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const { setUser } = useContext(UserContext);
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8080/auth/login', {
                email,
                password
            });
            if (response.data.code === 200) {
                setUser({ email, token: response.data.token });
                navigate('/'); // Redirect to home
            } else {
                setError('Login failed. Please check your credentials.');
            }
        } catch (error) {
            setError('Login error. Please try again later.');
            console.error(error);
        }
    };

    return (
        <div className="login-container">
            <form onSubmit={handleLogin} className="login-form">
                <h2>Login</h2>
                <input type="email" value={email} onChange={e => setEmail(e.target.value)} placeholder="Email"/>
                <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Password"/>
                <button type="submit">Login</button>
                {error && <p>{error}</p>}
            </form>
        </div>
    );
};

export default Login;
