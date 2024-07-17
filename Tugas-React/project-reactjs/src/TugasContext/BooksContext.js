import React, { createContext, useState, useEffect } from 'react';
import axios from 'axios';

const baseURL = process.env.REACT_APP_BASE_URL;
const BooksContext = createContext();

export const BooksProvider = ({ children }) => {
    const [books, setBooks] = useState([]);
    const [token, setToken] = useState('');
    const [formData, setFormData] = useState({
        title: '',
        description: '',
        image_url: '',
        release_year: '',
        price: '',
        total_page: ''
    });
    const [editing, setEditing] = useState(false);
    const [editingId, setEditingId] = useState(null);
    const [loading, setLoading] = useState(false);
    const [loadingMessage, setLoadingMessage] = useState('');

    useEffect(() => {
        autoLogin();
    }, []);

    const autoLogin = () => {
        setLoading(true);
        setLoadingMessage('Generate JWT...');
        const userData = { username: "adam2", password: "1234qwer" };
        axios.post(`${baseURL}/login`, userData)
            .then(response => {
                setToken(response.data.token);
                fetchBooks(response.data.token);
            })
            .catch(error => console.error('Gagal login:', error))
            .finally(() => {
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
        .catch(error => console.error('Gagal mengambil data:', error))
        .finally(() => {
            setLoading(false);
            setLoadingMessage('');
        });
    };

    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setFormData(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        const currentYear = new Date().getFullYear();
        const payload = {
            ...formData,
            release_year: parseInt(formData.release_year, 10),
            price: formData.price,
            total_page: parseInt(formData.total_page, 10)
        };

        if (payload.release_year < 0 || payload.release_year > currentYear) {
            alert(`Release Year harus antara 0 dan ${currentYear}`);
            return;
        }

        if (editing) {
            updateBook(editingId, payload);
        } else {
            addBook(payload);
        }
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
            setFormData({
                title: "",
                description: "",
                image_url: "",
                release_year: '',
                price: "",
                total_page: ""
            });
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
            setFormData({
                title: '',
                description: '',
                image_url: '',
                release_year: '',
                price: '',
                total_page: ''
            });
            setEditing(false);
            setEditingId(null);
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
        <BooksContext.Provider value={{
            books, formData, setFormData, handleInputChange, handleSubmit, loading, loadingMessage,
            editing, setEditing, setEditingId, deleteBook
        }}>
            {children}
        </BooksContext.Provider>
    );
};

export default BooksContext;
