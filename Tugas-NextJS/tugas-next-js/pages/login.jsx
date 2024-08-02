// pages/login.jsx
import { useState } from 'react';
import useAuthStore from '../store/authStore'; // Pastikan path ini benar
import Swal from 'sweetalert2';
import Navbar from '../components/Navbar';

const LoginPage = () => {
  const { login } = useAuthStore();
  const [Username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ Username, password })
      });

      if (response.ok) {
        const data = await response.json();
        login(data.user, data.token); // Menyimpan user dan token ke store
        Swal.fire("Success", "Login successful!", "success");
        // Redirect atau update UI di sini
      } else {
        const errorData = await response.json();
        throw new Error(`Failed to login: ${errorData.message}`);
      }
    } catch (error) {
      console.error('Error during login:', error);
      Swal.fire("Error", error.message, "error");
    }
  };

  return (
    <>
      <Navbar />
      <div className="container mx-auto mt-8">
        <h1 className="text-2xl font-bold mb-4">Login</h1>
        <form onSubmit={handleSubmit} className="flex flex-col gap-3">
          <input
            type="text"
            name="Username"
            placeholder="Username"
            value={Username}
            onChange={(e) => setUsername(e.target.value)}
            required
            className="p-2 border border-gray-300 rounded"
          />
          <input
            type="password"
            name="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            className="p-2 border border-gray-300 rounded"
          />
          <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded">
            Login
          </button>
        </form>
      </div>
    </>
  );
};

export default LoginPage;
