import { useState } from 'react';
import Navbar from '../components/Navbar';
import Swal from 'sweetalert2';
import Router from 'next/router';
import useAuthStore from '../store/authStore';

const Register = () => {
    const setUser = useAuthStore(state => state.setUser);
    const [user, setUserState] = useState({
        email: "",
        username: "",
        password: ""
    });

    const handleChange = (e) => {
        setUserState({ ...user, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(user)
            });
            const data = await response.json();
            if (response.ok) {
                setUser(data.user); // Set user in global state
                Swal.fire('Success', 'Registration successful!', 'success');
                Router.push('/login');
            } else {
                throw new Error(data.message || "Registration failed");
            }
        } catch (error) {
            Swal.fire('Error', error.message, 'error');
        }
    };

    return (
        <>
            <Navbar />
            <div className="container mx-auto mt-8">
                <h1 className="text-2xl font-bold mb-4">Register</h1>
                <form onSubmit={handleSubmit} className="flex flex-col gap-3">
                    <input
                        type="email"
                        name="email"
                        placeholder="Email"
                        value={user.email}
                        onChange={handleChange}
                        required
                        className="p-2 border border-gray-300 rounded"
                    />
                    <input
                        type="text"
                        name="username"
                        placeholder="Username"
                        value={user.username}
                        onChange={handleChange}
                        required
                        className="p-2 border border-gray-300 rounded"
                    />
                    <input
                        type="password"
                        name="password"
                        placeholder="Password"
                        value={user.password}
                        onChange={handleChange}
                        required
                        className="p-2 border border-gray-300 rounded"
                    />
                    <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded">
                        Register
                    </button>
                </form>
            </div>
        </>
    );
};

export default Register;
