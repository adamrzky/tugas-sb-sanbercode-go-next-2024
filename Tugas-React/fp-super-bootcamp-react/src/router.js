// src/Router.js
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import HomePage from './pages/HomePage';
import RestaurantDetail from './pages/RestaurantDetail';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import ProfilePage from './pages/ProfilePage';
import ManageUsersPage from './pages/ManageUsersPage';
import ProtectedRoute from './components/ProtectedRoute'; // For routes that require login

const RouterConfig = () => (
  <Router>
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/restaurant/:id" element={<RestaurantDetail />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/profile" element={<ProtectedRoute component={ProfilePage} />} />
      <Route path="/manage-users" element={<ProtectedRoute component={ManageUsersPage} />} />
    </Routes>
  </Router>
);

export default RouterConfig;
