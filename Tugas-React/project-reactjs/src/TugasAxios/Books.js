import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Books.css'; 

const baseURL = "https://backend-tugas-reactjs-mocha.vercel.app";

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

    useEffect(() => {
        autoLogin();
    }, []);

    const autoLogin = () => {
        const userData = { username: "adam2", password: "1234qwer" };
        axios.post(`${baseURL}/login`, userData)
            .then(response => {
                setToken(response.data.token);
                fetchBooks(response.data.token);
            })
            .catch(error => console.error('Gagal login:', error));
    };

    const fetchBooks = (authToken) => {
        axios.get(`${baseURL}/books`, {
            headers: {
                Authorization: `Bearer ${authToken}`
            }
        })
        .then(response => {
            setBooks(response.data.data);
        })
        .catch(error => console.error('Gagal mengambil data:', error));
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
        console.log("Data buku yang dikirim:", bookData);
    
        axios.post(`${baseURL}/books`, bookData, {
            headers: { Authorization: `Bearer ${token}` }
        })
        .then(response => {
            setBooks([...books, response.data]);
            setFormData({
                title: '',
                description: '',
                image_url: '',
                release_year: '',
                price: '',
                total_page: ''
            });
        })
        .catch(error => {
            console.error('Gagal menambah buku:', error);
            if (error.response) {
                console.error('Response error:', error.response.data);
            }
        });
    };
    
    

    const updateBook = (id, bookData) => {
        axios.patch(`${baseURL}/books/${id}`, bookData, {
            headers: { Authorization: `Bearer ${token}` }
        })
        .then(response => {
            const updatedBooks = books.map(book => book.id === id ? response.data : book);
            setBooks(updatedBooks);
            setEditing(false);
            setFormData({
                title: '',
                description: '',
                image_url: '',
                release_year: '',
                price: '',
                total_page: ''
            });
            setEditingId(null);
        })
        .catch(error => console.error('Gagal mengupdate buku:', error));
    };

    const deleteBook = (id) => {
        axios.delete(`${baseURL}/books/${id}`, {
            headers: { Authorization: `Bearer ${token}` }
        })
        .then(() => {
            setBooks(books.filter(book => book.id !== id));
        })
        .catch(error => console.error('Gagal menghapus buku:', error));
    };

    return (
        <div className="container">
            <h1>Books Manager</h1>
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
