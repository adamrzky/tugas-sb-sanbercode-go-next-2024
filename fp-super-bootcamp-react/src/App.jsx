// File: src/App.jsx
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar'; // Import Navbar from the components folder
import Home from './pages/Home'; // Define or import these components
// import Reviews from './pages/Reviews';
// import ManageUsers from './pages/ManageUsers';
// import Profile from './pages/Profile';
// import Login from './pages/Login'; 

const App = () => {
    return (
        <Router>
            <div>
                <Navbar />
                <Routes>
                    <Route path="/" element={<Home />} />
                    {/* <Route path="/reviews" element={<Reviews />} />
                    <Route path="/manage" element={<ManageUsers />} />
                    <Route path="/profile" element={<Profile />} />
                    <Route path="/login" element={<Login />} /> */}
                </Routes>
            </div>
        </Router>
    );
};

export default App;
