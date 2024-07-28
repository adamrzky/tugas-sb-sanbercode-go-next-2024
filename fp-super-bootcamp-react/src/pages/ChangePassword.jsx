// ChangePassword.jsx
import React, { useState, useContext } from 'react';
import axios from 'axios';
import { UserContext } from '../contexts/UserContext';
import './Login.css';  // Menggunakan CSS yang sama untuk konsistensi

const ChangePassword = () => {
    const [oldPassword, setOldPassword] = useState('');
    const [newPassword, setNewPassword] = useState('');
    const [message, setMessage] = useState('');
    const { user } = useContext(UserContext);

    const handleChangePassword = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('https://fp-super-bootcamp-go.vercel.app/auth/change-password', {
                old_password: oldPassword,
                new_password: newPassword
            }, {
                headers: { Authorization: `Bearer ${user.token}` }
            });
            if (response.data.message) {
                setMessage(response.data.message);
            }
        } catch (error) {
            setMessage('Failed to change password. Please try again.');
        }
    };

    return (
        <div className="login-container">  
            <form onSubmit={handleChangePassword} className="login-form">
                <h2>Change Password</h2>
                <input type="password" value={oldPassword} onChange={e => setOldPassword(e.target.value)} placeholder="Old Password" required />
                <input type="password" value={newPassword} onChange={e => setNewPassword(e.target.value)} placeholder="New Password" required />
                <button type="submit">Submit</button>
                {message && <p>{message}</p>}
            </form>
        </div>
    );
};

export default ChangePassword;
