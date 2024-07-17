import React, { useContext } from 'react';
import './Books.css'; // Mengimpor CSS untuk styling
import BooksContext from './BooksContext';

function Books() {
    const {
        books, formData, setFormData, handleInputChange, handleSubmit, loading, loadingMessage,
        editing, setEditing, setEditingId, deleteBook
    } = useContext(BooksContext);

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
                                    <button onClick={() => { setEditing(true); setEditingId(book.id); setFormData({
                                        title: book.title,
                                        description: book.description,
                                        image_url: book.image_url,
                                        release_year: book.release_year ? book.release_year.toString() : '',
                                        price: book.price,
                                        total_page: book.total_page ? book.total_page.toString() : ''
                                    }); }} className="btn-edit">Edit</button>
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
