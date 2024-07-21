import React, { useState } from 'react';
import axios from 'axios';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/auth/login', { email, password });
      console.log('Token JWT:', response.data.token);
      alert('Login berhasil!');
    } catch (error) {
      console.error('Login gagal:', error);
      alert('Login gagal!');
    }
  };

  return (
    <div className="container">
      <h2>Login</h2>
      <form onSubmit={handleLogin}>
        <div className="form-group">
          <label className="label">Email:</label>
          <input type="email" className="input" value={email} onChange={(e) => setEmail(e.target.value)} required />
        </div>
        <div className="form-group">
          <label className="label">Password:</label>
          <input type="password" className="input" value={password} onChange={(e) => setPassword(e.target.value)} required />
        </div>
        <button type="submit" className="button">Masuk</button>
      </form>
    </div>
  );
}

export default Login;
