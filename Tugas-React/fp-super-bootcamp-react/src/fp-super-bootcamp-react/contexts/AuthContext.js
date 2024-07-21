// src/fp-super-bootcamp-react/contexts/AuthContext.js
import React, { createContext, useState } from 'react';
import axios from 'axios';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);

  const login = async (credentials) => {
    const response = await axios.post('/api/login', credentials);
    setUser(response.data.user);
    localStorage.setItem('token', response.data.token);
  };

  const logout = () => {
    setUser(null);
    localStorage.removeItem('token');
  };

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};
