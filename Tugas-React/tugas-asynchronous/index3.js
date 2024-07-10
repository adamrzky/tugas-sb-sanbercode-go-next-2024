// di file index3.js
var filterBooksPromise = require('./promise2.js');

// Lanjutkan code untuk menjalankan function filterBookPromise
console.log('Soal 3');

// Kondisi 1: Bukunya berwarna dan jumlah halamannya 50
filterBooksPromise(true, 50)
    .then(result => {
        console.log(result);
    })
    .catch(error => {
        console.log(error.message);
    });

// Kondisi 2: Bukunya tidak berwarna dan jumlah halamannya 250 (menggunakan async/await)
(async function() {
    try {
        const result = await filterBooksPromise(false, 250);
        console.log(result);
    } catch (error) {
        console.log(error.message);
    }
})();

// Kondisi 3: Bukunya berwarna dan jumlah halamannya 30 (menggunakan async/await)
(async function() {
    try {
        const result = await filterBooksPromise(true, 30);
        console.log(result);
    } catch (error) {
        console.log(error.message);
    }
})();

console.log(); // Baris kosong untuk pemisah
