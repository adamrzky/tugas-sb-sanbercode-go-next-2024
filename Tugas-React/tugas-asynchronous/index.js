var readBooks = require('./callback.js');

var books = [
    { name: 'LOTR', timeSpent: 3000 },
    { name: 'Fidas', timeSpent: 2000 },
    { name: 'Kalkulus', timeSpent: 4000 },
    { name: 'Komik', timeSpent: 1000 }
];

function readAllBooks(time, books, index) {
    if (index < books.length && time > 0) {
        readBooks(time, books[index], (sisaWaktu) => {
            readAllBooks(sisaWaktu, books, index + 1);
        });
    } else {
        console.log("Semua buku sudah selesai dibaca atau waktu sudah habis.");
    }
}

console.log('Soal 1');
readAllBooks(10000, books, 0);