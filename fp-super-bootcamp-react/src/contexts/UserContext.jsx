// src/contexts/UserContext.jsx
import React, { createContext, useState, useEffect } from 'react';

export const UserContext = createContext({
  user: null,
  setUser: () => {},
  logout: () => {}
});

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem('token');
    const email = localStorage.getItem('email');
    if (token && email) {
      setUser({ email, token });
    }
  }, []);

  const saveUser = (userData) => {
    localStorage.setItem('token', userData.token);
    localStorage.setItem('email', userData.email);
    setUser(userData);
  };

  const logout = (navigateCallback) => {
    localStorage.removeItem('token');
    localStorage.removeItem('email');
    setUser(null);
    if (typeof navigateCallback === 'function') {
      navigateCallback();  // Memanggil navigateCallback jika itu fungsi
    } else {
      console.error('Navigate callback is not a function');  // Menambahkan pesan error jika callback bukan fungsi
    }
  };
  return (
    <UserContext.Provider value={{ user, setUser: saveUser, logout }}>
      {children}
    </UserContext.Provider>
  );
};
