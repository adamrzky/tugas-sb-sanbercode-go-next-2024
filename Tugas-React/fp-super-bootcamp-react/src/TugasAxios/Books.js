import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Books.css'; // Mengimpor CSS untuk styling

const baseURL = process.env.REACT_APP_BASE_URL;

function Books() {
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
    const [loading, setLoading] = useState(false); // State untuk mengelola status loading
    const [loadingMessage, setLoadingMessage] = useState(''); // State untuk pesan loading

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
        // Konversi release_year ke integer jika masih string
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
                release_year: 0, 
                price: "",
                total_page: ""
            });
            fetchBooks();
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
            setFormData(response.data); // Update formData dengan data terbaru
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
        <div className="container">
            <h1>Books Manager</h1>
            {loading && <p className="loading-message">{loadingMessage}</p>}
            <form onSubmit={handleSubmit} className="form">
                <input type="text" name="title" value={formData.title} onChange={handleInputChange} placeholder="Title" required />
                <input type="text" name="description" value={formData.description} onChange={handleInputChange} placeholder="Description" required />
                <input type="url" name="image_url" value={formData.image_url} onChange={handleInputChange} placeholder="Image URL" required />
                <input type="number" name="release_year" value={formData.release_year} onChange={handleInputChange} placeholder="Release Year" required />
                <input type="text" name="price" value={formData.price} onChange={handleInputChange} placeholder="Price" required />
                <input type="number" name="total_page" value={formData.total_page} onChange={handleInputChange} placeholder="Total Pages" required />
                <button type="submit" className="btn-submit">{editing ? "Update" : "Add"}</button>
            </form>
            <div className="book-list">
                <table>
                    <thead>
                        <tr>
                            <th>No</th>
                            <th>Title</th>
                            <th>Description</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {books.map((book, index) => (
                            <tr key={book.id}>
                                <td>{index + 1}</td>
                                <td>{book.title}</td>
                                <td>{book.description}</td>
                                <td>
                                    <button onClick={() => { setEditing(true); setEditingId(book.id); setFormData(book); }} className="btn-edit">Edit</button>
                                    <button onClick={() => deleteBook(book.id)} className="btn-delete">Delete</button>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
}

export default Books;
