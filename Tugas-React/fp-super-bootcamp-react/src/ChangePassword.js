import React, { useState } from 'react';
import axios from 'axios';

function ChangePassword() {
  const [oldPassword, setOldPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');

  const handleChangePassword = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/auth/change-password', {
        old_password: oldPassword,
        new_password: newPassword
      });
      alert('Password berhasil diubah!');
    } catch (error) {
      console.error('Gagal mengubah password:', error);
      alert('Gagal mengubah password!');
    }
  };

  return (
    <div style={{ padding: 20 }}>
      <h2>Ganti Password</h2>
      <form onSubmit={handleChangePassword}>
        <div>
          <label>Password Lama:</label>
          <input type="password" value={oldPassword} onChange={(e) => setOldPassword(e.target.value)} required />
        </div>
        <div>
          <label>Password Baru:</label>
          <input type="password" value={newPassword} onChange={(e) => setNewPassword(e.target.value)} required />
        </div>
        <button type="submit">Ganti Password</button>
      </form>
    </div>
  );
}

export default ChangePassword;
