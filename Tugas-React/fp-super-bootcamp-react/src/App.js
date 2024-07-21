import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Navbar from './components/Navbar';
import Home from './Home';
import Login from './Login';
import './App.css';

function App() {
  return (
    <Router>
      <Navbar />
      <div className="app-content">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          {/* Tambahkan rute lainnya jika ada */}
        </Routes>
      </div>
    </Router>
  );
}

export default App;
