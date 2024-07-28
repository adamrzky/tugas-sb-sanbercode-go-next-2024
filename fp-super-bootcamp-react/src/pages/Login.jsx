import React, { useState, useContext } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../contexts/UserContext';
import './Login.css'; // Pastikan ini benar
import Swal from 'sweetalert2';

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [fullName, setFullName] = useState('');
    const [bio, setBio] = useState('');
    const [isRegistering, setIsRegistering] = useState(false);
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
                navigate('/'); // Redirect ke home setelah login
                Swal.fire('Login Berhasil', 'Anda berhasil login.', 'success');
            } else {
                Swal.fire('Login Gagal', 'Periksa kredensial Anda dan coba lagi.', 'error');
            }
        } catch (error) {
            Swal.fire('Error', 'Terjadi kesalahan saat login. Coba lagi nanti.', 'error');
        }
    };

    const handleRegister = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('https://fp-super-bootcamp-go.vercel.app/auth/register', {
                email,
                password,
                profile: {
                    full_name: fullName,
                    bio
                }
            });
            if (response.data.message === "Registration successful") {
                Swal.fire('Registrasi Berhasil', 'Silakan login dengan akun baru Anda.', 'success');
                setIsRegistering(false);
            } else {
                Swal.fire('Registrasi Gagal', response.data.error || 'Coba lagi.', 'error');
            }
        } catch (error) {
            Swal.fire('Error', 'Terjadi kesalahan saat registrasi. Coba lagi nanti.', 'error');
        }
    };

    return (
        <div className="login-container">
            <form onSubmit={isRegistering ? handleRegister : handleLogin} className="login-form">
                <h2>{isRegistering ? 'Register' : 'Login'}</h2>
                <input type="email" value={email} onChange={e => setEmail(e.target.value)} placeholder="Email"/>
                <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Password"/>
                {isRegistering && (
                    <>
                        <input type="text" value={fullName} onChange={e => setFullName(e.target.value)} placeholder="Full Name" />
                        <textarea value={bio} onChange={e => setBio(e.target.value)} placeholder="Bio" />
                    </>
                )}
                <button type="submit">{isRegistering ? 'Register' : 'Login'}</button>
                <button type="button" onClick={() => setIsRegistering(!isRegistering)} className="toggle-button">
                    {isRegistering ? 'Already have an account? Login' : 'No account? Register'}
                </button>
            </form>
        </div>
    );
};

export default Login;
