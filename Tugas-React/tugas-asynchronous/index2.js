var readBooksPromise = require('./promise.js');

var books = [
    { name: 'LOTR', timeSpent: 3000 },
    { name: 'Fidas', timeSpent: 2000 },
    { name: 'Kalkulus', timeSpent: 4000 }
];

function readAllBooks(time, books, index) {
    if (index < books.length) {
        readBooksPromise(time, books[index])
            .then((sisaWaktu) => {
                readAllBooks(sisaWaktu, books, index + 1);
            })
            .catch((sisaWaktu) => {
                console.log("Waktu sudah habis atau tidak cukup untuk membaca buku berikutnya.");
            });
    } else {
        console.log("Semua buku sudah selesai dibaca atau waktu sudah habis.");
    }
}

console.log('Soal 2');
readAllBooks(10000, books, 0);
