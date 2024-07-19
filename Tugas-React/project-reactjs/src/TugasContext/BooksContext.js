// src/TugasContext/BooksContext.js
import React, { createContext, useState, useEffect } from 'react';
import axios from 'axios';

const baseURL = process.env.REACT_APP_BASE_URL;
const BooksContext = createContext();

export const BooksProvider = ({ children }) => {
    const [books, setBooks] = useState([]);
    const [token, setToken] = useState('');
    const [loading, setLoading] = useState(false);
    const [loadingMessage, setLoadingMessage] = useState('');

    useEffect(() => {
        if (!token) {
            autoLogin();
        } else {
            fetchBooks(token);
        }
    }, [token]);

    const autoLogin = () => {
        setLoading(true);
        setLoadingMessage('Generate JWT...');
        const userData = { username: "adam2", password: "1234qwer" };
        axios.post(`${baseURL}/login`, userData)
            .then(response => {
                setToken(response.data.token);
            })
            .catch(error => {
                console.error('Gagal login:', error);
                setLoading(false);
                setLoadingMessage('');
            });
    };

    const fetchBooks = (authToken) => {
        setLoading(true);
        setLoadingMessage('Fetching data...');
        axios.get(`${baseURL}/books`, {
            headers: {
                Authorization: `Bearer ${authToken}`
            }
        })
        .then(response => {
            setBooks(response.data.data);
        })
        .catch(error => {
            console.error('Gagal mengambil data:', error);
        })
        .finally(() => {
            setLoading(false);
            setLoadingMessage('');
        });
    };

    const addBook = (bookData) => {
        const payload = {
            ...bookData,
            release_year: parseInt(bookData.release_year, 10)
        };

        axios.post(`${baseURL}/books`, payload, {
            headers: { Authorization: `Bearer ${token}` }
        })
        .then(response => {
            if (response.data && response.data.id) {
                setBooks(prevBooks => [...prevBooks, response.data]);
            } else {
                console.error("Unexpected response structure:", response);
            }
        })
        .catch(error => {
            console.error("Failed to add book:", error);
        });
    };

    const updateBook = (id, bookData) => {
        setLoading(true);
        setLoadingMessage('Updating book...');
        axios.patch(`${baseURL}/books/${id}`, bookData, {
            headers: { Authorization: `Bearer ${token}` }
        })
        .then(response => {
            const updatedBooks = books.map(book => book.id === id ? response.data : book);
            setBooks(updatedBooks);
        })
        .catch(error => console.error('Gagal mengupdate buku:', error))
        .finally(() => {
            setLoading(false);
            setLoadingMessage('');
        });
    };

    const deleteBook = (id) => {
        setLoading(true);
        setLoadingMessage('Deleting book...');
        axios.delete(`${baseURL}/books/${id}`, {
            headers: { Authorization: `Bearer ${token}` }
        })
        .then(() => {
            setBooks(books.filter(book => book.id !== id));
        })
        .catch(error => console.error('Gagal menghapus buku:', error))
        .finally(() => {
            setLoading(false);
            setLoadingMessage('');
        });
    };

    return (
        <BooksContext.Provider value={{ books, addBook, updateBook, deleteBook, fetchBooks, loading, loadingMessage }}>
            {children}
        </BooksContext.Provider>
    );
};

export { BooksContext };
