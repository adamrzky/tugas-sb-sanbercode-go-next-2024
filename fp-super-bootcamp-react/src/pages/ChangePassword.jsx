import React, { useState, useContext } from 'react';
import axios from 'axios';
import { UserContext } from '../contexts/UserContext'; // Pastikan path benar

const ChangePassword = () => {
    const [oldPassword, setOldPassword] = useState('');
    const [newPassword, setNewPassword] = useState('');
    const [message, setMessage] = useState('');
    const { user } = useContext(UserContext);

    const handleChangePassword = async (e) => {
        e.preventDefault();
        if (!oldPassword || !newPassword) {
            setMessage("Please fill in all fields.");
            return;
        }
        try {
            const response = await axios.post('http://localhost:8080/auth/change-password', {
                old_password: oldPassword,
                new_password: newPassword
            }, {
                headers: {
                    Authorization: `Bearer ${user.token}`  // Pastikan token ini valid
                }
            });
            if (response.data.code === 200) {
                setMessage('Password successfully changed.');
            } else {
                setMessage(response.data.message || 'Failed to change password.');  // Gunakan pesan dari server jika tersedia
            }
        } catch (error) {
            setMessage('Error changing password. ' + (error.response?.data?.message || error.message));
            console.error(error);
        }
    };
    

    return (
        <div className="login-container"> {/* Using the same style as login for consistency */}
            <form onSubmit={handleChangePassword} className="login-form">
                <h2>Change Password</h2>
                <input
                    type="password"
                    placeholder="Old Password"
                    value={oldPassword}
                    onChange={e => setOldPassword(e.target.value)}
                    required
                />
                <input
                    type="password"
                    placeholder="New Password"
                    value={newPassword}
                    onChange={e => setNewPassword(e.target.value)}
                    required
                />
                <button type="submit">Change Password</button>
                <div>{message}</div>
            </form>
        </div>
    );
};

export default ChangePassword;
