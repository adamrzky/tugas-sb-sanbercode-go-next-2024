// src/TugasContext/BooksTable.js
import React, { useContext, useEffect } from 'react';
import { BooksContext } from './BooksContext';
import { Link } from 'react-router-dom';

const BooksTable = () => {
  const { books, fetchBooks, deleteBook } = useContext(BooksContext);

  useEffect(() => {
    fetchBooks();
  }, []);

  return (
    <div className="container">
      <h1>Books Manager</h1>
      <Link to="/context/create" className="btn-submit">Add New Book</Link>
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
                  <Link to={`/context/edit/${book.id}`} className="btn-edit">Edit</Link>
                  <button onClick={() => deleteBook(book.id)} className="btn-delete">Delete</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
        
      </div>
    </div>
  );
};

export default BooksTable;
