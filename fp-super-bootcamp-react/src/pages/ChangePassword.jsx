// ChangePassword.jsx
import React, { useState, useContext } from 'react';
import axios from 'axios';
import { UserContext } from '../contexts/UserContext';

const ChangePassword = () => {
    const [oldPassword, setOldPassword] = useState('');
    const [newPassword, setNewPassword] = useState('');
    const [message, setMessage] = useState('');
    const { user } = useContext(UserContext);

    const handleChangePassword = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8080/auth/change-password', {
                old_password: oldPassword,
                new_password: newPassword
            }, {
                headers: { Authorization: `Bearer ${user.token}` }
            });
            setMessage(response.data.message);
        } catch (error) {
            setMessage('Failed to change password. Please try again.');
        }
    };

    return (
        <div>
            <h2>Change Password</h2>
            <form onSubmit={handleChangePassword}>
                <input type="password" value={oldPassword} onChange={e => setOldPassword(e.target.value)} placeholder="Old Password" required />
                <input type="password" value={newPassword} onChange={e => setNewPassword(e.target.value)} placeholder="New Password" required />
                <button type="submit">Submit</button>
            </form>
            {message && <p>{message}</p>}
        </div>
    );
};

export default ChangePassword;
