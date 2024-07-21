import React, { useState } from 'react';
import axios from 'axios';

function Register() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [fullName, setFullName] = useState('');
  const [bio, setBio] = useState('');

  const handleRegister = async (e) => {
    e.preventDefault();
    const body = {
      email, 
      password, 
      profile: { full_name: fullName, bio }
    };
    try {
      const response = await axios.post('http://localhost:8080/auth/register', body);
      alert(response.data.message);
    } catch (error) {
      console.error('Registrasi gagal:', error);
      alert('Registrasi gagal!');
    }
  };

  return (
    <div style={{ padding: 20 }}>
      <h2>Register</h2>
      <form onSubmit={handleRegister}>
        <div>
          <label>Email:</label>
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
        </div>
        <div>
          <label>Password:</label>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
        </div>
        <div>
          <label>Nama Lengkap:</label>
          <input type="text" value={fullName} onChange={(e) => setFullName(e.target.value)} required />
        </div>
        <div>
          <label>Bio:</label>
          <input type="text" value={bio} onChange={(e) => setBio(e.target.value)} />
        </div>
        <button type="submit">Daftar</button>
      </form>
    </div>
  );
}

export default Register;
