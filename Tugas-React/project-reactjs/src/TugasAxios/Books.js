import React, { useState, useEffect } from 'react';
import axios from 'axios';

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
                console.log('Login successful:', response.data);
                setToken(response.data.token);
                fetchBooks(response.data.token);
            })
            .catch(error => console.error('Failed to login:', error));
    };

    const fetchBooks = (authToken) => {
      axios.get(`${baseURL}/books`, {
          headers: {
              Authorization: `Bearer ${authToken}`
          }
      })
      .then(response => {
          console.log('Data buku berhasil diambil:', response.data);
          setBooks(response.data.data); // Menggunakan data yang benar dari respons
      })
      .catch(error => console.error('Gagal mengambil data:', error));
  };

    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setFormData(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        if (editing) {
            updateBook(editingId, formData);
        } else {
            addBook(formData);
        }
    };

    const addBook = (bookData) => {
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
        .catch(error => console.error('Failed to add book:', error));
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
        .catch(error => console.error('Failed to update book:', error));
    };

    const deleteBook = (id) => {
        axios.delete(`${baseURL}/books/${id}`, {
            headers: { Authorization: `Bearer ${token}` }
        })
        .then(() => {
            setBooks(books.filter(book => book.id !== id));
        })
        .catch(error => console.error('Failed to delete book:', error));
    };

    return (
        <div>
            <h1>Books Manager</h1>
            <form onSubmit={handleSubmit}>
                <input type="text" name="title" value={formData.title} onChange={handleInputChange} placeholder="Title" required />
                <input type="text" name="description" value={formData.description} onChange={handleInputChange} placeholder="Description" required />
                <input type="url" name="image_url" value={formData.image_url} onChange={handleInputChange} placeholder="Image URL" required />
                <input type="number" name="release_year" value={formData.release_year} onChange={handleInputChange} placeholder="Release Year" required />
                <input type="text" name="price" value={formData.price} onChange={handleInputChange} placeholder="Price" required />
                <input type="number" name="total_page" value={formData.total_page} onChange={handleInputChange} placeholder="Total Pages" required />
                <button type="submit">{editing ? "Update" : "Add"}</button>
            </form>
            <ul>
                {books.map((book, index) => (
                    <li key={book.id}>
                        {book.title} - {book.description}
                        <button onClick={() => { setEditing(true); setEditingId(book.id); setFormData(book); }}>Edit</button>
                        <button onClick={() => deleteBook(book.id)}>Delete</button>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default Books;
