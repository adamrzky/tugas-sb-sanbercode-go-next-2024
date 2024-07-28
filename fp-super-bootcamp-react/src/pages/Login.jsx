import React, { useState, useContext } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../contexts/UserContext';
import './Login.css';
import Swal from 'sweetalert2'; // Import SweetAlert2

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const { setUser } = useContext(UserContext);
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('https://fp-super-bootcamp-go.vercel.app/auth/login', {
                email,
                password
            });
            if (response.data.code === 200) {
                setUser({ email, token: response.data.token });
                Swal.fire({ // Menampilkan SweetAlert sukses
                    icon: 'success',
                    title: 'Logged in!',
                    text: 'You have successfully logged in.',
                }).then(() => {
                    navigate('/'); // Redirect ke home setelah mengklik OK
                });
            } else {
                Swal.fire({ // Menampilkan SweetAlert error
                    icon: 'error',
                    title: 'Login Failed',
                    text: 'Please check your credentials.',
                });
            }
        } catch (error) {
            Swal.fire({ // Menampilkan SweetAlert untuk error teknis
                icon: 'error',
                title: 'Login Error',
                text: 'Login error. Please try again later.',
            });
            console.error('Login request failed:', error);
        }
    };

    return (
        <div className="login-container">
            <form onSubmit={handleLogin} className="login-form">
                <h2>Login</h2>
                <input type="email" value={email} onChange={e => setEmail(e.target.value)} placeholder="Email"/>
                <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Password"/>
                <button type="submit">Login</button>
            </form>
        </div>
    );
};

export default Login;
