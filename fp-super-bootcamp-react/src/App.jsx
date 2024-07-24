// src/App.jsx
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { UserProvider } from './contexts/UserContext';
import Navbar from './components/Navbar';
import Home from './pages/Home';
import Login from './pages/Login';
import ChangePassword from './pages/ChangePassword'; // Import ChangePassword


function App() {
  return (
    <UserProvider>
      <Router>
        <Navbar />
        <Routes> {/* Mengganti Switch dengan Routes */}
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<Home />} />
          <Route path="/change-password" element={<ChangePassword />} /> {/* Route for Change Password */}
          </Routes>
      </Router>
    </UserProvider>
  );
}

export default App;
