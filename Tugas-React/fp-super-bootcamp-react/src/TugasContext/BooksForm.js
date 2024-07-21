// src/TugasContext/BooksForm.js
import React, { useState, useContext, useEffect } from 'react';
import { BooksContext } from './BooksContext';
import { useParams, useNavigate } from 'react-router-dom';

const BooksForm = () => {
  const { addBook, updateBook, books } = useContext(BooksContext);
  const { id } = useParams();
  const navigate = useNavigate();
  const isEditing = id !== undefined;

  const [formData, setFormData] = useState({
    title: '',
    description: '',
    image_url: '',
    release_year: '',
    price: '',
    total_page: ''
  });

  useEffect(() => {
    if (isEditing) {
      const book = books.find(b => b.id === parseInt(id));
      if (book) {
        setFormData(book);
      }
    }
  }, [id, books, isEditing]);

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    if (isEditing) {
      updateBook(id, formData);
    } else {
      addBook(formData);
    }
    navigate('/context');
  };

  return (
    <div className="container">
      <h1>{isEditing ? 'Edit Book' : 'Add New Book'}</h1>
      <form onSubmit={handleSubmit} className="form">
        <input type="text" name="title" value={formData.title} onChange={handleInputChange} placeholder="Title" required />
        <input type="text" name="description" value={formData.description} onChange={handleInputChange} placeholder="Description" required />
        <input type="url" name="image_url" value={formData.image_url} onChange={handleInputChange} placeholder="Image URL" required />
        <input type="number" name="release_year" value={formData.release_year} onChange={handleInputChange} placeholder="Release Year" required />
        <input type="text" name="price" value={formData.price} onChange={handleInputChange} placeholder="Price" required />
        <input type="number" name="total_page" value={formData.total_page} onChange={handleInputChange} placeholder="Total Pages" required />
        <button type="submit" className="btn-submit">{isEditing ? 'Update' : 'Add'}</button>
      </form>
    </div>
  );
};

export default BooksForm;
